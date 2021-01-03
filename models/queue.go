package models

import (
	"database/sql"
	"github.com/jackc/pgtype"
	"time"
)

type Queue struct {
	ID         int              `json:"id,omitempty"`
	QueueCode  int              `json:"queue_code,omitempty"`
	TerminalID sql.NullInt64    `json:"terminal_id,omitempty"`
	UserID     sql.NullInt64    `json:"user_id,omitempty"`
	CityID     sql.NullInt64    `json:"city_id,omitempty"`
	BranchID   sql.NullInt64    `json:"branch_id,omitempty"`
	PurposeID  sql.NullInt64    `json:"purpose_id,omitempty"`
	TimeID     sql.NullInt64    `json:"time_id,omitempty"`
	Status     string           `json:"status,omitempty"`
	Date       time.Time        `json:"date,omitempty"`
	StartAt    pgtype.Timestamp `json:"start_at,omitempty"`
	FinishAt   pgtype.Timestamp `json:"finish_at,omitempty"`
	CreatedAt  pgtype.Timestamp `json:"created_at,omitempty"`
}
