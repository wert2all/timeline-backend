package main

import (
	"context"
	"log"
	"time"
	"timeline/backend/app"
	"timeline/backend/db"
	"timeline/backend/db/model/event"
	"timeline/backend/db/model/tag"
	"timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
	eventRepository "timeline/backend/db/repository/event"
	timelineRepository "timeline/backend/db/repository/timeline"
	userRepository "timeline/backend/db/repository/user"
	tagRepository "timeline/backend/db/repository/tag"
	"timeline/backend/di"

	"github.com/getsentry/sentry-go"
	"github.com/sakirsensoy/genv"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
)

func main() {
	appConfig := readConfig()
	client := db.NewClient(appConfig.Postgres)
	defer client.Close()

	err := sentry.Init(sentry.ClientOptions{Dsn: appConfig.SentryDsn, Debug: true})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
	defer sentry.Flush(time.Second)

	ctx := context.Background()

	userModel := user.NewUserModel(userRepository.NewUserRepository(ctx, client))
	timelineModel := timeline.NewTimelineModel(timelineRepository.NewTimelineRepository(ctx, client))
	eventModel := event.NewEventModel(eventRepository.NewRepository(ctx, client))
	tagModel := tag.NewTagModel(tagRepository.NewRepository(ctx, client))

	models := di.NewAllModels(userModel, timelineModel, eventModel, tagModel)
	app.NewApplication(app.NewAppState(models, appConfig), di.NewServiceLocator(ctx, client)).Start()
}

func readConfig() di.Config {
	return di.Config{
		Port: genv.Key("PORT").Default("8000").String(),
		CORS: di.CORS{
			Debug:         genv.Key("CORS_DEBUG").Default(false).Bool(),
			AllowedOrigin: genv.Key("CORS_ALLOWED_ORIGIN").String(),
		},
		Postgres: di.Postgres{
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
