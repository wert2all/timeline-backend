package di

import (
	"log"
	"strconv"
	"strings"

	"timeline/backend/app/config"
	"timeline/backend/ent"
	"timeline/backend/ent/migrate"
	"timeline/backend/lib/utils"

	"github.com/getsentry/sentry-go"
	_ "github.com/lib/pq"
	"golang.org/x/net/context"
)

func init() ServiceLocator {
	config := readConfig()
	applicationContext := context.Background()

	initSentry(config.Sentry.Dsn)

	client, err := newDBClient(applicationContext, config.Postgres)
	if err != nil {
		utils.F("Could not connect to database: %v", err)
	}
	locator := serviceLocator{
		config:  config,
		context: applicationContext,
		client:  client,
	}
	locator.repositoriesServiceLocator = newRepositoriesServiceLocator(locator)
	locator.modelsServiceLocator = newModelsServiceLocator(locator)
	locator.resolversServiceLocator = newResolversServiceLocator(locator)
	locator.middlewares = newMiddlewares(locator)
	return locator
}

func newDBClient(context context.Context, config config.Postgres) (*ent.Client, error) {
	client, err := ent.Open("postgres", createConnectionURL(config))
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(
		context,
		migrate.WithDropIndex(true),
		migrate.WithDropColumn(true)); err != nil {
		return nil, err
	}

	return client, nil
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

func initSentry(dsn string) {
	err := sentry.Init(sentry.ClientOptions{Dsn: dsn, Debug: true})
	if err != nil {
		log.Fatalf("sentry.Init: %s", err)
	}
}
