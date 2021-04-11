package models

type User struct {
	User_id       int    `json:"user_id,omitempty"`
	Firstname     string `json:"firstname,omitempty"`
	Lastname      string `json:"lastname,omitempty"`
	Email         string `json:"email,omitempty"`
	Reg_date      string `json:"reg_date,omitempty"`
	User_password string `json:"user_password,omitempty"`
	Credit_card   int    `json:"credit_card,omitempty"`
	Token_id      int    `json:"token_id,omitempty"`
}
