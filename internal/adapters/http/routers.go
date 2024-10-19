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

			root := auth.Group("", handlers.RoleRootMiddleware())
			{
				root.GET("", userHandler.GetListUsers)
				root.POST("", userHandler.CreateUser)
				root.DELETE("/:id", userHandler.DeleteUserByID)
			}
		}
	}
}

// RegisterWarehouseRoute is a option function to return register warehouse router function
func RegisterWarehouseRoute(token ports.ITokenService, warehouseHandler *handlers.WarehouseHandler) RegisterRouterFunc {
	return func(e gin.IRouter) {
		auth := e.Group("/warehouses")
		auth.Use(handlers.AuthMiddleware(token))
		{
			auth.GET("", warehouseHandler.GetListWarehouses)
			auth.GET("/:id", warehouseHandler.GetWarehouseByID)
			auth.GET("/:id/used_capacity", warehouseHandler.GetUsedCapacityByID)

			root := auth.Group("", handlers.RoleRootMiddleware())
			{
				root.POST("", warehouseHandler.CreateWarehouse)
				root.PATCH("/:id", warehouseHandler.UpdateWarehouse)
				root.DELETE("/:id", warehouseHandler.DeleteWarehouse)
			}
		}
	}
}

// RegisterRiceRoute is a option function to return register rice router function
func RegisterRiceRoute(token ports.ITokenService, riceHandler *handlers.RiceHandler) RegisterRouterFunc {
	return func(e gin.IRouter) {
		auth := e.Group("/rice", handlers.AuthMiddleware(token))
		{
			auth.GET("", riceHandler.GetListRice)
			auth.GET("/:id", riceHandler.GetRiceByID)
			root := auth.Group("", handlers.RoleRootMiddleware())
			{
				root.POST("", riceHandler.CreateRice)
				root.PATCH("/:id", riceHandler.UpdateRice)
				root.DELETE("/:id", riceHandler.DeleteRice)
			}
		}
	}
}

// RegisterCustomerRoute is a option function to return register customer router function
func RegisterCustomerRoute(token ports.ITokenService, customerHandler *handlers.CustomerHandler) RegisterRouterFunc {
	return func(e gin.IRouter) {
		auth := e.Group("/customers", handlers.AuthMiddleware(token))
		{
			auth.GET("", customerHandler.GetListCustomers)
			auth.GET("/:id", customerHandler.GetCustomerByID)
			root := auth.Group("", handlers.RoleRootMiddleware())
			{
				root.POST("", customerHandler.CreateCustomer)
				root.PATCH("/:id", customerHandler.UpdateCustomer)
				root.DELETE("/:id", customerHandler.DeleteCustomer)
			}
		}
	}
}

// RegisterImportInvoiceRoute is a option function to return register import invoice router function
func RegisterImportInvoiceRoute(token ports.ITokenService, imInvHandler *handlers.ImportInvoiceHandler) RegisterRouterFunc {
	return func(e gin.IRouter) {
		auth := e.Group("/import_invoices", handlers.AuthMiddleware(token))
		{
			auth.POST("", imInvHandler.CreateImInvoice)
			auth.GET("", imInvHandler.GetListImInvoices)
			auth.GET("/:id", imInvHandler.GetImInvoiceByID)
			// root := auth.Group("", handlers.RoleRootMiddleware())
			// {
			// }
		}
	}
}
