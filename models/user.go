package models

import "time"

type User struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Surname   string    `json:"surname,omitempty"`
	Login     string    `json:"login,omitempty"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Status    bool      `json:"status,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type UserWithRole struct {
	ID        int       `json:"id,omitempty"`
	Name      string    `json:"name,omitempty"`
	Surname   string    `json:"surname,omitempty"`
	Login     string    `json:"login,omitempty"`
	Password  string    `json:"password,omitempty"`
	Email     string    `json:"email,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Role 	  string	`json:"role,omitempty"`
	Status    bool      `json:"status,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}


