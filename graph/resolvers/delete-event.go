package resolvers

import (
	"context"
	"errors"
	appContext "timeline/backend/app/context"
	"timeline/backend/db/model/event"
	"timeline/backend/db/model/user"
	eventRepository "timeline/backend/db/repository/event"
	"timeline/backend/ent"
	"timeline/backend/graph/model"
)

type (
	DeleteEventArguments      struct{ eventID int }
	ValidDeleteEventArguments struct {
		event *ent.Event
	}
	DeleteEventArgumentFactory struct{}
	deleteEventResolverImpl    struct {
		eventRepository eventRepository.Repository
	}
	validatorDeleteEventImpl struct {
		usersModel  user.UserModel
		eventsModel event.Model
	}
)

func (d deleteEventResolverImpl) Resolve(ctx context.Context, arguments ValidArguments[ValidDeleteEventArguments]) (model.Status, error) {
	deleteError := d.eventRepository.Delete(ctx, arguments.GetArguments().event)
	if deleteError != nil {
		return model.StatusError, deleteError
	}
	return model.StatusSuccess, nil
}

func (v ValidDeleteEventArguments) GetArguments() ValidDeleteEventArguments { return v }

func (v validatorDeleteEventImpl) Validate(ctx context.Context, input Arguments[DeleteEventArguments]) (ValidArguments[ValidDeleteEventArguments], error) {
	event, error := v.eventsModel.GetEventByID(input.GetArguments().eventID)
	if error != nil {
		return nil, error
	}
	timeline, error := event.QueryTimeline().Only(ctx)
	if error != nil {
		return nil, error
	}
	userEntity, error := timeline.QueryUser().Only(ctx)
	if error != nil {
		return nil, error
	}
	userAutorized, error := v.usersModel.GetUser(appContext.GetUserID(ctx))
	if error != nil {
		return nil, error
	}
	if userEntity.ID != userAutorized.ID {
		return nil, errors.New("could not delete event")
	}

	return ValidDeleteEventArguments{event: event}, nil
}

func (d DeleteEventArguments) GetArguments() DeleteEventArguments { return d }

func (f DeleteEventArgumentFactory) New(eventID int) Arguments[DeleteEventArguments] {
	return DeleteEventArguments{eventID: eventID}
}

func NewDeleteEventValidator(usersModel user.UserModel, eventsModel event.Model) Validator[DeleteEventArguments, ValidDeleteEventArguments] {
	return validatorDeleteEventImpl{usersModel: usersModel, eventsModel: eventsModel}
}

func NewDeleteEventResolver(eventRepository eventRepository.Repository) Resolver[model.Status, ValidDeleteEventArguments] {
	return deleteEventResolverImpl{eventRepository: eventRepository}
}
