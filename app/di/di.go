package di

import (
	"context"
	"log/slog"
	"net/http"
	"time"

	"timeline/backend/app"
	"timeline/backend/app/config"
	middlewares "timeline/backend/app/middleware"
	"timeline/backend/graph/client/previewly"

	"timeline/backend/ent"
	"timeline/backend/lib/log"
	"timeline/backend/lib/utils"

	userModel "timeline/backend/db/model/user"
	"timeline/backend/db/repository/account"
	"timeline/backend/db/repository/user"

	"github.com/Khan/genqlient/graphql"
	"github.com/getsentry/sentry-go"
	"github.com/golobby/container/v3"

	domainUser "timeline/backend/domain/user"
)

func InitContainer(config config.Config, appContext context.Context) {
	initVoidServices(config)

	initService(func() context.Context { return appContext })
	initService(func(context context.Context) (*ent.Client, error) {
		return createClient(context, config.Postgres)
	})

	initService(func() graphql.Client { return graphql.NewClient(config.Previewly.API, http.DefaultClient) })
	initService(func(client graphql.Client) previewly.MutationClient { return previewly.NewMutationClient(client) })

	initRepositories()
	initModels()

	initService(
		func(userRepository user.Repository,
			accountRepository account.Repository,
			mutationClient previewly.MutationClient,
		) userModel.Authorize {
			return userModel.NewUserModel(userRepository, accountRepository, mutationClient)
		},
	)
	initService(func(userModel userModel.Authorize) domainUser.UserExtractor {
		return domainUser.NewUserExtractor(config.Google.ClientID, userModel)
	})

	initOperationsResolvers()

	initApplication(config)
}

func initApplication(config config.Config) {
	initService(func(entClient *ent.Client) app.Application {
		return app.NewApplication(
			newRouter(middlewares.NewMiddlewares(), config.App.Development),
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
