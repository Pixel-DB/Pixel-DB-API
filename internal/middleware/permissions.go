package middleware

var rolePermissions = map[string][]string{
	"user": {
		"pixelart.upload",
	},
	"moderator": {
		"pixelart.upload",
		"pixelart.review",
		"users.view",
	},
	"admin": {
		"pixelart.upload",
		"pixelart.review",
		"users.view",
		"users.delete",
	},
}

func HasPermission(role string, permission string) bool {
	permissions := rolePermissions[role]
	for _, p := range permissions {
		if p == permission {
			return true
		}
	}
	return false
}
