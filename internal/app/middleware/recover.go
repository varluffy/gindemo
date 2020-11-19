package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/varluffy/gindemo/internal/app/ginwrap"
	"github.com/varluffy/gindemo/pkg/errors"
	"github.com/varluffy/gindemo/pkg/logger"
	"net"
	"net/http/httputil"
	"os"
	"runtime/debug"
	"strings"
)

func RecoverMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			ctx := c.Request.Context()
			if err := recover(); err != nil {
				var brokenPipe bool
				if ne, ok := err.(*net.OpError); ok {
					if se, ok := ne.Err.(*os.SyscallError); ok {
						if strings.Contains(strings.ToLower(se.Error()), "broken pipe") || strings.Contains(strings.ToLower(se.Error()), "connection reset by peer") {
							brokenPipe = true
						}
					}
				}

				httpRequest, _ := httputil.DumpRequest(c.Request, false)
				if brokenPipe {
					logger.Error(ctx, "[Recovery from panic]", "err", err.(string), "request", string(httpRequest), "stack", string(debug.Stack()))
					c.Error(err.(error)) // nolint: errcheck
					c.Abort()
					return
				} else {
					logger.Error(ctx, "[Recovery from panic]", "err", err.(string), "request", string(httpRequest), "stack", string(debug.Stack()))
				}
				ginwrap.ResError(c, errors.ErrInternal)
				c.Abort()
				return
			}
		}()
		c.Next()
	}
}