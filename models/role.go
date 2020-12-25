package models

type Role struct {
	ID int				`json:"id"`
	Name string			`json:"name"`
	DisplayName string	`json:"display_name"`
	Description string	`json:"description"`
}

type (
	RoleList struct {
		Roles []Role
	}
)