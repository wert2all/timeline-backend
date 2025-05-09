package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.70

import (
	"context"
	"fmt"
	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/account"
	"timeline/backend/db/model/settings"
	"timeline/backend/db/model/timeline"
	domainUser "timeline/backend/domain/user"
	"timeline/backend/graph/convert"
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
	addAccountResolver "timeline/backend/graph/resolvers/mutation/account/add"
	saveAccountResolver "timeline/backend/graph/resolvers/mutation/account/save"
	settingsResolver "timeline/backend/graph/resolvers/mutation/account/settings"
	getEventResover "timeline/backend/graph/resolvers/query/getevent"
	getEventsResolver "timeline/backend/graph/resolvers/query/getevents"
	getTimelineResover "timeline/backend/graph/resolvers/query/gettimeline"
	myAccountTimelines "timeline/backend/graph/resolvers/query/timeline"
	enumvalues "timeline/backend/lib/enum-values"
	"timeline/backend/lib/utils"

	container "github.com/golobby/container/v3"
	"github.com/xorcare/pointer"
)

// Authorize is the resolver for the authorize field.
func (r *mutationResolver) Authorize(ctx context.Context) (*model.User, error) {
	var factory resolvers.AuthorizeArgumentFactory
	var validator resolvers.Validator[resolvers.AuthorizeArguments, resolvers.ValidAuthorizeArguments]
	var resolver resolvers.Resolver[*model.User, resolvers.ValidAuthorizeArguments]

	errFactoryResolve := container.Resolve(&factory)
	if errFactoryResolve != nil {
		utils.F("Couldnt resolve Authorize factory: %v", errFactoryResolve)
		return nil, errFactoryResolve
	}

	errValidatorResolve := container.Resolve(&validator)
	if errValidatorResolve != nil {
		utils.F("Couldnt resolve Authorize validator: %v", errValidatorResolve)
		return nil, errValidatorResolve
	}

	errResolverResolve := container.Resolve(&resolver)
	if errResolverResolve != nil {
		utils.F("Couldnt resolve Authorize resolver: %v", errResolverResolve)
		return nil, errResolverResolve
	}

	return resolvers.Resolve(ctx, factory.New(), validator, resolver)
}

// AddTimeline is the resolver for the addTimeline field.
func (r *mutationResolver) AddTimeline(ctx context.Context, timeline *model.AddTimeline) (*model.Timeline, error) {
	var factory resolvers.AddTimelineArgumentFactory
	var validator resolvers.Validator[resolvers.AddTimelineArguments, resolvers.ValidAddTimelineArguments]
	var resolver resolvers.Resolver[*model.Timeline, resolvers.ValidAddTimelineArguments]

	errFactoryResolve := container.Resolve(&factory)
	if errFactoryResolve != nil {
		utils.F("Couldnt resolve AddTimeline factory: %v", errFactoryResolve)
		return nil, errFactoryResolve
	}

	errValidatorResolve := container.Resolve(&validator)
	if errValidatorResolve != nil {
		utils.F("Couldnt resolve AddTimeline validator: %v", errValidatorResolve)
		return nil, errValidatorResolve
	}

	errResolverResolve := container.Resolve(&resolver)
	if errResolverResolve != nil {
		utils.F("Couldnt resolve AddTimeline resolver: %v", errResolverResolve)
		return nil, errResolverResolve
	}

	return resolvers.Resolve(ctx, factory.New(timeline), validator, resolver)
}

// AddEvent is the resolver for the addEvent field.
func (r *mutationResolver) AddEvent(ctx context.Context, event model.TimelineEventInput) (*model.TimelineEvent, error) {
	var factory resolvers.AddEventArgumentFactory
	var validator resolvers.Validator[resolvers.AddEventArguments, resolvers.ValidAddEventArguments]
	var resolver resolvers.Resolver[*model.TimelineEvent, resolvers.ValidAddEventArguments]

	errFactoryResolve := container.Resolve(&factory)
	if errFactoryResolve != nil {
		utils.F("Couldnt resolve AddEvent factory: %v", errFactoryResolve)
		return nil, errFactoryResolve
	}

	errValidatorResolve := container.Resolve(&validator)
	if errValidatorResolve != nil {
		utils.F("Couldnt resolve AddEvent validator: %v", errValidatorResolve)
		return nil, errValidatorResolve
	}

	errResolverResolve := container.Resolve(&resolver)
	if errResolverResolve != nil {
		utils.F("Couldnt resolve AddEvent resolver: %v", errResolverResolve)
		return nil, errResolverResolve
	}

	return resolvers.Resolve(ctx, factory.New(event), validator, resolver)
}

// EditEvent is the resolver for the editEvent field.
func (r *mutationResolver) EditEvent(ctx context.Context, event model.ExistTimelineEventInput) (*model.TimelineEvent, error) {
	var factory resolvers.EditEventArgumentFactory
	var validator resolvers.Validator[resolvers.EditEventArguments, resolvers.ValidEditEventArguments]
	var resolver resolvers.Resolver[*model.TimelineEvent, resolvers.ValidEditEventArguments]

	errFactoryResolve := container.Resolve(&factory)
	if errFactoryResolve != nil {
		utils.F("Couldnt resolve EditEvent factory: %v", errFactoryResolve)
		return nil, errFactoryResolve
	}

	errValidatorResolve := container.Resolve(&validator)
	if errValidatorResolve != nil {
		utils.F("Couldnt resolve EditEvent validator: %v", errValidatorResolve)
		return nil, errValidatorResolve
	}

	errResolverResolve := container.Resolve(&resolver)
	if errResolverResolve != nil {
		utils.F("Couldnt resolve EditEvent resolver: %v", errResolverResolve)
		return nil, errResolverResolve
	}

	return resolvers.Resolve(ctx, factory.New(event), validator, resolver)
}

// DeleteEvent is the resolver for the deleteEvent field.
func (r *mutationResolver) DeleteEvent(ctx context.Context, eventID int) (model.Status, error) {
	var factory resolvers.DeleteEventArgumentFactory
	var validator resolvers.Validator[resolvers.DeleteEventArguments, resolvers.ValidDeleteEventArguments]
	var resolver resolvers.Resolver[model.Status, resolvers.ValidDeleteEventArguments]

	errFactoryResolve := container.Resolve(&factory)
	if errFactoryResolve != nil {
		utils.F("Couldnt resolve DeleteEvent factory: %v", errFactoryResolve)
		return model.StatusError, errFactoryResolve
	}

	errValidatorResolve := container.Resolve(&validator)
	if errValidatorResolve != nil {
		utils.F("Couldnt resolve DeleteEvent validator: %v", errValidatorResolve)
		return model.StatusError, errValidatorResolve
	}

	errResolverResolve := container.Resolve(&resolver)
	if errResolverResolve != nil {
		utils.F("Couldnt resolve DeleteEvent resolver: %v", errResolverResolve)
		return model.StatusError, errResolverResolve
	}

	return resolvers.Resolve(ctx, factory.New(eventID), validator, resolver)
}

// AddAccount is the resolver for the addAccount field.
func (r *mutationResolver) AddAccount(ctx context.Context, name string) (*model.ShortAccount, error) {
	var factory addAccountResolver.AddAccountArgumentFactory
	var validator resolvers.Validator[addAccountResolver.AddAccountArguments, addAccountResolver.ValidAddAccountArguments]
	var resolver resolvers.Resolver[*model.ShortAccount, addAccountResolver.ValidAddAccountArguments]

	errFactoryResolve := container.Resolve(&factory)
	if errFactoryResolve != nil {
		utils.F("Couldnt resolve AddAccount factory: %v", errFactoryResolve)
		return nil, errFactoryResolve
	}

	errValidatorResolve := container.Resolve(&validator)
	if errValidatorResolve != nil {
		utils.F("Couldnt resolve AddAccount validator: %v", errValidatorResolve)
		return nil, errValidatorResolve
	}

	errResolverResolve := container.Resolve(&resolver)
	if errResolverResolve != nil {
		utils.F("Couldnt resolve AddAccount resolver: %v", errResolverResolve)
		return nil, errResolverResolve
	}

	return resolvers.Resolve(ctx, factory.New(name), validator, resolver)
}

// SaveAccount is the resolver for the saveAccount field.
func (r *mutationResolver) SaveAccount(ctx context.Context, accountID int, account model.SaveAccountInput) (*model.ShortAccount, error) {
	var factory saveAccountResolver.SaveAccountArgumentsFactory
	var validator resolvers.Validator[saveAccountResolver.SaveAccountArguments, saveAccountResolver.ValidSaveAccountArguments]
	var resolver resolvers.Resolver[*model.ShortAccount, saveAccountResolver.ValidSaveAccountArguments]

	errFactoryResolve := container.Resolve(&factory)
	if errFactoryResolve != nil {
		utils.F("Couldnt resolve SaveAccount factory: %v", errFactoryResolve)
		return nil, errFactoryResolve
	}

	errValidatorResolve := container.Resolve(&validator)
	if errValidatorResolve != nil {
		utils.F("Couldnt resolve SaveAccount validator: %v", errValidatorResolve)
		return nil, errValidatorResolve
	}

	errResolverResolve := container.Resolve(&resolver)
	if errResolverResolve != nil {
		utils.F("Couldnt resolve SaveAccount resolver: %v", errResolverResolve)
		return nil, errResolverResolve
	}

	return resolvers.Resolve(ctx, factory.New(accountID, account), validator, resolver)
}

// SaveSettings is the resolver for the saveSettings field.
func (r *mutationResolver) SaveSettings(ctx context.Context, accountID int, settings []*model.AccountSettingInput) (model.Status, error) {
	var factory settingsResolver.SaveSettingsArgumentFactory
	var validator resolvers.Validator[settingsResolver.SaveSettingsArguments, settingsResolver.ValidSaveSettingsArguments]
	var resolver resolvers.Resolver[model.Status, settingsResolver.ValidSaveSettingsArguments]

	errFactoryResolve := container.Resolve(&factory)
	if errFactoryResolve != nil {
		utils.F("Couldnt resolve SaveSettings factory: %v", errFactoryResolve)
		return model.StatusError, errFactoryResolve
	}

	errValidatorResolve := container.Resolve(&validator)
	if errValidatorResolve != nil {
		utils.F("Couldnt resolve SaveSettings validator: %v", errValidatorResolve)
		return model.StatusError, errValidatorResolve
	}

	errResolverResolve := container.Resolve(&resolver)
	if errResolverResolve != nil {
		utils.F("Couldnt resolve SaveSettings resolver: %v", errResolverResolve)
		return model.StatusError, errResolverResolve
	}

	return resolvers.Resolve(ctx, factory.New(accountID, settings), validator, resolver)
}

// TimelineCursorEvents is the resolver for the timelineCursorEvents field.
func (r *queryResolver) TimelineCursorEvents(ctx context.Context, accountID *int, timelineID int, limit *model.Limit, cursor *string) (*model.TimelineEvents, error) {
	var factory getEventsResolver.GetCursorEventsArgumentFactory
	var validator resolvers.Validator[getEventsResolver.GetCursorEventsArguments, getEventsResolver.ValidGetCursorEventsArguments]
	var resolver resolvers.Resolver[*model.TimelineEvents, getEventsResolver.ValidGetCursorEventsArguments]

	errFactoryResolve := container.Resolve(&factory)
	if errFactoryResolve != nil {
		utils.F("Couldnt resolve GetCursorEvents factory: %v", errFactoryResolve)
		return nil, errFactoryResolve
	}

	errValidatorResolve := container.Resolve(&validator)
	if errValidatorResolve != nil {
		utils.F("Couldnt resolve GetCursorEvents validator: %v", errValidatorResolve)
		return nil, errValidatorResolve
	}

	errResolverResolve := container.Resolve(&resolver)
	if errResolverResolve != nil {
		utils.F("Couldnt resolve GetCursorEvents resolver: %v", errResolverResolve)
		return nil, errResolverResolve
	}

	return resolvers.Resolve(ctx, factory.New(timelineID, accountID, limit, cursor), validator, resolver)
}

// MyAccountTimelines is the resolver for the myAccountTimelines field.
func (r *queryResolver) MyAccountTimelines(ctx context.Context, accountID int) ([]*model.ShortTimeline, error) {
	var resolver myAccountTimelines.Resolver
	var userExtractor domainUser.UserExtractor

	errResolverResolve := container.Resolve(&resolver)
	if errResolverResolve != nil {
		utils.F("Couldnt resolve MyAccountTimelines resolver: %v", errResolverResolve)
		return nil, errResolverResolve
	}

	errExtractorResolve := container.Resolve(&userExtractor)
	if errExtractorResolve != nil {
		utils.F("Couldnt resolve MyAccountTimelines extractor: %v", errExtractorResolve)
		return nil, errExtractorResolve
	}

	token := appContext.GetToken(ctx)
	user, err := userExtractor.ExtractUserFromToken(ctx, token)
	if err != nil {
		return nil, err
	}

	return resolver.Resolve(ctx, accountID, user.ID)
}

// Timeline is the resolver for the timeline field.
func (r *queryResolver) Timeline(ctx context.Context, timelineID int) (*model.Timeline, error) {
	var factory getTimelineResover.GetTimelineArgumentFactory
	var validator resolvers.Validator[getTimelineResover.GetTimelineArguments, getTimelineResover.ValidGetTimelineArguments]
	var resolver resolvers.Resolver[*model.Timeline, getTimelineResover.ValidGetTimelineArguments]

	errFactoryResolve := container.Resolve(&factory)
	if errFactoryResolve != nil {
		utils.F("Couldnt resolve GetTimeline factory: %v", errFactoryResolve)
		return nil, errFactoryResolve
	}

	errValidatorResolve := container.Resolve(&validator)
	if errValidatorResolve != nil {
		utils.F("Couldnt resolve GetTimeline validator: %v", errValidatorResolve)
		return nil, errValidatorResolve
	}

	errResolverResolve := container.Resolve(&resolver)
	if errResolverResolve != nil {
		utils.F("Couldnt resolve GetTimeline resolver: %v", errResolverResolve)
		return nil, errResolverResolve
	}

	return resolvers.Resolve(ctx, factory.New(timelineID), validator, resolver)
}

// Event is the resolver for the event field.
func (r *queryResolver) Event(ctx context.Context, eventID int, accountID *int) (*model.TimelineEvent, error) {
	var factory getEventResover.GetEventArgumentFactory
	var validator resolvers.Validator[getEventResover.GetEventArguments, getEventResover.ValidGetEventArguments]
	var resolver resolvers.Resolver[*model.TimelineEvent, getEventResover.ValidGetEventArguments]

	errFactoryResolve := container.Resolve(&factory)
	if errFactoryResolve != nil {
		utils.F("Couldnt resolve GetEvent factory: %v", errFactoryResolve)
		return nil, errFactoryResolve
	}

	errValidatorResolve := container.Resolve(&validator)
	if errValidatorResolve != nil {
		utils.F("Couldnt resolve GetEvent validator: %v", errValidatorResolve)
		return nil, errValidatorResolve
	}

	errResolverResolve := container.Resolve(&resolver)
	if errResolverResolve != nil {
		utils.F("Couldnt resolve GetEvent resolver: %v", errResolverResolve)
		return nil, errResolverResolve
	}

	return resolvers.Resolve(ctx, factory.New(eventID, accountID), validator, resolver)
}

// Account is the resolver for the account field.
func (r *timelineResolver) Account(ctx context.Context, obj *model.Timeline) (*model.ShortAccount, error) {
	var accountModel account.Model
	var settingsModel settings.Model

	if err := container.Resolve(&accountModel); err != nil {
		return nil, fmt.Errorf("failed to resolve account model: %w", err)
	}

	if err := container.Resolve(&settingsModel); err != nil {
		return nil, fmt.Errorf("failed to resolve settings model: %w", err)
	}

	accountEntity, err := accountModel.GetAccount(obj.AccountID)
	if err != nil {
		return nil, err
	}
	settings := settingsModel.GetSettings(enumvalues.SettingsTypeAccount, accountEntity.ID)

	return convert.ToShortAccount(*accountEntity, settings), nil
}

// Timeline is the resolver for the timeline field.
func (r *timelineEventResolver) Timeline(ctx context.Context, obj *model.TimelineEvent) (*model.Timeline, error) {
	var timelineModel timeline.Timeline

	if err := container.Resolve(&timelineModel); err != nil {
		return nil, fmt.Errorf("failed to resolve timeline model: %w", err)
	}

	timelineEntity, err := timelineModel.GetTimeline(obj.TimelineID)
	if err != nil {
		return nil, err
	}
	return &model.Timeline{
		ID:        timelineEntity.ID,
		Name:      pointer.String(timelineEntity.Name),
		AccountID: timelineEntity.QueryAccount().OnlyIDX(ctx),
	}, nil
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

// Timeline returns TimelineResolver implementation.
func (r *Resolver) Timeline() TimelineResolver { return &timelineResolver{r} }

// TimelineEvent returns TimelineEventResolver implementation.
func (r *Resolver) TimelineEvent() TimelineEventResolver { return &timelineEventResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type timelineResolver struct{ *Resolver }
type timelineEventResolver struct{ *Resolver }
