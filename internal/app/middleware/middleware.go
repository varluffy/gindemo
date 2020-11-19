package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/varluffy/gindemo/internal/app/ginwrap"
	"github.com/varluffy/gindemo/pkg/errors"
)

// NoMethodMiddleware 405
func NoMethodMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ginwrap.ResError(c, errors.ErrorMethodNotAllow)
	}
}

// NoRouteMiddleware 404
func NoRouteMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ginwrap.ResError(c, errors.ErrNotFound)
	}
}
