package config

import (
	"flag"

	"github.com/getsentry/sentry-go"
)

type (
	Postgres struct {
		Port     int
		Host     string
		DB       string
		User     string
		Password string
	}

	Cors struct {
		AllowedOrigin string
	}

	ListenHost struct {
		Host string
		Port int
	}
	App struct {
		Cors        Cors
		Listen      ListenHost
		Development bool
	}

	Google struct {
		ClientID string
	}

	Previewly struct {
		API string
	}

	Config struct {
		App      App
		Postgres Postgres

		Google    Google
		Previewly Previewly
		Sentry    sentry.ClientOptions
	}
)

func NewConfig() Config {
	var (
		listenHostFlag     string
		listenPortFlag     int
		developmentFlag    bool
		postgresPort       int
		postgresHost       string
		postgresDB         string
		postgresUser       string
		postgresPassword   string
		sentryDsnFlag      string
		googleClientIDFlag string
		previewlyAPI       string
	)

	flag.StringVar(&listenHostFlag, "listen-host", "localhost", "Listen host")
	flag.IntVar(&listenPortFlag, "listen-port", 8000, "Listen port")

	flag.BoolVar(&developmentFlag, "development", false, "Development mode")

	flag.StringVar(&postgresHost, "postgres-host", "localhost", "Postgres host")
	flag.IntVar(&postgresPort, "postgres-port", 5432, "Postgres port")
	flag.StringVar(&postgresUser, "postgres-user", "timeline", "Postgres user")
	flag.StringVar(&postgresPassword, "postgres-password", "timeline", "Postgres password")
	flag.StringVar(&postgresDB, "postgres-db", "timeline", "Postgres DB")

	flag.StringVar(&googleClientIDFlag, "google-client-id", "", "Google client ID")
	flag.StringVar(&sentryDsnFlag, "sentry-dsn", "", "Sentry DSN")
	flag.StringVar(&previewlyAPI, "previewly-api", "https://api.previewly.top/graphql", "Previewly API URL")
	flag.Parse()

	config := Config{
		App: App{
			Cors: Cors{
				AllowedOrigin: "*",
			},
			Listen: ListenHost{
				Host: listenHostFlag,
				Port: listenPortFlag,
			},
			Development: developmentFlag,
		},
		Postgres: Postgres{
			Port:     postgresPort,
			Host:     postgresHost,
			DB:       postgresDB,
			User:     postgresUser,
			Password: postgresPassword,
		},

		Google: Google{
			ClientID: googleClientIDFlag,
		},
		Sentry: sentry.ClientOptions{
			Dsn:           sentryDsnFlag,
			EnableTracing: false,
		},

		Previewly: Previewly{
			API: previewlyAPI,
		},
	}

	return config
}
