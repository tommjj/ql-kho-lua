package http

import (
	"fmt"
	"time"

	"github.com/go-playground/validator/v10"
	custom_validator "github.com/tommjj/ql-kho-lua/internal/adapters/http/validator"
	"github.com/tommjj/ql-kho-lua/internal/config"
	"go.uber.org/zap"

	"github.com/gin-contrib/cors"
	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/tommjj/ql-kho-lua/internal/logger"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	docs "github.com/tommjj/ql-kho-lua/docs"
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
		return nil, err
	}

	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	// set CORS
	CORSConfig := cors.DefaultConfig()
	CORSConfig.AllowOrigins = conf.AllowedOrigins
	CORSConfig.AllowCredentials = true
	CORSConfig.AllowHeaders = []string{"authorization"}
	r.Use(cors.New(CORSConfig))

	// Custom validators
	v, ok := binding.Validator.Engine().(*validator.Validate)
	if ok {
		if err := v.RegisterValidation("user_role", custom_validator.UserRoleValidator); err != nil {
			return nil, err
		}
		if err := v.RegisterValidation("location", custom_validator.LocationValidator); err != nil {
			return nil, err
		}
		if err := v.RegisterValidation("image_file", custom_validator.ImageFileValidator); err != nil {
			return nil, err
		}
	}

	// Swagger
	docs.SwaggerInfo.BasePath = "/v1/api"
	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

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
