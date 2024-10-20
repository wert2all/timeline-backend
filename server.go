package main

import (
	"context"

	"timeline/backend/app"
	"timeline/backend/app/config"
	"timeline/backend/app/di"
	"timeline/backend/lib/utils"

	"github.com/golobby/container/v3"
)

func main() {
	di.InitContainer(config.NewConfig(), context.Background())

	var application app.Application
	err := container.Resolve(&application)
	if err != nil {
		utils.F("Coulnd not create application: %v", err)
	}

	application.Start()
	defer application.Closer()()
}
