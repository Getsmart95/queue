package models

import (
	"database/sql"
	"github.com/jackc/pgtype"
	"time"
)

type Queue struct {
	ID         int              `json:"id"`
	QueueCode  int              `json:"queue_code"`
	TerminalID sql.NullInt64              `json:"terminal_id"`
	UserID     sql.NullInt64              `json:"user_id"`
	CityID     sql.NullInt64              `json:"city_id"`
	BranchID   sql.NullInt64              `json:"branch_id"`
	PurposeID  sql.NullInt64              `json:"purpose_id"`
	TimeID     sql.NullInt64              `json:"time_id,omitempty"`
	Status     string           `json:"status"`
	Date       time.Time		    `json:"date"`
	StartAt    pgtype.Timestamp `json:"start_at"`
	FinishAt   pgtype.Timestamp `json:"finish_at"`
	CreatedAt  pgtype.Timestamp `json:"created_at"`
}