package http

import (
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/tommjj/ql-kho-lua/internal/config"
	"github.com/tommjj/ql-kho-lua/internal/logger"
)

func WithLogger(conf config.Logger) RegisterRouterFunc {
	return func(r gin.IRouter) {
		// set logger middleware
		logger, err := logger.New(conf)
		if err != nil {
			panic(err)
		}

		r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
		r.Use(ginzap.RecoveryWithZap(logger, true))
	}
}
