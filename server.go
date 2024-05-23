package main

import (
	"timeline/backend/app"
	"timeline/backend/app/router"
	"timeline/backend/di"
)

func main() {
	serviceLocator := di.Init()

	application := app.NewApplication(router.NewRouterFactory(serviceLocator).Create())
	application.Start()

	defer serviceLocator.Close()
}
