package main

import (
	"github.com/tommjj/ql-kho-lua/internal/adapters/http"
	"github.com/tommjj/ql-kho-lua/internal/adapters/http/handlers"
	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"
	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb/repository"
	"github.com/tommjj/ql-kho-lua/internal/config"
	"github.com/tommjj/ql-kho-lua/internal/core/auth"
	"github.com/tommjj/ql-kho-lua/internal/core/services"
	"github.com/tommjj/ql-kho-lua/internal/logger"
	"go.uber.org/zap"
)

func main() {
	conf, err := config.New()
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	err = logger.Set(*conf.Logger)
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	db, err := mysqldb.NewMysqlDB(*conf.DB)
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	userRepository := repository.NewUserRepository(db)

	tokenService := auth.NewJWTTokenService(*conf.Auth)
	authService := services.NewAuthService(userRepository, tokenService)

	authHandler := handlers.NewAuthHandler(authService)

	server, err := http.NewAdapter(conf.Http, http.RegisterPingRoute(),
		http.Group("/v1/api",
			http.RegisterAuthRoute(authHandler),
		),
	)
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	server.Serve()
}
