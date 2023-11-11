package main

import (
	"api_gateway_task/config"
	"api_gateway_task/pkg/logger"
	"api_gateway_task/services"
	"fmt"

	"api_gateway_task/api"
	"api_gateway_task/api/handlers"

	// "api_gateway_task/api/handlers"
	// "api_gateway_task/config"
	// "api_gateway_task/generates"
	// "api_gateway_task/pkg/logger"
	// "api_gateway_task/services"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg := config.Load()

	fmt.Printf("config: %+v\n", cfg)

	grpcSvcs, err := services.NewGrpcClients(cfg)
	if err != nil {
		panic(err)
	}

	var loggerLevel = new(string)

	*loggerLevel = logger.LevelDebug

	switch cfg.Environment {
	case config.DebugMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.DebugMode)
	case config.TestMode:
		*loggerLevel = logger.LevelDebug
		gin.SetMode(gin.TestMode)
	default:
		*loggerLevel = logger.LevelInfo
		gin.SetMode(gin.ReleaseMode)
	}
	log := logger.NewLogger("tu_go_admin_api_gateway", *loggerLevel)
	defer func() {
		err := logger.Cleanup(log)
		if err != nil {
			return
		}
	}()

	r := gin.New()

	r.Use(gin.Logger(), gin.Recovery())

	h := handlers.NewHandler(cfg, log, grpcSvcs)

	api.SetUpAPI(r, h, cfg)

	fmt.Println("Start api gateway....")

	if err := r.Run(cfg.HTTPPort); err != nil {
		return
	}
}
