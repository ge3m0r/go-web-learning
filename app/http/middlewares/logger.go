package middlewares

import (
	"bytes"
	"golearning/pkg/helpers"
	"golearning/pkg/logger"
	"io"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"go.uber.org/zap"
)

type responeBodyWriter struct{
	gin.ResponseWriter
	body *bytes.Buffer
}

func (r responeBodyWriter) Write(b []byte)(int, error){
	r.body.Write(b)
	return r.ResponseWriter.Write(b)

}

func Logger() gin.HandlerFunc{
	return func(c *gin.Context) {
		w := &responeBodyWriter{body: &bytes.Buffer{},ResponseWriter: c.Writer,}
	    c.Writer = w

		var requestBody []byte
		if c.Request.Body != nil{

			requestBody,_ = io.ReadAll(c.Request.Body)

			c.Request.Body = io.NopCloser(bytes.NewBuffer(requestBody))
		}

		start := time.Now()
		c.Next()

		cost := time.Since(start)
		responStatus := c.Writer.Status()

		logFields := []zap.Field{
			zap.Int("status", responStatus),
			zap.String("request", c.Request.Method+" "+c.Request.URL.String()),
			zap.String("query", c.Request.URL.RawQuery),
			zap.String("ip",c.ClientIP()),
			zap.String("user-agent",c.Request.UserAgent()),
			zap.String("errors",c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.String("time",helpers.MicrosecondsStr(cost)),
		}

		if c.Request.Method == "POST" || c.Request.Method == "PUT" || c.Request.Method == "DELETE"{
			logFields = append(logFields, zap.String("Request Body", string(requestBody)))

			logFields = append(logFields, zap.String("Response Body", w.body.String()))
		}
		if responStatus > 400 && responStatus <= 499 {
            // 除了 StatusBadRequest 以外，warning 提示一下，常见的有 403 404，开发时都要注意
            logger.Warn("HTTP Warning "+cast.ToString(responStatus), logFields...)
        } else if responStatus >= 500 && responStatus <= 599 {
            // 除了内部错误，记录 error
            logger.Error("HTTP Error "+cast.ToString(responStatus), logFields...)
        } else {
            logger.Debug("HTTP Access Log", logFields...)
        }
	}
}