package permisions

import (
	"github.com/les-cours/learning-api/types"
)

func CanCreate(user *types.UserToken) bool {

	if !user.Create.LEARNING {
		return false
	}

	return true
}
