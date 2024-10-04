package main

import (
	"time"

	"github.com/robfig/cron/v3"
	"github.com/tommjj/ql-kho-lua/internal/adapters/http"
	"github.com/tommjj/ql-kho-lua/internal/adapters/http/handlers"
	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/files"
	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb"
	"github.com/tommjj/ql-kho-lua/internal/adapters/storage/mysqldb/repository"
	"github.com/tommjj/ql-kho-lua/internal/config"
	"github.com/tommjj/ql-kho-lua/internal/core/auth"
	"github.com/tommjj/ql-kho-lua/internal/core/services"
	"github.com/tommjj/ql-kho-lua/internal/logger"
	"go.uber.org/zap"
)

// ql-kho-lua
//
//	@title						Qua Ly Kho Lua
//	@version					1.0
//	@description				This is a RESTful ql-kho-lua.
//
//	@BasePath					/v1/api
//	@schemes					http https
//
//	@securityDefinitions.apikey	JWTAuth
//	@in							header
//	@name						Authorization
//	@description				Type "JWT" followed by a space and the access token.
func main() {
	conf, err := config.New()
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	err = logger.Set(*conf.Logger)
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	// |> Start Storage
	zap.L().Info("Start create Storage")

	db, err := mysqldb.NewMysqlDB(*conf.DB)
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	fileStorage, err := files.NewFileStorage("./public", "./public/temp", time.Hour)
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	// |> Start CRON
	zap.L().Info("Start CRON")

	c := cron.New()
	// Every hour
	_, err = c.AddFunc("0 * * * *", func() {
		zap.L().Info("clean temp files")
		fileStorage.CleanupTempFiles()
	})
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	c.Start()

	// |> Start Repository
	zap.L().Info("Start create repository")

	userRepository := repository.NewUserRepository(db)

	// |> Start Service
	zap.L().Info("Start create service")

	uploadService := services.NewUploadService(fileStorage)
	tokenService := auth.NewJWTTokenService(*conf.Auth)
	authService := services.NewAuthService(userRepository, tokenService)

	// |> Start Handler
	zap.L().Info("Start create handler")

	uploadHandler := handlers.NewUploadHandler(uploadService)
	authHandler := handlers.NewAuthHandler(authService)

	// |> Start HTTP Server
	zap.L().Info("Start create http server")

	server, err := http.NewAdapter(conf.Http, http.RegisterPingRoute(),
		http.RegisterStatic("./public"),
		http.Group("/v1/api",
			http.RegisterUploadRoute(uploadHandler),
			http.RegisterAuthRoute(authHandler),
		),
	)
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	server.Serve()
}
