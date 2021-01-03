package models

type Terminal struct {
	ID             int `json:"id,omitempty"`
	TerminalNumber int `json:"terminal_number,omitempty"`
	CityID         int `json:"city_id,omitempty"`
	BranchID       int `json:"branch_id,omitempty"`
	UserID         int `json:"user_id,omitempty"`
}
