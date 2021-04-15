package models

import "time"

type Login struct {
	Login_id   int       `json:"login_id,omitempty"`
	User_id    User      `json:"user_id,omitempty"`
	Login_date time.Time `json:"login_date,omitempty"`
}
