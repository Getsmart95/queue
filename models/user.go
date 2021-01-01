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
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role 	  string	`json:"role"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}


