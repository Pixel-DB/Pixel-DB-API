package middleware

var rolePermissions = map[string][]string{
	"user": {
		"pixelart.upload",
	},
	"moderator": {
		"pixelart.upload",
		"pixelart.review",
	},
	"admin": {
		"pixelart.upload",
		"pixelart.review",
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
