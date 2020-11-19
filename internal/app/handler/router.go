package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	handler "github.com/varluffy/gindemo/internal/app/handler/demo"
	"github.com/varluffy/gindemo/internal/logic"
)

type Router struct {
	demo        *handler.Demo
	redisClient *redis.ClusterClient
}

// NewRouter 创建路由
func NewRouter(logic *logic.Logic, client *redis.ClusterClient) *Router {
	return &Router{
		demo: handler.NewDemo(logic, client),
	}
}

// Register 路由注册
func (a *Router) Register(r *gin.Engine) {
	g := r.Group("/api")
	g.GET("demo", a.demo.Get)
	g.POST("demo", a.demo.Post)
}
