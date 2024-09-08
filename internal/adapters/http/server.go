package http

import (
	"errors"
	"fmt"
	"time"

	"github.com/tommjj/ql-kho-lua/internal/config"
	"go.uber.org/zap"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/tommjj/ql-kho-lua/internal/logger"
)

type RegisterRouterFunc func(gin.IRouter)

type router struct {
	engine *gin.Engine
	Port   int
	Url    string
}

func NewAdapter(conf *config.HTTP, options ...RegisterRouterFunc) (*router, error) {
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
	CORSConfig := cors.DefaultConfig()
	CORSConfig.AllowOrigins = conf.AllowedOrigins
	CORSConfig.AllowCredentials = true
	r.Use(cors.New(CORSConfig))

	for _, option := range options {
		option(r)
	}

	return &router{
		engine: r,
		Port:   conf.Port,
		Url:    conf.URL,
	}, nil
}

// Serve is a method start server
func (r *router) Serve() {
	zap.L().Info(fmt.Sprintf("start server at http://%v:%v", r.Url, r.Port))

	err := r.engine.Run(fmt.Sprintf("%v:%v", r.Url, r.Port))
	if err != nil {
		zap.L().Error(err.Error())
	}
}
