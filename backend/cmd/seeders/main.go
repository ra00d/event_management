package main

import (
	"log"

	"github.com/ra00d/event_management/internals/configs"
	"github.com/ra00d/event_management/internals/constants/auth"
)

func main() {
	// HINT use transactions
	// create roles and add them to db
	configs.DataBasaInit()
	tx := configs.AppDB.MustBegin()
	for _, v := range auth.GetPermissions() {
		permission_id := v["permission_id"]
		permission_name := v["permission_name"]
		tx.MustExec(
			"insert into permissions (permission_id,permission_name) values(?,?)",
			permission_id,
			permission_name,
		)

	}
	for _, v := range auth.GetRoles() {
		role_id := v["role_id"]
		role_name := v["role_name"]
		tx.MustExec("insert into roles (role_id,role_name) values(?,?)",
			role_id,
			role_name,
		)

	}
	err := tx.Commit()
	if err != nil {
		log.Fatal(err.Error())
	}
	// create permisssions for all the system
	// link every permission to its role
}
