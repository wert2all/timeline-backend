package config

import "github.com/sakirsensoy/genv"

type appConfig struct {
	Port string
}

var AppConfig = &appConfig{
	Port: genv.Key("PORT").Default("8000").String(),
}
