package di

import (
	"timeline/backend/ent"

	"golang.org/x/net/context"
)

type Postgres struct {
	Port     string `yaml:"port"`
	Host     string `yaml:"host"`
	Db       string `yaml:"db"`
	User     string `yaml:"user"`
	Password string `yaml:"password"`
}

type Config struct {
	App struct {
		Cors struct {
			Debug         bool   `yaml:"debug"`
			AllowedOrigin string `yaml:"allowedOrigin"`
		} `yaml:"cors"`
	} `yaml:"app"`

	Google struct {
		ClientId string `yaml:"clientId"`
	} `yaml:"google"`

	Postgres Postgres `yaml:"postgres"`

	Sentry struct {
		Dsn string `yaml:"dsn"`
	} `yaml:"sentry"`
}

func NewServiceLocator(config Config, context context.Context, client *ent.Client) ServiceLocator {
	locator := serviceLocator{config: config, context: context, client: client}
	locator.repositoriesServiceLocator = newRepositoriesServiceLocator(locator)
	locator.modelsServiceLocator = newModelsServiceLocator(locator)
	locator.resolversServiceLocator = newResolversServiceLocator(locator)
	locator.middlewares = newMiddlewares(locator)
	return locator
}
