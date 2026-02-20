package handlers

import (
	"net/http"

	"github.com/Seasky89/go-report-api/internal/services"

	"github.com/gin-gonic/gin"
)

type PostHandler struct {
	service *services.PostService
}

func NewPostHandler(service *services.PostService) *PostHandler {
	return &PostHandler{service: service}
}

func (h *PostHandler) FindAll(c *gin.Context) {
	posts := h.service.FindAll()
	c.JSON(http.StatusOK, posts)
}
