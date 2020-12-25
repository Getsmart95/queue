package models

type CredentialStatus struct {
	Ok bool `json:"ok"`
	Login bool `json:"login"`
	Password bool `json:"password"`
}

type ResponseToken struct {
	Ok bool `json:"ok"`
	Token string `json:"token"`
	ExpiredIn int `json:"expired_In"`
	Status rune `json:"status"`
	User User
}
