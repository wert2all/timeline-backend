package main

import (
	"context"
	"timeline/backend/app"
	"timeline/backend/config"
	"timeline/backend/db"
	"timeline/backend/db/model/user"
	userRepository "timeline/backend/db/repository/user"
)

func main() {
	client := db.CreateClient(
		db.CreateConnectionURL(
			db.PostgresConfig{
				Host:     config.AppConfig.Postgres.Host,
				Port:     config.AppConfig.Postgres.Port,
				User:     config.AppConfig.Postgres.User,
				Password: config.AppConfig.Postgres.Password,
				Database: config.AppConfig.Postgres.Database,
			}))

	defer client.Close()
	ctx := context.Background()

	userModel := user.NewUserModel(userRepository.NewUserRepository(ctx, client))

	app.Start(userModel)
}
