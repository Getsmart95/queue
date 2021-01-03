package models

type Role struct {
	ID int				`json:"id,omitempty"`
	Name string			`json:"name,omitempty"`
	DisplayName string	`json:"display_name,omitempty"`
	Description string	`json:"description,omitempty"`
}