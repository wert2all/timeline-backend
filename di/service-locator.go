package di

import (
	"golang.org/x/net/context"
	"timeline/backend/db/model/timeline"
	"timeline/backend/db/model/user"
	timelineRepository "timeline/backend/db/repository/timeline"
	userRepository "timeline/backend/db/repository/user"
	"timeline/backend/ent"
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
)

type QueryResolversServiceLocator interface{}

type ResolverOperationServiceLocator[T any, V any, R any, F any] interface {
	ArgumentFactory() F
	Validator() resolvers.Validator[T, V]
	Resolver() resolvers.Resolver[R, V]
}

type MutationResolversServiceLocator interface {
	Authorize() ResolverOperationServiceLocator[resolvers.AuthorizeArguments, resolvers.ValidAuthorizeArguments, model.User, resolvers.AuthorizeArgumentFactory]
}

type ResolversServiceLocator interface {
	Query() QueryResolversServiceLocator
	Mutation() MutationResolversServiceLocator
}

type ModelsServiceLocator interface {
	Users() user.UserModel
	Timeline() timeline.UserTimeline
}

type RepositoriesServiceLocator interface {
	User() userRepository.Repository
	Timeline() timelineRepository.Repository
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
type queryResolverServiceLocator struct{}
type mutationResolversServiceLocator struct {
	authorizeServiceLocator ResolverOperationServiceLocator[resolvers.AuthorizeArguments, resolvers.ValidAuthorizeArguments, model.User, resolvers.AuthorizeArgumentFactory]
}

type modelsServiceLocator struct {
	locator ServiceLocator
}

type authorizeServiceLocator struct {
	locator ServiceLocator
}

type repositoriesServiceLocator struct {
	locator ServiceLocator
}

func (r repositoriesServiceLocator) Timeline() timelineRepository.Repository {
	return timelineRepository.NewTimelineRepository(r.locator.Context(), r.locator.DbClient())
}

func (r repositoriesServiceLocator) User() userRepository.Repository {
	return userRepository.NewUserRepository(r.locator.Context(), r.locator.DbClient())
}

func (m modelsServiceLocator) Users() user.UserModel {
	return user.NewUserModel(m.locator.Repositories().User())
}

func (m modelsServiceLocator) Timeline() timeline.UserTimeline {
	return timeline.NewTimelineModel(m.locator.Repositories().Timeline())
}
func (s serviceLocator) Repositories() RepositoriesServiceLocator { return s.repositoriesServiceLocator }

func (s serviceLocator) Models() ModelsServiceLocator { return s.modelsServiceLocator }

func (s serviceLocator) Context() context.Context { return s.context }

func (s serviceLocator) DbClient() *ent.Client { return s.client }

func (a authorizeServiceLocator) ArgumentFactory() resolvers.AuthorizeArgumentFactory {
	return resolvers.AuthorizeArgumentFactory{}
}

func (a authorizeServiceLocator) Validator() resolvers.Validator[resolvers.AuthorizeArguments, resolvers.ValidAuthorizeArguments] {
	return resolvers.NewAuthorizeValidator(a.locator.Models().Users())
}

func (a authorizeServiceLocator) Resolver() resolvers.Resolver[model.User, resolvers.ValidAuthorizeArguments] {
	return resolvers.NewAutorizeResolver(a.locator.Models().Timeline())
}

func (m mutationResolversServiceLocator) Authorize() ResolverOperationServiceLocator[resolvers.AuthorizeArguments, resolvers.ValidAuthorizeArguments, model.User, resolvers.AuthorizeArgumentFactory] {
	return m.authorizeServiceLocator
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

func newModelsServiceLocator(locator ServiceLocator) ModelsServiceLocator {
	return modelsServiceLocator{locator: locator}
}

func newResolversServiceLocator(locator ServiceLocator) ResolversServiceLocator {
	return resolversServiceLocator{queryResolver: newQueryResolver(), mutationResolversServiceLocator: newMutationResolversServiceLocator(locator)}
}

func newQueryResolver() QueryResolversServiceLocator { return queryResolverServiceLocator{} }

func newMutationResolversServiceLocator(locator ServiceLocator) MutationResolversServiceLocator {
	return mutationResolversServiceLocator{newAuthorizeServiceLocator(locator)}
}

func newAuthorizeServiceLocator(locator ServiceLocator) ResolverOperationServiceLocator[resolvers.AuthorizeArguments, resolvers.ValidAuthorizeArguments, model.User, resolvers.AuthorizeArgumentFactory] {
	return authorizeServiceLocator{locator: locator}
}

func newRepositoriesServiceLocator(locator ServiceLocator) RepositoriesServiceLocator {
	return repositoriesServiceLocator{locator: locator}
}
