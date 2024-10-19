package di

import (
	eventModel "timeline/backend/db/model/event"
	tagModel "timeline/backend/db/model/tag"
	timelineModel "timeline/backend/db/model/timeline"
	userModel "timeline/backend/db/model/user"
	"timeline/backend/db/repository/event"
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
)

func initOperationsResolvers() {
	initArgumentFactories()
	initValidators()
	initResolvers()
}

func initArgumentFactories() {
	initService(func() resolvers.AuthorizeArgumentFactory { return resolvers.AuthorizeArgumentFactory{} })
	initService(func() resolvers.AddTimelineArgumentFactory { return resolvers.AddTimelineArgumentFactory{} })
	initService(func() resolvers.AddEventArgumentFactory { return resolvers.AddEventArgumentFactory{} })
	initService(func() resolvers.DeleteEventArgumentFactory { return resolvers.DeleteEventArgumentFactory{} })
}

func initValidators() {
	initService(func(userModel userModel.UserModel) resolvers.Validator[resolvers.AuthorizeArguments, resolvers.ValidAuthorizeArguments] {
		return resolvers.NewAuthorizeValidator(userModel)
	})
	initService(func(userModel userModel.UserModel) resolvers.Validator[resolvers.AddTimelineArguments, resolvers.ValidAddTimelineArguments] {
		return resolvers.NewAddtimelineValidator(userModel)
	})
	initService(func(timelineModel timelineModel.UserTimeline) resolvers.Validator[resolvers.AddEventArguments, resolvers.ValidAddEventArguments] {
		return resolvers.NewAddEventValidator(timelineModel)
	})
	initService(func(userModel userModel.UserModel, eventModel eventModel.Model) resolvers.Validator[resolvers.DeleteEventArguments, resolvers.ValidDeleteEventArguments] {
		return resolvers.NewDeleteEventValidator(userModel, eventModel)
	})
}

func initResolvers() {
	initService(func(timelineModel timelineModel.UserTimeline) resolvers.Resolver[*model.User, resolvers.ValidAuthorizeArguments] {
		return resolvers.NewAutorizeResolver(timelineModel)
	})
	initService(func(timelineModel timelineModel.UserTimeline, userModel userModel.UserModel) resolvers.Resolver[*model.ShortUserTimeline, resolvers.ValidAddTimelineArguments] {
		return resolvers.NewAddTimelineResolver(userModel, timelineModel)
	})
	initService(func(timelineModel timelineModel.UserTimeline, eventModel eventModel.Model, tagModel tagModel.Model) resolvers.Resolver[*model.TimelineEvent, resolvers.ValidAddEventArguments] {
		return resolvers.NewAddEventResolver(eventModel, timelineModel, tagModel)
	})
	initService(func(eventRepository event.Repository) resolvers.Resolver[model.Status, resolvers.ValidDeleteEventArguments] {
		return resolvers.NewDeleteEventResolver(eventRepository)
	})
}
