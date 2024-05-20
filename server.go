package main

import (
	"context"
	"github.com/getsentry/sentry-go"
	"gopkg.in/yaml.v2"
	"log"
	"os"
	"time"
	"timeline/backend/app"
	"timeline/backend/db"
	"timeline/backend/db/model/event"
	"timeline/backend/db/model/tag"
	"timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
	eventRepository "timeline/backend/db/repository/event"
	tagRepository "timeline/backend/db/repository/tag"
	timelineRepository "timeline/backend/db/repository/timeline"
	userRepository "timeline/backend/db/repository/user"
	"timeline/backend/di"
)

func main() {
	appConfig := readConfig()
	client := db.NewClient(appConfig.Postgres)
	defer client.Close()

	err := sentry.Init(sentry.ClientOptions{Dsn: appConfig.Sentry.Dsn, Debug: true})
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
	f, err := os.Open("config.yaml")
	if err != nil {
		panic("cannot open config file")
	}
	defer f.Close()

	var cfg di.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic("cannot parse config ")
	}

	return cfg
}
