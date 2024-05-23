package main

import (
	"timeline/backend/app"
	appRouter "timeline/backend/app/router"
	"timeline/backend/app/router/route"
	"timeline/backend/di"
	"timeline/backend/graph"
)

func main() {
	serviceLocator := di.Init()

	gqlConfig := route.NewGQLConfig(&graph.Resolver{ServiceLocator: serviceLocator})
	gqlRoute := route.NewGQLRoute(gqlConfig, serviceLocator.Middlewares().AuthMiddleware())

	router := appRouter.NewRouterBuilder().
		SetMiddlewares(serviceLocator.Middlewares().Common()...).
		SetRoutes(gqlRoute).
		Build()

	application := app.NewApplication(router)
	application.Start()

	defer serviceLocator.Close()
}
