package main

import (
	"context"
	"timeline/backend/app"
	appContext "timeline/backend/app/context"
	"timeline/backend/db"
	"timeline/backend/db/model/user"
	userRepository "timeline/backend/db/repository/user"

	"github.com/sakirsensoy/genv"
	_ "github.com/sakirsensoy/genv/dotenv/autoload"
)

func main() {
	appConfig := readConfig()
	client := db.NewClient(appConfig.Postgres)
	defer client.Close()

	ctx := context.Background()

	userModel := user.NewUserModel(userRepository.NewUserRepository(ctx, client))

	models := appContext.NewModels(userModel)

	app.NewApplication(app.NewAppState(appContext.NewAppContext(ctx, models), appConfig)).Start()
}

func readConfig() app.AppConfig {
	return app.AppConfig{
		Port: genv.Key("PORT").Default("8000").String(),
		CORS: app.CORS{
			Debug:         genv.Key("CORS_DEBUG").Default(false).Bool(),
			AllowedOrigin: genv.Key("CORS_ALLOWED_ORIGIN").String(),
		},
		Postgres: app.Postgres{
			Host:     genv.Key("POSTGRES_HOST").Default("localhost").String(),
			Port:     genv.Key("POSTGRES_PORT").Default(5432).Int(),
			User:     genv.Key("POSTGRES_USER").String(),
			Password: genv.Key("POSTGRES_PASSWORD").String(),
			Database: genv.Key("POSTGRES_DB").String(),
		},
		GoogleClintID: genv.Key("GOOGLE_CLIENT_ID").String(),
	}
}
