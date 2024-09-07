package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Group is a option function to group register router functions.
func Group(path string, registerRouterFuncs ...RegisterRouterFunc) RegisterRouterFunc {
	return func(r gin.IRouter) {
		g := r.Group(path)
		for _, fn := range registerRouterFuncs {
			fn(g)
		}
	}
}

// RegisterPingRoute is a option function to return register ping router function.
func RegisterPingRoute() RegisterRouterFunc {
	return func(r gin.IRouter) {
		r.GET("/ping", func(ctx *gin.Context) {
			ctx.JSON(http.StatusOK, gin.H{"message": "pong"})
		})
	}
}
