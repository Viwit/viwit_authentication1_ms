package models

import "time"

type User struct {
	User_id       int       `json:"user_id,omitempty"`
	Firstname     string    `json:"firstname,omitempty"`
	Lastname      string    `json:"lastname,omitempty"`
	Email         string    `json:"email,omitempty"`
	Reg_date      time.Time `json:"reg_date,omitempty"`
	User_pasword  string    `json:"user_pasword,omitempty"`
	Wallet_id     int       `json:"wallet_id,omitempty"`
	Block_account bool      `json:"block_account,omitempty"`
}
