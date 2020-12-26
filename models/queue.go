package models

import (
	"github.com/jackc/pgtype"
)

type Queue struct {
	ID        int       `json:"id"`
	QueueCode string    `json:"queue_code"`
	UserID    int       `json:"user_id"`
	CityID    int       `json:"city_id"`
	BranchID  int       `json:"branch_id"`
	PurposeID int       `json:"purpose_id"`
	TimeID    int       `json:"time_id"`
	Status    string    `json:"status"`
	Date      pgtype.Date	`json:"date"`
	StartAt   pgtype.Timestamp `json:"start_at"`
	FinishAt  pgtype.Timestamp `json:"finish_at"`
	CreatedAt pgtype.Timestamp `json:"created_at"`
}