package di

import (
	"log"
	"os"
	"strings"
	"timeline/backend/ent"

	"github.com/getsentry/sentry-go"
	"golang.org/x/net/context"
	"gopkg.in/yaml.v2"
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

func Init() ServiceLocator {
	config := readConfig()
	applicationContext := context.Background()

	initSentry(config.Sentry.Dsn)

	locator := serviceLocator{
		config:  config,
		context: applicationContext,
		client:  newDBClient(applicationContext, config.Postgres),
	}
	locator.repositoriesServiceLocator = newRepositoriesServiceLocator(locator)
	locator.modelsServiceLocator = newModelsServiceLocator(locator)
	locator.resolversServiceLocator = newResolversServiceLocator(locator)
	locator.middlewares = newMiddlewares(locator)
	return locator
}

func newDBClient(context context.Context, config Postgres) *ent.Client {
	client, err := ent.Open("postgres", createConnectionURL(config))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	if err := client.Schema.Create(context); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func createConnectionURL(config Postgres) string {
	var sb strings.Builder

	optionsMap := map[string]string{
		"host":     config.Host,
		"port":     config.Port,
		"user":     config.User,
		"password": config.Password,
		"dbname":   config.Db,
		"sslmode":  "disable",
	}

	for key, val := range optionsMap {
		sb.WriteString(key + "=" + val + " ")
	}

	return sb.String()
}

func readConfig() Config {
	f, err := os.Open("config.yaml")
	if err != nil {
		panic("cannot open config file")
	}
	defer f.Close()

	var cfg Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		panic("cannot parse config ")
	}

	return cfg
}

func initSentry(dsn string) {
	err := sentry.Init(sentry.ClientOptions{Dsn: dsn, Debug: true})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}
