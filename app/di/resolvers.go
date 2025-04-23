package di

import (
	accountModel "timeline/backend/db/model/account"
	eventModel "timeline/backend/db/model/event"
	"timeline/backend/db/model/settings"
	tagModel "timeline/backend/db/model/tag"
	timelineModel "timeline/backend/db/model/timeline"
	userModel "timeline/backend/db/model/user"
	"timeline/backend/db/repository/event"
	domainUser "timeline/backend/domain/user"
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
	settingsResolver "timeline/backend/graph/resolvers/mutation/account/settings"
	eventValidator "timeline/backend/graph/resolvers/mutation/event"
	myAccountTimelines "timeline/backend/graph/resolvers/query/timeline"

	addAccountResolver "timeline/backend/graph/resolvers/mutation/account/add"
	saveAccountResolver "timeline/backend/graph/resolvers/mutation/account/save"
	getEventResolver "timeline/backend/graph/resolvers/query/getevent"
	getEventsResolver "timeline/backend/graph/resolvers/query/getevents"
	getTimelineResolver "timeline/backend/graph/resolvers/query/gettimeline"
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
	initService(func() settingsResolver.SaveSettingsArgumentFactory {
		return settingsResolver.SaveSettingsArgumentFactory{}
	})
	initService(func() getEventsResolver.GetCursorEventsArgumentFactory {
		return getEventsResolver.GetCursorEventsArgumentFactory{}
	})
	initService(func() saveAccountResolver.SaveAccountArgumentsFactory {
		return saveAccountResolver.SaveAccountArgumentsFactory{}
	})
	initService(func() addAccountResolver.AddAccountArgumentFactory {
		return addAccountResolver.AddAccountArgumentFactory{}
	})
	initService(func() getTimelineResolver.GetTimelineArgumentFactory {
		return getTimelineResolver.GetTimelineArgumentFactory{}
	})
	initService(func() getEventResolver.GetEventArgumentFactory { return getEventResolver.GetEventArgumentFactory{} })
}

func initValidators() {
	initService(func(timeline timelineModel.Timeline, userExtractor domainUser.UserExtractor) eventValidator.BaseValidator {
		return eventValidator.NewBaseValidator(timeline, userExtractor)
	})

	initService(func(userModel userModel.UserModel, userExtractor domainUser.UserExtractor) resolvers.Validator[resolvers.AuthorizeArguments, resolvers.ValidAuthorizeArguments] {
		return resolvers.NewAuthorizeValidator(userModel, userExtractor)
	})
	initService(func(userModel userModel.UserModel, userExtractor domainUser.UserExtractor) resolvers.Validator[resolvers.AddTimelineArguments, resolvers.ValidAddTimelineArguments] {
		return resolvers.NewAddtimelineValidator(userModel, userExtractor)
	})
	initService(func(baseValidator eventValidator.BaseValidator) resolvers.Validator[resolvers.AddEventArguments, resolvers.ValidAddEventArguments] {
		return resolvers.NewAddEventValidator(baseValidator)
	})
	initService(func(baseValidator eventValidator.BaseValidator, eventModel eventModel.Model, timelineModel timelineModel.Timeline) resolvers.Validator[resolvers.EditEventArguments, resolvers.ValidEditEventArguments] {
		return resolvers.NewEditEventValidator(baseValidator, eventModel, timelineModel)
	})
	initService(func(userModel userModel.UserModel, eventModel eventModel.Model, userExtractor domainUser.UserExtractor) resolvers.Validator[resolvers.DeleteEventArguments, resolvers.ValidDeleteEventArguments] {
		return resolvers.NewDeleteEventValidator(userModel, eventModel, userExtractor)
	})
	initService(func(userModel userModel.UserModel, userExtractor domainUser.UserExtractor) resolvers.Validator[settingsResolver.SaveSettingsArguments, settingsResolver.ValidSaveSettingsArguments] {
		return settingsResolver.NewSaveSettingsValidator(userModel, userExtractor)
	})
	initService(func(userModel userModel.UserModel, timelineModel timelineModel.Timeline, userExtractor domainUser.UserExtractor) resolvers.Validator[getEventsResolver.GetCursorEventsArguments, getEventsResolver.ValidGetCursorEventsArguments] {
		return getEventsResolver.NewValidator(userModel, timelineModel, userExtractor)
	})
	initService(func(userModel userModel.UserModel, userExtractor domainUser.UserExtractor) resolvers.Validator[saveAccountResolver.SaveAccountArguments, saveAccountResolver.ValidSaveAccountArguments] {
		return saveAccountResolver.NewValidator(userModel, userExtractor)
	})
	initService(func(userExtractor domainUser.UserExtractor, userModel userModel.UserModel) resolvers.Validator[addAccountResolver.AddAccountArguments, addAccountResolver.ValidAddAccountArguments] {
		return addAccountResolver.NewValidator(userExtractor, userModel)
	})
	initService(func() resolvers.Validator[getTimelineResolver.GetTimelineArguments, getTimelineResolver.ValidGetTimelineArguments] {
		return getTimelineResolver.NewValidator()
	})
	initService(func(userModel userModel.UserModel, userExtractor domainUser.UserExtractor) resolvers.Validator[getEventResolver.GetEventArguments, getEventResolver.ValidGetEventArguments] {
		return getEventResolver.NewValidator(userModel, userExtractor)
	})
}

func initResolvers() {
	initService(func(timelineModel timelineModel.Timeline, userModel userModel.UserModel, settings settings.Model) resolvers.Resolver[*model.User, resolvers.ValidAuthorizeArguments] {
		return resolvers.NewAutorizeResolver(timelineModel, userModel, settings)
	})
	initService(func(timelineModel timelineModel.Timeline, userModel userModel.UserModel) resolvers.Resolver[*model.Timeline, resolvers.ValidAddTimelineArguments] {
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
	initService(func(timelineModel timelineModel.Timeline, userModel userModel.UserModel) myAccountTimelines.Resolver {
		return myAccountTimelines.NewMyAccountTimelinesResolver(timelineModel, userModel)
	})
	initService(func(settingsModel settings.Model) resolvers.Resolver[model.Status, settingsResolver.ValidSaveSettingsArguments] {
		return settingsResolver.NewSaveSettingsResolver(settingsModel)
	})
	initService(func(eventModel eventModel.Model, tagModel tagModel.Model) resolvers.Resolver[*model.TimelineEvents, getEventsResolver.ValidGetCursorEventsArguments] {
		return getEventsResolver.NewResolver(eventModel, tagModel)
	})
	initService(func(userModel userModel.UserModel, settingsModel settings.Model) resolvers.Resolver[*model.ShortAccount, saveAccountResolver.ValidSaveAccountArguments] {
		return saveAccountResolver.NewResolver(userModel, settingsModel)
	})
	initService(func(accountModel accountModel.Model, settingsModel settings.Model) resolvers.Resolver[*model.ShortAccount, addAccountResolver.ValidAddAccountArguments] {
		return addAccountResolver.NewResolver(accountModel, settingsModel)
	})
	initService(func(timelineModel timelineModel.Timeline, settingsModel settings.Model) resolvers.Resolver[*model.Timeline, getTimelineResolver.ValidGetTimelineArguments] {
		return getTimelineResolver.NewResolver(timelineModel, settingsModel)
	})
	initService(func(eventModel eventModel.Model, tagModel tagModel.Model) resolvers.Resolver[*model.TimelineEvent, getEventResolver.ValidGetEventArguments] {
		return getEventResolver.NewResolver(eventModel, tagModel)
	})
}
