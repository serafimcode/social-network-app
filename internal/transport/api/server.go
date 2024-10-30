package api

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"social-network-app/internal/transport/api/handlers"
)

type Server struct {
	Router         *gin.Engine
	log            *slog.Logger
	userHandler    *handlers.UserHandler
	meetingHandler *handlers.MeetingHandler
}

func BuildServer(
	r *gin.Engine,
	log *slog.Logger,
	userHandler *handlers.UserHandler,
	meetingHandler *handlers.MeetingHandler,
) *Server {
	s := Server{
		Router:         r,
		log:            log,
		userHandler:    userHandler,
		meetingHandler: meetingHandler,
	}

	api := r.Group("/api")
	{
		api.GET("/users", s.userHandler.GetUsers)

		api.GET("/meetings", s.meetingHandler.GetMeetings)
		api.POST("/meetings", s.meetingHandler.CreateMeeting)
	}

	return &s
}

func (s *Server) Start() error {
	s.log.Info("Starting server...")
	return http.ListenAndServe(":8080", s.Router)
}
