package di

import (
	"context"
	"log/slog"
	"time"

	"timeline/backend/app"
	"timeline/backend/app/config"
	middlewares "timeline/backend/app/middleware"

	"timeline/backend/ent"
	"timeline/backend/lib/log"
	"timeline/backend/lib/utils"

	userModel "timeline/backend/db/model/user"
	"timeline/backend/db/repository/user"

	"github.com/getsentry/sentry-go"
	"github.com/golobby/container/v3"
)

func InitContainer(config config.Config, appContext context.Context) {
	initVoidServices(config)

	initService(func() context.Context { return appContext })
	initService(func(context context.Context) (*ent.Client, error) {
		return createClient(context, config.Postgres)
	})

	initRepositories()
	initModels()

	initOperationsResolvers()

	initService(func(repository user.Repository) userModel.Authorize { return userModel.NewUserModel(repository) })

	initApplication(config)
}

func initApplication(config config.Config) {
	initService(func(entClient *ent.Client, userModel userModel.Authorize) app.Application {
		return app.NewApplication(
			newRouter(middlewares.NewMiddlewares(userModel, config.Google.ClientID)),
			config.App.Listen,
			func() {
				entClient.Close()
				sentry.Flush(time.Second)
			})
	})
}

func initVoidServices(config config.Config) {
	initSentry(config.Sentry)
	initLogger()
}

func initSentry(options sentry.ClientOptions) {
	err := sentry.Init(options)
	if err != nil {
		utils.F("Sentry initialization failed: %v\n", err)
	}
}

func initLogger() {
	logger := slog.New(log.NewHandler(nil))
	slog.SetDefault(logger)
}

func initService(resolver interface{}) {
	err := container.Singleton(resolver)
	if err != nil {
		utils.F("Couldnt inititalize service: %v", err)
	}
}
