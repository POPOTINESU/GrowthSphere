package domain

type Permission struct {
	value string
}

func NewPermission(permission string) (Permission, error) {
	switch permission {
	case "admin":
		return Permission{value: permission}, nil
	default:
		return Permission{value: "user"}, nil
	}
}
