package dto

import "time"

type Page struct {
	PageNo    int         `json:"pageNo"`
	PageSize  int         `json:"pageSize"`
	HasResult bool        `json:"hasResult"`
	Result    interface{} `json:"result"`
}

type LoginDto struct {
	Id    int64  `json:"id"`
	User  string `json:"user"`
	Token string `json:"token"`
}

type ListFileDto struct {
	Id         int64
	FileName   string
	FileUrl    string
	CreateTime time.Time
}
