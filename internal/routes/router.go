package routes

import (
	"github.com/Seasky89/go-report-api/internal/app"

	"github.com/gin-gonic/gin"
)

func SetupRoutes(app *app.App) *gin.Engine {
	router := gin.Default()
	router.GET("/report", app.ReportHandler.GenerateReport)
	router.GET("/report/users", app.UserHandler.FindAll)
	router.GET("/report/posts", app.PostHandler.FindAll)

	return router
}
