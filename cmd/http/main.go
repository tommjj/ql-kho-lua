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
	"github.com/tommjj/ql-kho-lua/internal/core/mapmutex"
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
	// |> Start load config
	conf, err := config.New()
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	// |> Start set logger
	err = logger.Set(*conf.Logger)
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	// |> Start Storage
	zap.L().Info("Start create Storage")

	db, err := Retry(func() (*mysqldb.MysqlDB, error) {
		return mysqldb.NewMysqlDB(*conf.DB)
	}, time.Second, 10)
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	fileStorage, err := files.NewFileStorage("./public/", "./public/temp", time.Hour)
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	defer fileStorage.CleanupTempFiles()

	// |> Start CRON
	zap.L().Info("Start CRON")

	c := cron.New()
	_, err = c.AddFunc("@hourly", func() {
		zap.L().Info("clean temp files")
		fileStorage.CleanupTempFiles()
	})
	if err != nil {
		zap.L().Fatal(err.Error())
	}
	c.Start()

	// |> Start Repository
	zap.L().Info("Start create repository")

	keyRepository := repository.NewKeyRepository(db)
	userRepository := repository.NewUserRepository(db)
	storehouseRepository := repository.NewStorehouseRepository(db)
	accessControlRepository := repository.NewAccessControlRepository(db)
	riceRepository := repository.NewRiceRepository(db)
	customerRepository := repository.NewCustomerRepository(db)
	imInvoiceRepository := repository.NewImInvoicesRepository(db)

	// |> Start Service
	zap.L().Info("Start create service")

	uploadService := services.NewUploadService(fileStorage)
	tokenService := auth.NewJWTTokenService(*conf.Auth, keyRepository)
	authService := services.NewAuthService(userRepository, tokenService)
	userService := services.NewUserService(userRepository)
	accessControlService := services.NewAccessControlService(accessControlRepository)
	storehouseService := services.NewStorehouseService(storehouseRepository, fileStorage)
	riceService := services.NewRiceService(riceRepository)
	customerService := services.NewCustomerService(customerRepository)
	imInvoiceService := services.NewImInvoicesService(imInvoiceRepository, storehouseRepository, &mapmutex.Mapmutex{})

	// |> Start Handler
	zap.L().Info("Start create handler")

	uploadHandler := handlers.NewUploadHandler(uploadService)
	authHandler := handlers.NewAuthHandler(authService)
	userHandler := handlers.NewUserHandler(userService)
	storeHouseHandler := handlers.NewStorehouseHandler(storehouseService, accessControlService)
	riceHandler := handlers.NewRiceHandler(riceService)
	customerHandler := handlers.NewCustomerHandler(customerService)
	imInvoiceHandler := handlers.NewImportInvoiceHandler(imInvoiceService, accessControlService)

	// |> Start HTTP Server
	zap.L().Info("Start create http server")

	server, err := http.NewAdapter(conf.Http, http.RegisterPingRoute(),
		http.RegisterStatic("./public"),
		http.Group("/v1/api",
			http.RegisterUploadRoute(uploadHandler),
			http.RegisterAuthRoute(authHandler),
			http.RegisterUsersRoute(tokenService, userHandler),
			http.RegisterStorehouseRoute(tokenService, storeHouseHandler),
			http.RegisterRiceRoute(tokenService, riceHandler),
			http.RegisterCustomerRoute(tokenService, customerHandler),
			http.RegisterImportInvoiceRoute(tokenService, imInvoiceHandler),
		),
	)
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	server.Serve()
}

// Retry is a helper func to retry a function fc a specified number of times if it encounters an error.
func Retry[T any](fc func() (T, error), duration time.Duration, times int) (T, error) {
	var result T
	var err error

	result, err = fc()
	if err == nil {
		return result, nil
	}

	for i := range times {
		zap.S().Infof("times: %v. retry...", i+1)
		result, err = fc()
		if err == nil {
			zap.L().Info("success")
			return result, err
		}

		zap.L().Error(err.Error())
		time.Sleep(duration)
	}
	return result, err
}
