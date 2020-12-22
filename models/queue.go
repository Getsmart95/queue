package models

import "time"

type Queue struct {
 	ID int
 	UserID int
 	PurposeID int
 	TimeID int
 	CityID int
 	BranchID int
 	Status string
 	Date time.Time
 	StartAt time.Time
 	FinishAt time.Time
 	CreatedAt time.Time
 }