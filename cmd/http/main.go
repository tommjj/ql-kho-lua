package main

import (
	"github.com/tommjj/ql-kho-lua/internal/adapters/http"
	"github.com/tommjj/ql-kho-lua/internal/config"
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

	server, err := http.NewAdapter(conf.Http, http.RegisterPingRoute())
	if err != nil {
		zap.L().Fatal(err.Error())
	}

	server.Serve()
}
