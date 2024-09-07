package main

import (
	"github.com/tommjj/ql-kho-lua/internal/adapters/http"
	"github.com/tommjj/ql-kho-lua/internal/config"
	"github.com/tommjj/ql-kho-lua/internal/logger"
)

func main() {
	conf, err := config.New()
	if err != nil {
		logger.Fatal(err.Error())
	}

	err = logger.Set(*conf.Logger)
	if err != nil {
		logger.Fatal(err.Error())
	}

	server, err := http.New(conf.Http, http.RegisterPingRoute())
	if err != nil {
		logger.Fatal(err.Error())
	}

	server.Serve()
}
