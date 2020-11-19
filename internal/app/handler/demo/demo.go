package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	"github.com/varluffy/gindemo/internal/app/ginwrap"
	"github.com/varluffy/gindemo/internal/logic"
	"github.com/varluffy/gindemo/pkg/errors"
	"time"
)

type Demo struct {
	logic       *logic.Logic
	redisClient *redis.ClusterClient
}

func NewDemo(logic *logic.Logic, client *redis.ClusterClient) *Demo {
	return &Demo{
		logic:       logic,
		redisClient: client,
	}
}

func (a *Demo) Get(c *gin.Context) {
	b, err := a.redisClient.SetNX(c.Request.Context(), "123", "1231", time.Second*2).Result()
	if err != nil {
		ginwrap.ResError(c, err)
		return
	}
	if !b {
		ginwrap.ResError(c, errors.ErrorTooManyRequests)
		return
	}
	m := make(map[string]string)
	m["demo"] = "demo"
	ginwrap.ResSuccess(c, m)
}


func (a *Demo) Post(c *gin.Context) {
	type req struct {
		A int `form:"a"`
		B string `form:"b"`
	}
	var r req
	if err := ginwrap.ParseForm(c, &r); err != nil {
		ginwrap.ResError(c, err)
		return
	}
	ginwrap.ResSuccess(c, r)
}