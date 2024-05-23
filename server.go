package main

import (
	"timeline/backend/app"
	appRouter "timeline/backend/app/router"
	"timeline/backend/app/router/handler"
	"timeline/backend/app/router/route"
	"timeline/backend/di"
)

func main() {
	serviceLocator := di.Init()

	router := appRouter.NewRouterBuilder().
		SetMiddlewares(serviceLocator.Middlewares().Common()...).
		SetRoutes(route.NewGQLRoute(handler.NewGQLHandler(serviceLocator), serviceLocator.Middlewares().AuthMiddleware())).
		Build()

	application := app.NewApplication(router)
	application.Start()

	defer serviceLocator.Close()
}
