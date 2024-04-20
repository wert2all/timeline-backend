package context

import (
	"context"
	"timeline/backend/db/model/user"
)

type AppContextUserKey struct{}

type (
	AppModelsKey struct{}
	AppModels    struct {
		Users user.UserModel
	}
)

type AppContext interface {
	GetContext() context.Context
	GetModels() AppModels
	SetUserID(id int) context.Context
	GetUserID() *int
}

type AppContextImpl struct {
	context context.Context
}

// AddUserId implements AppContext.
func (a AppContextImpl) SetUserID(id int) context.Context {
	a.context = context.WithValue(a.context, AppContextUserKey{}, id)
	return a.GetContext()
}

func (a AppContextImpl) GetUserID() *int {
	value := a.context.Value(AppContextUserKey{})
	if value == nil {
		return nil
	}
	return value.(*int)
}
func (a AppContextImpl) GetContext() context.Context { return a.context }
func (a AppContextImpl) GetModels() AppModels        { return a.context.Value(AppModelsKey{}).(AppModels) }

func NewAppContext(ctx context.Context, models AppModels) AppContext {
	return AppContextImpl{context: context.WithValue(ctx, AppModelsKey{}, models)}
}

func NewModels(userModel user.UserModel) AppModels { return AppModels{Users: userModel} }
