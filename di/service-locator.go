package di

import (
	"timeline/backend/db/model/event"
	"timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
	eventRepository "timeline/backend/db/repository/event"
	tagRepository "timeline/backend/db/repository/tag"
	timelineRepository "timeline/backend/db/repository/timeline"
	userRepository "timeline/backend/db/repository/user"
	"timeline/backend/ent"
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"

	"golang.org/x/net/context"
)

type QueryResolversServiceLocator interface{}

type ResolverOperationServiceLocator[T any, V any, R any, F any] interface {
	ArgumentFactory() F
	Validator() resolvers.Validator[T, V]
	Resolver() resolvers.Resolver[R, V]
}

type MutationResolversServiceLocator interface {
	Authorize() ResolverOperationServiceLocator[resolvers.AuthorizeArguments, resolvers.ValidAuthorizeArguments, *model.User, resolvers.AuthorizeArgumentFactory]
	AddTimeline() ResolverOperationServiceLocator[resolvers.AddTimelineArguments, resolvers.ValidAddTimelineArguments, *model.ShortUserTimeline, resolvers.AddTimelineArgumentFactory]
	AddEvent() ResolverOperationServiceLocator[resolvers.AddEventArguments, resolvers.ValidAddEventArguments, *model.TimelineEvent, resolvers.AddEventArgumentFactory]
	DeleteEvent() ResolverOperationServiceLocator[resolvers.DeleteEventArguments, resolvers.ValidDeleteEventArguments, model.Status, resolvers.DeleteEventArgumentFactory]
}

type ResolversServiceLocator interface {
	Query() QueryResolversServiceLocator
	Mutation() MutationResolversServiceLocator
}

type ModelsServiceLocator interface {
	Users() user.UserModel
	Timeline() timeline.UserTimeline
	Events() event.Model
}

type RepositoriesServiceLocator interface {
	User() userRepository.Repository
	Timeline() timelineRepository.Repository
	Event() eventRepository.Repository
	Tag() tagRepository.Repository
}

type ServiceLocator interface {
	Resolvers() ResolversServiceLocator
	Models() ModelsServiceLocator
	Repositories() RepositoriesServiceLocator
	Context() context.Context
	DbClient() *ent.Client
}
type serviceLocator struct {
	context                    context.Context
	client                     *ent.Client
	repositoriesServiceLocator RepositoriesServiceLocator
	resolversServiceLocator    ResolversServiceLocator
	modelsServiceLocator       ModelsServiceLocator
}

type resolversServiceLocator struct {
	queryResolver                   QueryResolversServiceLocator
	mutationResolversServiceLocator MutationResolversServiceLocator
}
type (
	queryResolverServiceLocator     struct{}
	mutationResolversServiceLocator struct {
		authorizeServiceLocator   ResolverOperationServiceLocator[resolvers.AuthorizeArguments, resolvers.ValidAuthorizeArguments, *model.User, resolvers.AuthorizeArgumentFactory]
		addTimelineServiceLocator ResolverOperationServiceLocator[resolvers.AddTimelineArguments, resolvers.ValidAddTimelineArguments, *model.ShortUserTimeline, resolvers.AddTimelineArgumentFactory]
		addEventServiceLocator    ResolverOperationServiceLocator[resolvers.AddEventArguments, resolvers.ValidAddEventArguments, *model.TimelineEvent, resolvers.AddEventArgumentFactory]
		deleteEventServiceLocator ResolverOperationServiceLocator[resolvers.DeleteEventArguments, resolvers.ValidDeleteEventArguments, model.Status, resolvers.DeleteEventArgumentFactory]
	}
)

func (m mutationResolversServiceLocator) AddEvent() ResolverOperationServiceLocator[resolvers.AddEventArguments, resolvers.ValidAddEventArguments, *model.TimelineEvent, resolvers.AddEventArgumentFactory] {
	return m.addEventServiceLocator
}

func (m mutationResolversServiceLocator) Authorize() ResolverOperationServiceLocator[resolvers.AuthorizeArguments, resolvers.ValidAuthorizeArguments, *model.User, resolvers.AuthorizeArgumentFactory] {
	return m.authorizeServiceLocator
}

func (m mutationResolversServiceLocator) AddTimeline() ResolverOperationServiceLocator[resolvers.AddTimelineArguments, resolvers.ValidAddTimelineArguments, *model.ShortUserTimeline, resolvers.AddTimelineArgumentFactory] {
	return m.addTimelineServiceLocator
}

func (m mutationResolversServiceLocator) DeleteEvent() ResolverOperationServiceLocator[resolvers.DeleteEventArguments, resolvers.ValidDeleteEventArguments, model.Status, resolvers.DeleteEventArgumentFactory] {
	return m.deleteEventServiceLocator
}

type modelsServiceLocator struct {
	locator ServiceLocator
}

func (m modelsServiceLocator) Users() user.UserModel {
	return user.NewUserModel(m.locator.Repositories().User())
}

func (m modelsServiceLocator) Timeline() timeline.UserTimeline {
	return timeline.NewTimelineModel(m.locator.Repositories().Timeline())
}

func (m modelsServiceLocator) Events() event.Model {
	return event.NewEventModel(m.locator.Repositories().Event(), m.locator.Repositories().Tag())
}

type authorizeServiceLocator struct {
	locator ServiceLocator
}

type addTimelineServiceLocator struct {
	locator ServiceLocator
}

type deleteEventServiceLocator struct {
	locator ServiceLocator
}

func (d deleteEventServiceLocator) ArgumentFactory() resolvers.DeleteEventArgumentFactory {
	return resolvers.DeleteEventArgumentFactory{}
}

func (d deleteEventServiceLocator) Validator() resolvers.Validator[resolvers.DeleteEventArguments, resolvers.ValidDeleteEventArguments] {
	return resolvers.NewDeleteEventValidator(d.locator.Models().Users(), d.locator.Models().Events())
}

func (d deleteEventServiceLocator) Resolver() resolvers.Resolver[model.Status, resolvers.ValidDeleteEventArguments] {
	return resolvers.NewDeleteEventResolver(d.locator.Repositories().Event())
}

type addEventServiceLocator struct {
	locator ServiceLocator
}

func (a addEventServiceLocator) ArgumentFactory() resolvers.AddEventArgumentFactory {
	return resolvers.AddEventArgumentFactory{}
}

func (a addEventServiceLocator) Validator() resolvers.Validator[resolvers.AddEventArguments, resolvers.ValidAddEventArguments] {
	return resolvers.NewAddEventValidator(a.locator.Models().Timeline())
}

func (a addEventServiceLocator) Resolver() resolvers.Resolver[*model.TimelineEvent, resolvers.ValidAddEventArguments] {
	return resolvers.NewAddEventResolver(a.locator.Models().Events(), a.locator.Models().Timeline())
}

func (a addTimelineServiceLocator) ArgumentFactory() resolvers.AddTimelineArgumentFactory {
	return resolvers.AddTimelineArgumentFactory{}
}

func (a addTimelineServiceLocator) Validator() resolvers.Validator[resolvers.AddTimelineArguments, resolvers.ValidAddTimelineArguments] {
	return resolvers.NewAddtimelineValidator(a.locator.Models().Users())
}

func (a addTimelineServiceLocator) Resolver() resolvers.Resolver[*model.ShortUserTimeline, resolvers.ValidAddTimelineArguments] {
	return resolvers.NewAddTimelineResolver(a.locator.Models().Users(), a.locator.Models().Timeline())
}

type repositoriesServiceLocator struct {
	locator ServiceLocator
}

func (r repositoriesServiceLocator) Tag() tagRepository.Repository {
	return tagRepository.NewRepository(r.locator.Context(), r.locator.DbClient())
}

func (r repositoriesServiceLocator) Event() eventRepository.Repository {
	return eventRepository.NewRepository(r.locator.Context(), r.locator.DbClient())
}

func (r repositoriesServiceLocator) Timeline() timelineRepository.Repository {
	return timelineRepository.NewTimelineRepository(r.locator.Context(), r.locator.DbClient())
}

func (r repositoriesServiceLocator) User() userRepository.Repository {
	return userRepository.NewUserRepository(r.locator.Context(), r.locator.DbClient())
}

func (s serviceLocator) Repositories() RepositoriesServiceLocator {
	return s.repositoriesServiceLocator
}

func (s serviceLocator) Models() ModelsServiceLocator { return s.modelsServiceLocator }

func (s serviceLocator) Context() context.Context { return s.context }

func (s serviceLocator) DbClient() *ent.Client { return s.client }

func (a authorizeServiceLocator) ArgumentFactory() resolvers.AuthorizeArgumentFactory {
	return resolvers.AuthorizeArgumentFactory{}
}

func (a authorizeServiceLocator) Validator() resolvers.Validator[resolvers.AuthorizeArguments, resolvers.ValidAuthorizeArguments] {
	return resolvers.NewAuthorizeValidator(a.locator.Models().Users())
}

func (a authorizeServiceLocator) Resolver() resolvers.Resolver[*model.User, resolvers.ValidAuthorizeArguments] {
	return resolvers.NewAutorizeResolver(a.locator.Models().Timeline())
}

func (r resolversServiceLocator) Query() QueryResolversServiceLocator { return r.queryResolver }

func (r resolversServiceLocator) Mutation() MutationResolversServiceLocator {
	return r.mutationResolversServiceLocator
}

func (s serviceLocator) Resolvers() ResolversServiceLocator { return s.resolversServiceLocator }

func NewServiceLocator(context context.Context, client *ent.Client) ServiceLocator {
	locator := serviceLocator{context: context, client: client}
	locator.repositoriesServiceLocator = newRepositoriesServiceLocator(locator)
	locator.modelsServiceLocator = newModelsServiceLocator(locator)
	locator.resolversServiceLocator = newResolversServiceLocator(locator)
	return locator
}

func newRepositoriesServiceLocator(locator ServiceLocator) RepositoriesServiceLocator {
	return repositoriesServiceLocator{locator: locator}
}

func newModelsServiceLocator(locator ServiceLocator) ModelsServiceLocator {
	return modelsServiceLocator{locator: locator}
}

func newResolversServiceLocator(locator ServiceLocator) ResolversServiceLocator {
	return resolversServiceLocator{queryResolver: newQueryResolver(), mutationResolversServiceLocator: newMutationResolversServiceLocator(locator)}
}

func newQueryResolver() QueryResolversServiceLocator { return queryResolverServiceLocator{} }

func newMutationResolversServiceLocator(locator ServiceLocator) MutationResolversServiceLocator {
	return mutationResolversServiceLocator{
		authorizeServiceLocator:   newAuthorizeServiceLocator(locator),
		addTimelineServiceLocator: newAddTimelineServiceLocator(locator),
		addEventServiceLocator:    newAddEventServiceLocator(locator),
		deleteEventServiceLocator: newDeleteEventServiceLocator(locator),
	}
}
func newAddTimelineServiceLocator(locator ServiceLocator) ResolverOperationServiceLocator[resolvers.AddTimelineArguments, resolvers.ValidAddTimelineArguments, *model.ShortUserTimeline, resolvers.AddTimelineArgumentFactory] {
	return addTimelineServiceLocator{locator: locator}
}

func newAuthorizeServiceLocator(locator ServiceLocator) ResolverOperationServiceLocator[resolvers.AuthorizeArguments, resolvers.ValidAuthorizeArguments, *model.User, resolvers.AuthorizeArgumentFactory] {
	return authorizeServiceLocator{locator: locator}
}

func newAddEventServiceLocator(locator ServiceLocator) ResolverOperationServiceLocator[resolvers.AddEventArguments, resolvers.ValidAddEventArguments, *model.TimelineEvent, resolvers.AddEventArgumentFactory] {
	return addEventServiceLocator{locator: locator}
}

func newDeleteEventServiceLocator(locator ServiceLocator) ResolverOperationServiceLocator[resolvers.DeleteEventArguments, resolvers.ValidDeleteEventArguments, model.Status, resolvers.DeleteEventArgumentFactory] {
	return deleteEventServiceLocator{locator: locator}
}
