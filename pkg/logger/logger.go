package logger

import (
	"encoding/json"
	"fmt"
	"golearning/pkg/app"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *zap.Logger

func InitLogger(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string, level string){
	writeSyncer := getLogWriter(filename,maxSize,maxBackup,maxAge,compress,logType)

	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil{
		fmt.Println("日志初始化错误，日志级别设置有误。请修改 config/log.go 文件中的 log.lever配置项")
	}

	core := zapcore.NewCore(getEncoder(), writeSyncer,logLevel)
	Logger = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1), zap.AddStacktrace(zap.ErrorLevel))

	zap.ReplaceGlobals(Logger)
}

func getEncoder() zapcore.Encoder{
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "caller",
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.CapitalColorLevelEncoder,
		EncodeTime:     customTimeEncoder,
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	}

	if app.IsLocal(){
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		return zapcore.NewConsoleEncoder(encoderConfig)
	}
	return zapcore.NewJSONEncoder(encoderConfig)
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder){
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string) zapcore.WriteSyncer{
	if logType == "daily"{
		logname := time.Now().Format("2006-01-02.log")
		filename = strings.ReplaceAll(filename,"logs.log",logname)
	}

	lumberjackLogger := &lumberjack.Logger{
		Filename:     filename,
		MaxSize:      maxSize,
		MaxBackups:   maxBackup,
		MaxAge:       maxAge,
		Compress:     compress,
	}

	if app.IsLocal(){
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberjackLogger))
	}else{
		return zapcore.AddSync(lumberjackLogger)
	}

}

func Dump(value interface{}, msg ...string){
	vauleString := jsonString(value)

	if len(msg) > 0{
		Logger.Warn("Dump", zap.String(msg[0], vauleString))
	}else{
		Logger.Warn("Dump", zap.String("data", vauleString))
	}
}

// LogIf 当 err != nil 时记录 error 等级的日志
func LogIf(err error) {
    if err != nil {
        Logger.Error("Error Occurred:", zap.Error(err))
    }
}

func logWarnIf(err error){
	if err != nil{
		Logger.Warn("Error Occurred:", zap.Error(err))
	}
}

func logInfoIf(err error){
	if err != nil{
		Logger.Info("Error Occurred:", zap.Error(err))
	}
}

func Debug(moduleName string, fields ...zap.Field){
	Logger.Debug(moduleName,fields...)
}

func Info(moduleName string, fields ...zap.Field){
	Logger.Info(moduleName,fields...)
}

func Warn(moduleName string, fields ...zap.Field){
	Logger.Warn(moduleName,fields...)
}

func Error(moduleName string, fields ...zap.Field){
	Logger.Error(moduleName,fields...)
}

func Fatal(moduleName string, fields ...zap.Field){
	Logger.Fatal(moduleName,fields...)
}

func DebugString(moduleName, name, msg string){
	Logger.Debug(moduleName,zap.String(name,msg))
}

func InfoString(moduleName, name, msg string){
	Logger.Info(moduleName,zap.String(name,msg))
}

func WarnString(moduleName, name, msg string){
	Logger.Warn(moduleName,zap.String(name,msg))
}

func ErrorString(moduleName, name, msg string){
	Logger.Error(moduleName,zap.String(name,msg))
}



func FatalString(moduleName, name, msg string){
	Logger.Debug(moduleName,zap.String(name,msg))
}

func DebugJson(moduleName, name string, value interface{}){
	Logger.Debug(moduleName,zap.String(name,jsonString(value)))
}

func InfoJson(moduleName, name string, value interface{}){
	Logger.Info(moduleName,zap.String(name,jsonString(value)))
}

func WarnJson(moduleName, name string, value interface{}){
	Logger.Warn(moduleName,zap.String(name,jsonString(value)))
}

func ErrorJson(moduleName, name string, value interface{}){
	Logger.Error(moduleName,zap.String(name,jsonString(value)))
}

func FatalJson(moduleName, name string, value interface{}){
	Logger.Fatal(moduleName,zap.String(name,jsonString(value)))
}

func jsonString(value interface{}) string {
	b, err := json.Marshal(value)
	if err != nil{
		Logger.Error("Logger", zap.String("JSON marshal error", err.Error()))
	}
	return string(b)
}
