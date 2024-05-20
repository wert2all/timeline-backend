package db

import (
	"context"
	"log"
	"strings"
	"timeline/backend/di"
	"timeline/backend/ent"

	_ "github.com/lib/pq"
)

func NewClient(config di.Postgres) *ent.Client {
	client, err := ent.Open("postgres", createConnectionURL(config))
	if err != nil {
		log.Fatalf("failed opening connection to postgres: %v", err)
	}
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}

func createConnectionURL(config di.Postgres) string {
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
