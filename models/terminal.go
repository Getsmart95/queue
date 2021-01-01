package models

type Terminal struct {
	ID int
	TerminalNumber int `json:"terminal_number"`
	CityID int `json:"city_id"`
	BranchID int `json:"branch_id"`
	UserID int `json:"user_id"`
}