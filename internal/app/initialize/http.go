package initialize

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"github.com/varluffy/gindemo/pkg/logger"
	"net/http"
	"time"
)

// InitHTTPServer 初始化HTTP服务
func InitHTTPServer(ctx context.Context, handler http.Handler) func() {
	addr := fmt.Sprintf("%s:%d",
		viper.GetString("HTTP.Host"),
		viper.GetInt("HTTP.Port"),
	)
	server := http.Server{
		Addr:    addr,
		Handler: handler,
	}

	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logger.Fatal(ctx, "start http server error", err.Error())
		}
	}()
	return func() {
		ctx, cancel := context.WithTimeout(ctx, time.Second*time.Duration(viper.GetInt64("HTTP.ShutdownTimeout")))
		defer cancel()
		server.SetKeepAlivesEnabled(false)
		if err := server.Shutdown(ctx); err != nil {
			logger.Error(ctx, "http server shutdown error", err.Error())
		}
	}
}
