package permisions

import (
	"context"
	"github.com/les-cours/learning-api/types"
)

func Admin(ctx context.Context) (*types.UserToken, error) {

	user, ok := ctx.Value("user").(*types.UserToken)
	if !ok || *user == (types.UserToken{}) {
		return nil, ErrAuth
	}

	if user.UserType != "admin" {
		return nil, ErrPermissionDenied
	}

	return user, nil
}
