package main

import (
	"context"
	"timeline/backend/app"
	"timeline/backend/db"
	"timeline/backend/db/model/user"
	userRepository "timeline/backend/db/repository/user"
)

func main() {
	state := app.NewAppState(app.ReadConfig())

	client := db.CreateClient(db.CreateConnectionURL(state.Config.Postgres))

	defer client.Close()
	ctx := context.Background()

	userModel := user.NewUserModel(userRepository.NewUserRepository(ctx, client))
	app := app.NewApplication(state, userModel)
	app.Start()
}
