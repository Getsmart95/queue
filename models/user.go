package models

import "time"

type User struct {
		ID        int
		Name      string
		Surname   string
		Login     string
		Password  string
		Email     string
		Phone     string
		Status    bool
		CreatedAt time.Time
	}
