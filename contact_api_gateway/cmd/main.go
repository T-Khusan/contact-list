package main

import (
	// "fmt"

	"contact_api_gateway/api"
	"contact_api_gateway/config"
	"contact_api_gateway/pkg/logger"
	"contact_api_gateway/services"
	// "github.com/gomodule/redigo/redis"
)

func main() {
	cfg := config.Load()
	log := logger.New(cfg.LogLevel, "contact_api_gateway")

	gprcClients, _ := services.NewGrpcClients(&cfg)

	server := api.New(&api.RouterOptions{
		Log:      log,
		Cfg:      cfg,
		Services: gprcClients,
	})

	server.Run(cfg.HttpPort)
}
