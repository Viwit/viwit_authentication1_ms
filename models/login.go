package models

type Login struct {
	Login_id   int  `json:"login_id,omitempty"`
	User_id    User `json:"user_id,omitempty"`
	Login_date int  `json:"login_date,omitempty"`
}
