package models

type Role struct {
	ID int
	Name string
	DisplayName string
	Description string
}

type (
	RoleList struct {
		Roles []Role
	}
)