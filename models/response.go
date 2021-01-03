package models

import "time"

type CredentialStatus struct {
	Ok       bool `json:"ok,omitempty"`
	Login    bool `json:"login,omitempty"`
	Password bool `json:"password,omitempty"`
}

type ResponseUser struct {
	Name      string    `json:"name,omitempty"`
	Surname   string    `json:"surname,omitempty"`
	Login     string    `json:"login,omitempty"`
	Email     string    `json:"email,omitempty"`
	Phone     string    `json:"phone,omitempty"`
	Role 	  string 	`json:"role,omitempty"`
	Status    bool      `json:"status,omitempty"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}
type ResponseToken struct {
		Ok        bool   `json:"ok,omitempty"`
		Token     string `json:"token,omitempty"`
		User      ResponseUser
}

type ResponseStatus struct {
		Ok      bool   `json:"ok,omitempty"`
		Message string `json:"message,omitempty"`
}

type JWTUserRole struct {
		RoleID int    `json:"role_id,omitempty"`
		UserID int    `json:"user_id,omitempty"`
		Name   string `json:"name,omitempty"`
}
