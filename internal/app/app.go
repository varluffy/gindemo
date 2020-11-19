package app

import (
	"context"
	"github.com/spf13/viper"
	"github.com/varluffy/gindemo/internal/app/handler"
	"github.com/varluffy/gindemo/internal/app/initialize"
	"github.com/varluffy/gindemo/internal/logic"
	"github.com/varluffy/gindemo/internal/repository"
	"github.com/varluffy/gindemo/pkg/logger"
	"os"
	"os/signal"
	"syscall"
)

type options struct {
	ConfigFile string
	Version    string
}

type Option func(*options)

func Run(ctx context.Context, opts ...Option) error {
	logger.Infof(ctx, "服务启动, 运行模式:%s，版本号：%s，进程号：%s", viper.GetString("RunMode"), "1.0.0", os.Getegid())
	db, cleanDBFunc, err := initialize.InitDB()
	if err != nil {
		return err
	}
	repo := repository.NewRepository(db)
	client, cleanRedisFunc, err := initialize.InitRedis()
	if err != nil {
		return err
	}
	lo := logic.NewLogic(repo)
	r := handler.NewRouter(lo, client)
	engine := initialize.InitRouter(r)
	cleanHttpFunc := initialize.InitHTTPServer(ctx, engine)

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	select {
	case sig := <-c:
		logger.Infof(ctx, "接收到信号[%s]", sig.String())
		cleanDBFunc()
		cleanRedisFunc()
		cleanHttpFunc()
		logger.Info(ctx, "服务退出")
		os.Exit(0)
	}
	return nil
}
