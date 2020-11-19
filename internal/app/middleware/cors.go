package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"time"
)

// CORSMiddleware 跨域中间件
func CORSMiddleware() gin.HandlerFunc {
	return cors.New(cors.Config{
		AllowOrigins:     viper.GetStringSlice("CORS.AllowOrigins"),
		AllowMethods:     viper.GetStringSlice("CORS.AllowMethods"),
		AllowHeaders:     viper.GetStringSlice("CORS.AllowHeaders"),
		AllowCredentials: viper.GetBool("CORS.AllowCredentials"),
		MaxAge:           time.Second * time.Duration(viper.GetInt64("CORS.MaxAge")),
	},
	)
}
