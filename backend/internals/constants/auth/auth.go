package auth

const (
	ADMIN           = 1005
	USER            = 2009
	ORGANIZAER      = 2222
	ADMIN_NAME      = "admin"
	USER_NAME       = "user"
	ORGANIZAER_NAME = "organizer"

	// permissions numbers
	VIEWEVENTS  = 1
	VIEWUSERS   = 2
	CREATEEVNET = 3
	DELETEEVENT = 4
	UPDATEEVENT = 5
	CREATEUSER  = 6
	DELETEUSER  = 7
	UPDATEUSER  = 8
	ENABLEUSER  = 9
	DISABLEUSER = 10

	VIEWEVENTS_NAME  = "view-events"
	VIEWUSERS_NAME   = "view-users"
	CREATEEVNET_NAME = "create-event"
	DELETEEVENT_NAME = "delete-event"
	UPDATEEVENT_NAME = "update-event"
	CREATEUSER_NAME  = "create-user"
	DELETEUSER_NAME  = "delete-user"
	UPDATEUSER_NAME  = "update-user"
	ENABLEUSER_NAME  = "enable-user"
	DISABLEUSER_NAME = "disable-user"
)

func GetRoles() []map[string]interface{} {
	return []map[string]interface{}{
		{"role_id": ADMIN, "role_name": ADMIN_NAME},
		{"role_id": USER, "role_name": USER_NAME},
		{"role_id": ORGANIZAER, "role_name": ORGANIZAER_NAME},
	}
}

func GetPermissions() []map[string]interface{} {
	return []map[string]interface{}{
		{"permission_id": CREATEUSER, "permission_name": CREATEUSER_NAME},
		{"permission_id": CREATEEVNET, "permission_name": CREATEEVNET_NAME},
		{"permission_id": DELETEUSER, "permission_name": DELETEUSER_NAME},
		{"permission_id": DELETEEVENT, "permission_name": DELETEEVENT_NAME},
		{"permission_id": UPDATEUSER, "permission_name": UPDATEUSER_NAME},
		{"permission_id": UPDATEEVENT, "permission_name": UPDATEEVENT_NAME},
		{"permission_id": ENABLEUSER, "permission_name": ENABLEUSER_NAME},
		{"permission_id": DISABLEUSER, "permission_name": DISABLEUSER_NAME},
		{"permission_id": VIEWUSERS, "permission_name": VIEWUSERS_NAME},
		{"permission_id": VIEWEVENTS, "permission_name": VIEWEVENTS_NAME},
	}
}

func GetUserPermissions(role int) []int {
	switch role {
	case ADMIN:
		return getAdminPermissions()
	case ORGANIZAER:
		return getOrganizerPermissions()
	default:
		return getUserPermissions()
	}
}

func getAdminPermissions() []int {
	// declare all permissions for the admin
	return []int{
		VIEWUSERS,
		VIEWEVENTS,
		CREATEEVNET,
		DELETEEVENT,
		UPDATEEVENT,
		CREATEUSER,
		DELETEUSER,
		UPDATEUSER,
		DISABLEUSER,
		ENABLEUSER,
	}
}

func getUserPermissions() []int {
	// declare all permissions for the admin
	return []int{
		VIEWEVENTS,
	}
}

func getOrganizerPermissions() []int {
	return []int{
		VIEWEVENTS,
		CREATEEVNET,
		DELETEEVENT,
		UPDATEEVENT,
	}
}

// add another permission group
// # e.g func GetUserPermissons() []string {
// return []string{"any permission"}
// }
