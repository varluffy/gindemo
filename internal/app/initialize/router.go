package initialize

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/varluffy/gindemo/internal/app/handler"
	"github.com/varluffy/gindemo/internal/app/middleware"
)

// InitRouter 初始化路由
func InitRouter(router *handler.Router) *gin.Engine {
	gin.SetMode(viper.GetString("RunMode"))
	engine := gin.New()
	engine.NoMethod(middleware.NoMethodMiddleware())
	engine.NoRoute(middleware.NoRouteMiddleware())

	// 日志追踪
	engine.Use(middleware.TraceMiddleware())
	// 跨域
	engine.Use(middleware.CORSMiddleware())
	engine.Use(middleware.RecoverMiddleware())
	router.Register(engine)
	return engine
}
