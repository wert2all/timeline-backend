package save

import (
	"timeline/backend/graph/model"
	"timeline/backend/graph/resolvers"
)

type (
	SaveAccountArgumentsFactory struct{}
	SaveAccountArguments        struct {
		accountID int
		name      string
		avatarID  *int
	}
)

func (s SaveAccountArgumentsFactory) New(accountID int, account model.SaveAccountInput) resolvers.Arguments[SaveAccountArguments] {
	return SaveAccountArguments{accountID: accountID, name: account.Name, avatarID: account.AvatarID}
}

func (g SaveAccountArguments) GetArguments() SaveAccountArguments { return g }
