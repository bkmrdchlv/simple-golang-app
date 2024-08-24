package main

import (
	"market_go_apigateway/api"
	"market_go_apigateway/config"
	"market_go_apigateway/grpc/client"
	"market_go_apigateway/pkg/logger"
	"market_go_apigateway/api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	var loggerLevel string

	switch cfg.Environment {
	case config.DebugMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}

	log := logger.NewLogger(cfg.ServiceName, loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			log.Error("!!!Cleanup->Error-->>>", logger.Error(err))
		}
	}()

	svcs, err := client.NewGrpcClients(cfg)
	if err != nil {
		log.Panic("client.NewGrpcClients", logger.Error(err))
	}

	h := handlers.NewHandler(cfg, log, svcs)

	r := api.SetUpRouter(h, cfg)

	log.Info("HTTP: Server being started...", logger.String("port", cfg.HTTPPort))

	err = r.Run(cfg.HTTPPort)
	if err != nil {
		panic(err)
	}
}
