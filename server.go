package main

import (
	"context"
	"github.com/getsentry/sentry-go"
	"github.com/sakirsensoy/genv"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
	"log"
	"time"
	"timeline/backend/app"
	"timeline/backend/db"
	"timeline/backend/db/model"
	"timeline/backend/db/model/event"
	"timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
	eventRepository "timeline/backend/db/repository/event"
	timelineRepository "timeline/backend/db/repository/timeline"
	userRepository "timeline/backend/db/repository/user"
	"timeline/backend/graph/resolvers"
)

func main() {
	appConfig := readConfig()
	client := db.NewClient(appConfig.Postgres)
	defer client.Close()

	error := sentry.Init(sentry.ClientOptions{Dsn: appConfig.SentryDsn, Debug: true})
	if error != nil {
		log.Fatalf("sentry.Init: %s", error)
	}
	defer sentry.Flush(time.Second)

	ctx := context.Background()

	userModel := user.NewUserModel(userRepository.NewUserRepository(ctx, client))
	timelineModel := timeline.NewTimelineModel(timelineRepository.NewTimelineRepository(ctx, client))
	eventModel := event.NewEventModel(eventRepository.NewRepository(ctx, client))

	models := model.NewAllModels(userModel, timelineModel, eventModel)
	resolvers := resolvers.New(eventModel, userModel, timelineModel)
	app.NewApplication(app.NewAppState(models, appConfig, resolvers)).Start()
}

func readConfig() app.AppConfig {
	return app.AppConfig{
		Port: genv.Key("PORT").Default("8000").String(),
		CORS: app.CORS{
			Debug:         genv.Key("CORS_DEBUG").Default(false).Bool(),
			AllowedOrigin: genv.Key("CORS_ALLOWED_ORIGIN").String(),
		},
		Postgres: app.Postgres{
			Host:     genv.Key("POSTGRES_HOST").Default("localhost").String(),
			Port:     genv.Key("POSTGRES_PORT").Default(5432).Int(),
			User:     genv.Key("POSTGRES_USER").String(),
			Password: genv.Key("POSTGRES_PASSWORD").String(),
			Database: genv.Key("POSTGRES_DB").String(),
		},
		GoogleClintID: genv.Key("GOOGLE_CLIENT_ID").String(),
		SentryDsn:     genv.Key("SENTRY_DSN").String(),
	}
}
