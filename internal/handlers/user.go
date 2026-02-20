package handlers

import (
	"net/http"

	"github.com/Seasky89/go-report-api/internal/services"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	service *services.UserService
}

func NewUserHandler(service *services.UserService) *UserHandler {
	return &UserHandler{service: service}
}

func (h *UserHandler) FindAll(c *gin.Context) {
	users := h.service.FindAll()
	c.JSON(http.StatusOK, users)
}
