package app

import (
	"github.com/Seasky89/go-report-api/internal/bootstrap"
	"github.com/Seasky89/go-report-api/internal/handlers"
	"github.com/Seasky89/go-report-api/internal/services"
)

type App struct {
	UserHandler   *handlers.UserHandler
	PostHandler   *handlers.PostHandler
	ReportHandler *handlers.ReportHandler
}

func NewApp() *App {
	memoryStore := bootstrap.InitMemoryStore()

	reportService := services.NewReportService(memoryStore)
	postService := services.NewPostService(memoryStore)
	userService := services.NewUserService(memoryStore)

	return &App{
		UserHandler:   handlers.NewUserHandler(userService),
		PostHandler:   handlers.NewPostHandler(postService),
		ReportHandler: handlers.NewReportHandler(reportService),
	}
}
