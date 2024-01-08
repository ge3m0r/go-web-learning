package config

import (
	"golearning/pkg/helpers"
	"os"
	
	"github.com/spf13/cast"
	viperlib "github.com/spf13/viper"
)

var viper *viperlib.Viper

type ConfigFunc func() map[string]interface{}

var ConfigFuncs map[string]ConfigFunc

func init(){
	viper = viperlib.New()
	viper.SetConfigType("env")
	viper.AddConfigPath(".")
	viper.SetEnvPrefix("appenv")
	viper.AutomaticEnv()

	ConfigFuncs = make(map[string]ConfigFunc)
}

func InitConfig(env string){
	loadEnv(env)
	loadConfig()
}

func loadConfig(){
	for name, fn := range ConfigFuncs{
		viper.Set(name,fn())
	}
}

func loadEnv(envSuffix string){
	var envPath string = ".env"
	if len(envSuffix) > 0{
		filepath := ".env." + envSuffix
		if _,err := os.Stat(filepath); err == nil{
			envPath = filepath
		}
	}
	viper.SetConfigName(envPath)
	if err := viper.ReadInConfig(); err != nil{
		panic(err)
	}
	viper.WatchConfig()

}

func Env(envName string,defaultValue ...interface{})interface{}{
	if len(defaultValue) > 0{
		return internalGet(envName,defaultValue[0])
	}
	return internalGet(envName)
}

func Add(name string, configFn ConfigFunc){
	ConfigFuncs[name] = configFn
}

func Get(path string, defaultValue ...interface{}) string{
	return GetString(path,defaultValue...)
}
func internalGet(path string, defaultValue ...interface{}) interface{} {
    // config 或者环境变量不存在的情况
    if !viper.IsSet(path) || helpers.Empty(viper.Get(path)) {
        if len(defaultValue) > 0 {
            return defaultValue[0]
        }
        return nil
    }
    return viper.Get(path)
}

// GetString 获取 String 类型的配置信息
func GetString(path string, defaultValue ...interface{}) string {
    return cast.ToString(internalGet(path, defaultValue...))
}

// GetInt 获取 Int 类型的配置信息
func GetInt(path string, defaultValue ...interface{}) int {
    return cast.ToInt(internalGet(path, defaultValue...))
}

// GetFloat64 获取 float64 类型的配置信息
func GetFloat64(path string, defaultValue ...interface{}) float64 {
    return cast.ToFloat64(internalGet(path, defaultValue...))
}

// GetInt64 获取 Int64 类型的配置信息
func GetInt64(path string, defaultValue ...interface{}) int64 {
    return cast.ToInt64(internalGet(path, defaultValue...))
}

// GetUint 获取 Uint 类型的配置信息
func GetUint(path string, defaultValue ...interface{}) uint {
    return cast.ToUint(internalGet(path, defaultValue...))
}

// GetBool 获取 Bool 类型的配置信息
func GetBool(path string, defaultValue ...interface{}) bool {
    return cast.ToBool(internalGet(path, defaultValue...))
}

// GetStringMapString 获取结构数据
func GetStringMapString(path string) map[string]string {
    return viper.GetStringMapString(path)
}
