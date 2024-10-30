package main

import (
	"github.com/gin-gonic/gin"
	"os"
	"social-network-app/internal/config"
	"social-network-app/internal/infrastructure/datastore"
	"social-network-app/internal/infrastructure/log"
	"social-network-app/internal/repositories"
	"social-network-app/internal/services"
	"social-network-app/internal/transport/api"
	"social-network-app/internal/transport/api/handlers"
)

func main() {
	defaultLog := log.InitDefaultLogger()

	cfg, err := config.Load(defaultLog)
	if err != nil {
		defaultLog.Error("Terminate execution", err)
	}

	logger := log.InitLogger()
	db := datastore.InitDB(cfg.DB, logger)
	r := gin.Default()

	userRepository := repositories.BuildUserRepository(db)
	meetingRepository := repositories.BuildMeetingRepository(db)

	userService := services.BuildUserService(userRepository, logger)
	meetingService := services.BuildMeetingService(meetingRepository, logger)

	userHandler := handlers.BuildUserHandler(userService, logger)
	meetingHandler := handlers.BuildMeetingHandler(meetingService, logger)

	server := api.BuildServer(r, logger, userHandler, meetingHandler)

	if err := server.Start(); err != nil {
		os.Exit(1)
	}
}
