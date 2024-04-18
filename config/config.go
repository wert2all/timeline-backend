package config

import (
	"github.com/sakirsensoy/genv"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
)

type CORS struct {
	AllowedOrigin string
	Debug         bool
}
type appConfig struct {
	Port          string
	CORS          CORS
	GoogleClintID string
}

var AppConfig = &appConfig{
	Port: genv.Key("PORT").Default("8000").String(),
	CORS: CORS{
		Debug:         genv.Key("CORS_DEBUG").Default(false).Bool(),
		AllowedOrigin: genv.Key("CORS_ALLOWED_ORIGIN").String(),
	},
	GoogleClintID: genv.Key("GOOGLE_CLIENT_ID").String(),
}
