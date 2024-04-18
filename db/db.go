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
	Host     string
	Port     int
	User     string
	Password string
	Database string
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
	sb.WriteString("host=")
	sb.WriteString(config.Host)

	sb.WriteString(" port=")
	sb.WriteString(strconv.Itoa(config.Port))

	sb.WriteString(" user=")
	sb.WriteString(config.User)
	sb.WriteString(" password=")
	sb.WriteString(config.Password)

	sb.WriteString(" dbname=")
	sb.WriteString(config.Database)

	sb.WriteString(" sslmode=disable")

	return sb.String()
}
