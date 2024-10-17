package di

import (
	"log"
	"os"
	"strconv"
	"strings"

	"timeline/backend/app/config"
	"timeline/backend/ent"
	"timeline/backend/lib/dumper"

	"github.com/getsentry/sentry-go"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
	"gopkg.in/yaml.v2"
)

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

func newDBClient(context context.Context, config config.Postgres) *ent.Client {
	client, err := ent.Open("postgres", createConnectionURL(config))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	if err := client.Schema.Create(context); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func createConnectionURL(config config.Postgres) string {
	var sb strings.Builder

	optionsMap := map[string]string{
		"host":     config.Host,
		"port":     strconv.Itoa(config.Port),
		"user":     config.User,
		"password": config.Password,
		"dbname":   config.DB,
		"sslmode":  "disable",
	}

	for key, val := range optionsMap {
		sb.WriteString(key + "=" + val + " ")
	}

	return sb.String()
}

func readConfig() config.Config {
	f, err := os.Open("config.yaml")
	if err != nil {
		panic("cannot open config file")
	}
	defer f.Close()

	var cfg config.Config
	decoder := yaml.NewDecoder(f)
	err = decoder.Decode(&cfg)
	if err != nil {
		dumper.D(err, cfg)
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
