package dto

type LoginDto struct {
	Id    int64  `json:"id"`
	User  string `json:"user"`
	Token string `json:"token"`
}
