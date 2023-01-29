package model

import "time"

type UserModel struct {
	Id         int64
	Username   string
	Password   string
	CreateTime time.Time
}

type MediaModel struct {
	Id             int64
	FileName       string
	StorePath      string
	FileCreateTime time.Time
	MediaType      int
	CreateTime     time.Time
}
