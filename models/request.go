package models

type RequestDate struct {
	Date string `json:"date,omitempty"`
}

type RequestStatus struct {
	Status string `json:"status,omitempty"`
}

type RequestUser struct {
	UserID int `json:"user_id,omitempty"`
}

type RequestLogin struct {
	Login string `json:"login,omitempty"`
}

type RequestTerminal struct {
	PurposeID int `json:"purpose_id,omitempty"`
	Date string   `json:"date,omitempty"`
}

type RequestQueue struct {
	CityID     int      `json:"city_id,omitempty"`
	BranchID   int      `json:"branch_id,omitempty"`
	PurposeID  int      `json:"purpose_id,omitempty"`
	Date       string	`json:"date,omitempty"`
}

