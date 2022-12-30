package models

type UserDB struct {
	Id       string `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
	Enabled  bool   `json:"enabled"`
}
