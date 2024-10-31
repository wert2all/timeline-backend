package di

import (
	"context"
	"strconv"
	"strings"

	"timeline/backend/app/config"
	eventModel "timeline/backend/db/model/event"
	tagModel "timeline/backend/db/model/tag"
	timelineModel "timeline/backend/db/model/timeline"
	userModel "timeline/backend/db/model/user"
	"timeline/backend/db/repository/event"
	"timeline/backend/db/repository/tag"
	"timeline/backend/db/repository/timeline"
	"timeline/backend/db/repository/user"
	"timeline/backend/ent"
	"timeline/backend/ent/migrate"

	_ "github.com/lib/pq"
)

func initRepositories() {
	initService(func(ctx context.Context, client *ent.Client) user.Repository {
		return user.NewUserRepository(ctx, client)
	})
	initService(func(ctx context.Context, client *ent.Client) timeline.Repository {
		return timeline.NewTimelineRepository(ctx, client)
	})
	initService(func(ctx context.Context, client *ent.Client) event.Repository {
		return event.NewRepository(ctx, client)
	})
	initService(func(ctx context.Context, client *ent.Client) tag.Repository {
		return tag.NewRepository(ctx, client)
	})
}

func initModels() {
	initService(func(repository user.Repository) userModel.UserModel {
		return userModel.NewUserModel(repository)
	})

	initService(func(repository tag.Repository) tagModel.Model {
		return tagModel.NewTagModel(repository)
	})
	initService(func(repository timeline.Repository) timelineModel.UserTimeline {
		return timelineModel.NewTimelineModel(repository)
	})
	initService(func(repository event.Repository, timelineModel timelineModel.UserTimeline, tagModel tagModel.Model) eventModel.Model {
		return eventModel.NewEventModel(repository, tagModel, timelineModel)
	})
}

func createClient(context context.Context, config config.Postgres) (*ent.Client, error) {
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
