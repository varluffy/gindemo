package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/varluffy/gindemo/internal/app/ginwrap"
	"github.com/varluffy/gindemo/pkg/logger"
	"github.com/varluffy/gindemo/pkg/util"
	"net/http"
	"time"
)

// TraceMiddleware 追踪日志
func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceId := c.GetHeader("X-Request-Id")
		if traceId == "" {
			traceId = util.NewTraceID()
		}
		ctx := logger.NewTraceIDContext(c.Request.Context(), traceId)
		c.Request = c.Request.WithContext(ctx)

		p := c.Request.URL.Path
		method := c.Request.Method
		span := logger.StartSpan(c.Request.Context(), logger.SetSpanTitle("trace log"))
		start := time.Now()
		fields := make(map[string]interface{})
		fields["ip"] = c.ClientIP()
		fields["method"] = c.Request.Method
		fields["url"] = c.Request.URL.String()
		fields["header"] = c.Request.Header
		fields["user_agent"] = c.Request.UserAgent()
		if method == http.MethodPost || method == http.MethodPut {

		}
		c.Next()
		if v, ok := c.Get(ginwrap.ResBodyKey); ok {
			if buf, ok := v.([]byte); ok {
				fields["res_body"] = string(buf)
			}
		}
		timeConsume := time.Since(start) / 1e6
		fields["res_status"] = c.Writer.Status()
		fields["res_length"] = c.Writer.Size()
		fields["user_id"] = ginwrap.GetUserID(c)
		span.WithFields(fields).Infof("[http] %s-%s-%s-%d(%dms)", p, method, c.ClientIP(), c.Writer.Status(), timeConsume)
	}
}
