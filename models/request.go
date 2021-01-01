package models

type RequestDate struct {
	Date string `json:"date"`
}

type RequestStatus struct {
	Status string `json:"status"`
}

type RequestUser struct {
	UserID int `json:"user_id"`
}

type RequestLogin struct {
	Login string `json:"login"`
}

type RequestTerminal struct {
	PurposeID int `json:"purpose_id"`
	Date string `json:"date"`
}


