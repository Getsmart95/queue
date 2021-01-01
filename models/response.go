package models

import "time"

type CredentialStatus struct {
	Ok bool `json:"ok"`
	Login bool `json:"login"`
	Password bool `json:"password"`
}

type ResponseUser struct {
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Login     string    `json:"login"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Role 	  string 	`json:"role"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
type ResponseToken struct {
		Ok        bool   `json:"ok"`
		Token     string `json:"token"`
		User      ResponseUser
}

type ResponseStatus struct {
		Ok      bool   `json:"ok"`
		Message string `json:"message"`
		Status rune `json:"status"`
}

type JWTUserRole struct {
		RoleID int    `json:"role_id"`
		UserID int    `json:"user_id"`
		Name   string `json:"name"`
}
