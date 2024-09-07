package http

import (
	"errors"
	"fmt"
	"time"

	"github.com/tommjj/ql-kho-lua/internal/config"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/tommjj/ql-kho-lua/internal/logger"
)

type RegisterRouterFunc func(gin.IRouter)

type Router struct {
	*gin.Engine
	Port int
	Url  string
}

func New(conf *config.HTTP, options ...RegisterRouterFunc) (*Router, error) {
	if conf.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.New()

	// set logger middleware
	logger, err := logger.New(conf.Logger)
	if err != nil {
		return nil, errors.New("http logger conf is not valid")
	}

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	// set CORS
	ginConfig := cors.DefaultConfig()
	ginConfig.AllowOrigins = conf.AllowedOrigins
	r.Use(cors.New(ginConfig))

	for _, option := range options {
		option(r)
	}

	return &Router{
		Engine: r,
		Port:   conf.Port,
		Url:    conf.URL,
	}, nil
}

// Serve is a method start server
func (r *Router) Serve() {
	logger.Info(fmt.Sprintf("start server at http://%v:%v", r.Url, r.Port))

	err := r.Run(fmt.Sprintf("%v:%v", r.Url, r.Port))
	if err != nil {
		logger.Error(err.Error())
	}
}
