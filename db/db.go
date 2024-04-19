package db

import (
	"context"
	"log"
	"strconv"
	"strings"
	"timeline/backend/ent"

	_ "github.com/lib/pq"
)

type PostgresConfig struct {
	Host, User, Password, Database string
	Port                           int
}

func CreateClient(connectionURL string) *ent.Client {
	client, err := ent.Open("postgres", connectionURL)
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func CreateConnectionURL(config PostgresConfig) string {
	var sb strings.Builder

	optionsMap := map[string]string{
		"host":     config.Host,
		"port":     strconv.Itoa(config.Port),
		"user":     config.User,
		"password": config.Password,
		"dbname":   config.Database,
		"sslmode":  "disable",
	}

	for key, val := range optionsMap {
		sb.WriteString(key + "=" + val + " ")
	}

	return sb.String()
}
