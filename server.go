package main

import (
	"context"
	"log"
	"os"
	"time"
	"timeline/backend/app"
	"timeline/backend/db"
	"timeline/backend/di"

	"github.com/getsentry/sentry-go"
	"gopkg.in/yaml.v2"
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

	app.NewApplication(app.NewAppState(appConfig), di.NewServiceLocator(context.Background(), client)).Start()
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
