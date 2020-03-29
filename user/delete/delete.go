package delete

import (
	"compose/user/daos"
)

func deleteUser(model *RequestModel) error {
	return daos.GetUserDao().DeleteUser(model.email)
}
