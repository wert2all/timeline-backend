package config

import (
	"github.com/sakirsensoy/genv"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
)

type CORS struct {
	AllowedOrigin string
	Debug         bool
}

type Postgres struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
}

type appConfig struct {
	Port          string
	CORS          CORS
	Postgres      Postgres
	GoogleClintID string
}

var AppConfig = &appConfig{
	Port: genv.Key("PORT").Default("8000").String(),
	CORS: CORS{
		Debug:         genv.Key("CORS_DEBUG").Default(false).Bool(),
		AllowedOrigin: genv.Key("CORS_ALLOWED_ORIGIN").String(),
	},
	Postgres: Postgres{
		Host:     genv.Key("POSTGRES_HOST").Default("localhost").String(),
		Port:     genv.Key("POSTGRES_PORT").Default(5432).Int(),
		User:     genv.Key("POSTGRES_USER").String(),
		Password: genv.Key("POSTGRES_PASSWORD").String(),
		Database: genv.Key("POSTGRES_DB").String(),
	},
	GoogleClintID: genv.Key("GOOGLE_CLIENT_ID").String(),
}
