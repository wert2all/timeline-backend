package main

import (
	"timeline/backend/app"
	"timeline/backend/di"
)

func main() {
	serviceLocator := di.Init()

	application := app.NewApplication(serviceLocator)
	application.Start()

	defer application.Close()
}
