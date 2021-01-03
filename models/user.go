package models

import "time"

type User struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
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


