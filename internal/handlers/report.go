package handlers

import (
	"net/http"

	"github.com/Seasky89/go-report-api/internal/services"

	"github.com/gin-gonic/gin"
)

type ReportHandler struct {
	service *services.ReportService
}

func NewReportHandler(service *services.ReportService) *ReportHandler {
	return &ReportHandler{service: service}
}

func (h *ReportHandler) GenerateReport(c *gin.Context) {
	report := h.service.GenerateReport()
	c.JSON(http.StatusOK, report)
}
