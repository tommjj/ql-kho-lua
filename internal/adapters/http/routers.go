package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/tommjj/ql-kho-lua/internal/adapters/http/handlers"
	"github.com/tommjj/ql-kho-lua/internal/core/ports"
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

// RegisterAuthRoute is a option function to return register auth router function
func RegisterAuthRoute(authHandler *handlers.AuthHandler) RegisterRouterFunc {
	return func(e gin.IRouter) {
		r := e.Group("/auth")
		{
			r.POST("/login", authHandler.Login)
		}
	}
}

// RegisterUploadRoute is a option function to return register upload router function
func RegisterUploadRoute(uploadHandler *handlers.UploadHandler) RegisterRouterFunc {
	return func(e gin.IRouter) {
		e.POST("/upload", uploadHandler.UploadImage)
	}
}

// RegisterStatic is a option function to return register static router function
func RegisterStatic(root string) RegisterRouterFunc {
	return func(i gin.IRouter) {
		i.Static("/static", root)
	}
}

// RegisterUsersRoute is a option function to return register user router function
func RegisterUsersRoute(token ports.ITokenService, userHandler *handlers.UserHandler) RegisterRouterFunc {
	return func(e gin.IRouter) {
		auth := e.Group("/users", handlers.AuthMiddleware(token))
		{
			auth.GET("/:id", userHandler.GetUserByID)
			auth.PATCH("/:id", userHandler.UpdateUser)

			root := auth.Group("/", handlers.RoleRootMiddleware())
			{

				root.GET("/", userHandler.GetListUsers)
				root.POST("/", userHandler.CreateUser)
				root.DELETE("/:id", userHandler.DeleteUserByID)
			}
		}
	}
}
