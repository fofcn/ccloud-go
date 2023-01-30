package cmd

import "mime/multipart"

type PageRequest struct {
	PageNo   int
	PageSize int
}

type LoginCmd struct {
	User string `json:"username"`
	Pass string `json:"password"`
}

type UploadCmd struct {
	File       multipart.File
	FileHeader *multipart.FileHeader
	CreateTime string
	MediaType  string
	UserId     string
}

type ListFileCmd struct {
	PageRequest PageRequest
	UserId      int64
	OrderBy     string
}
