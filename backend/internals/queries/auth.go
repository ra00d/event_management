package queries

import (
	"github.com/ra00d/event_management/internals/configs"
	"github.com/ra00d/event_management/internals/constants/auth"
)

// THIS FUNCTION PANICS
func MustCreateUser(role int, username string, password string, email string) error {
	permissions := auth.GetUserPermissions(role)
	tx := configs.AppDB.MustBegin()
	res := tx.MustExec(
		"INSERT INTO users (username,password_hash,email,is_active,role_id) VALUES(?,?,?,true,?)",
		username,
		password,
		email,
		role,
	)
	userId, _ := res.LastInsertId()
	for _, v := range permissions {
		tx.MustExec("INSERT INTO user_permissions (user_id,permission_id) VALUES (?,?)", userId, v)
	}
	err := tx.Commit()
	if err != nil {
		return err
	}
	return nil
}
