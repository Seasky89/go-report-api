package main

import (
	"github.com/Seasky89/go-report-api/internal/app"
	"github.com/Seasky89/go-report-api/internal/routes"
)

func main() {
	app := app.NewApp()

	router := routes.SetupRoutes(app)
	if err := router.Run(":8080"); err != nil {
		panic(err)
	}
}
