package save

import (
	"context"

	"timeline/backend/db/model/settings"
	"timeline/backend/db/model/user"
	"timeline/backend/graph/convert"
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
	enumvalues "timeline/backend/lib/enum-values"
)

type (
	resolverImpl struct {
		userModel     user.UserModel
		settingsModel settings.Model
	}
)

// Resolve implements resolvers.Resolver.
func (r resolverImpl) Resolve(ctx context.Context, arguments resolvers.ValidArguments[ValidSaveAccountArguments]) (*model.ShortAccount, error) {
	args := arguments.GetArguments()
	accountEntity, err := r.userModel.SaveUserAccount(args.account, args.name, args.about, args.avatarID)
	if err != nil {
		return nil, err
	}

	settings := r.settingsModel.GetSettings(enumvalues.SettingsTypeAccount, accountEntity.ID)
	return convert.ToShortAccount(*accountEntity, settings), nil
}

func NewResolver(userModel user.UserModel, settings settings.Model) resolvers.Resolver[*model.ShortAccount, ValidSaveAccountArguments] {
	return resolverImpl{userModel: userModel, settingsModel: settings}
}
