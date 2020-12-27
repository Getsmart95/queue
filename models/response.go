package models

import "time"

type CredentialStatus struct {
	Ok bool `json:"ok"`
	Login bool `json:"login"`
	Password bool `json:"password"`
}

type ResponseUser struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Surname   string    `json:"surname"`
	Login     string    `json:"login"`
	Email     string    `json:"email"`
	Phone     string    `json:"phone"`
	Status    bool      `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}
type ResponseToken struct {
	Ok bool `json:"ok"`
	Token string `json:"token"`
	ExpiredIn int `json:"expired_In"`
	Status rune `json:"status"`
	User ResponseUser
}

type ResponseStatus struct {
	Ok bool `json:"ok"`
	Status rune `json:"status"`
	Message string `json:"message"`
}

type JWTUserRole struct {
	RoleID int `json:"role_id"`
	UserID int `json:"user_id"`
	Name string `json:"name"`
}