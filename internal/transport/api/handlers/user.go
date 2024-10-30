package handlers

import (
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"social-network-app/internal/services"
	"social-network-app/internal/transport/domain_errors"
	"social-network-app/internal/transport/dto"
	"social-network-app/internal/transport/utils"
)

type UserHandler struct {
	log         *slog.Logger
	userService *services.UserService
}

func BuildUserHandler(userService *services.UserService, log *slog.Logger) *UserHandler {
	return &UserHandler{userService: userService, log: log}
}

func (uh *UserHandler) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()
	users, err := uh.userService.GetUsers(ctx)

	if err != nil {
		uh.log.Error("Failed to get users", err)
		utils.BuildErrorResponse(c, http.StatusInternalServerError, domain_errors.InternalError)
		return
	}

	usersDto := dto.MapUsersToDto(users)

	utils.BuildSuccessResponse(c, usersDto)
}
