package middleware

import (
	"bytes"
	"io/ioutil"
	"time"
	"wapi/src/pkg/logging"

	"github.com/gin-gonic/gin"
)

func StructuredLog(logger logging.Logger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		start := time.Now()
		path := ctx.FullPath()
		raw := ctx.Request.URL.RawQuery
		bdyByte, _ := ioutil.ReadAll(ctx.Request.Body)
		ctx.Request.Body.Close()
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bdyByte))
		ctx.Next()

		param := gin.LogFormatterParams{}
		param.TimeStamp = time.Now()
		param.Latency = param.TimeStamp.Sub(start)
		param.ClientIP = ctx.ClientIP()
		param.Method = ctx.Request.Method
		param.ErrorMessage = ctx.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = ctx.Writer.Size()

		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		keys := map[logging.ExtraKey]interface{}{}

		keys[logging.ClientIp] = param.ClientIP
		keys[logging.Method] = param.Method
		keys[logging.Latency] = param.TimeStamp
		keys[logging.BodySize] = param.BodySize
		keys[logging.ErrorMessage] = param.ErrorMessage
		keys[logging.RequestBody] = string(bdyByte)
		keys[logging.Path] = param.Path

		logger.Info(logging.RequestResponse, logging.Api, "", keys)
	}
}
