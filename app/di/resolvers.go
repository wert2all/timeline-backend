package di

import (
	eventModel "timeline/backend/db/model/event"
	"timeline/backend/db/model/settings"
	tagModel "timeline/backend/db/model/tag"
	"timeline/backend/db/model/timeline"
	timelineModel "timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
	userModel "timeline/backend/db/model/user"
	"timeline/backend/db/repository/event"
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
	eventValidator "timeline/backend/graph/resolvers/mutation/event"
	myAccountTimelines "timeline/backend/graph/resolvers/query/timeline"
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
	initService(func() resolvers.EditEventArgumentFactory { return resolvers.EditEventArgumentFactory{} })
	initService(func() resolvers.DeleteEventArgumentFactory { return resolvers.DeleteEventArgumentFactory{} })
}

func initValidators() {
	initService(func(timeline timelineModel.Timeline) eventValidator.BaseValidator {
		return eventValidator.NewBaseValidator(timeline)
	})

	initService(func(userModel userModel.UserModel) resolvers.Validator[resolvers.AuthorizeArguments, resolvers.ValidAuthorizeArguments] {
		return resolvers.NewAuthorizeValidator(userModel)
	})
	initService(func(userModel userModel.UserModel) resolvers.Validator[resolvers.AddTimelineArguments, resolvers.ValidAddTimelineArguments] {
		return resolvers.NewAddtimelineValidator(userModel)
	})
	initService(func(baseValidator eventValidator.BaseValidator) resolvers.Validator[resolvers.AddEventArguments, resolvers.ValidAddEventArguments] {
		return resolvers.NewAddEventValidator(baseValidator)
	})
	initService(func(baseValidator eventValidator.BaseValidator, eventModel eventModel.Model, timelineModel timeline.Timeline) resolvers.Validator[resolvers.EditEventArguments, resolvers.ValidEditEventArguments] {
		return resolvers.NewEditEventValidator(baseValidator, eventModel, timelineModel)
	})
	initService(func(userModel userModel.UserModel, eventModel eventModel.Model) resolvers.Validator[resolvers.DeleteEventArguments, resolvers.ValidDeleteEventArguments] {
		return resolvers.NewDeleteEventValidator(userModel, eventModel)
	})
}

func initResolvers() {
	initService(func(timelineModel timelineModel.Timeline, userModel user.UserModel, settings settings.Model) resolvers.Resolver[*model.User, resolvers.ValidAuthorizeArguments] {
		return resolvers.NewAutorizeResolver(timelineModel, userModel, settings)
	})
	initService(func(timelineModel timelineModel.Timeline, userModel userModel.UserModel) resolvers.Resolver[*model.ShortTimeline, resolvers.ValidAddTimelineArguments] {
		return resolvers.NewAddTimelineResolver(userModel, timelineModel)
	})
	initService(func(timelineModel timelineModel.Timeline, eventModel eventModel.Model, tagModel tagModel.Model) resolvers.Resolver[*model.TimelineEvent, resolvers.ValidAddEventArguments] {
		return resolvers.NewAddEventResolver(eventModel, timelineModel, tagModel)
	})
	initService(func(eventModel eventModel.Model) resolvers.Resolver[*model.TimelineEvent, resolvers.ValidEditEventArguments] {
		return resolvers.NewEditEventResolver(eventModel)
	})
	initService(func(eventRepository event.Repository) resolvers.Resolver[model.Status, resolvers.ValidDeleteEventArguments] {
		return resolvers.NewDeleteEventResolver(eventRepository)
	})

	initService(func(timelineModel timelineModel.Timeline, userModel user.UserModel) myAccountTimelines.Resolver {
		return myAccountTimelines.NewMyAccountTimelinesResolver(timelineModel, userModel)
	})
}
