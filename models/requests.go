package models

type RequestQueue struct {
	QueueHash string
	UserID int
	CityID int
	BranchID int
	TimeID int
	PurposeID int
	Status string
	Date Time
	StartAt Time
	FinishAt Time
	CreatedAt Time
}

type Date struct {
	Date string `json:"date"`
}