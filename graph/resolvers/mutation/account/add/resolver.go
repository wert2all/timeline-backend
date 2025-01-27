package add

import (
	"context"

	"timeline/backend/db/model/account"
	"timeline/backend/db/model/settings"
	"timeline/backend/graph/convert"
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
	enumvalues "timeline/backend/lib/enum-values"
)

type (
	resolverImpl struct {
		accountModel  account.Model
		settingsModel settings.Model
	}
)

// Resolve implements resolvers.Resolver.
func (r resolverImpl) Resolve(ctx context.Context, arguments resolvers.ValidArguments[ValidAddAccountArguments]) (*model.ShortAccount, error) {
	args := arguments.GetArguments()
	account, err := r.accountModel.NewAccount(ctx, args.user, args.name)
	if err != nil {
		return nil, err
	}
	settings := r.settingsModel.GetSettings(enumvalues.SettingsTypeAccount, account.ID)
	return convert.ToShortAccount(*account, settings), nil
}

func NewResolver(model account.Model, settingsModel settings.Model) resolvers.Resolver[*model.ShortAccount, ValidAddAccountArguments] {
	return resolverImpl{accountModel: model, settingsModel: settingsModel}
}
