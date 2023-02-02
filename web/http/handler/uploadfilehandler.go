package handler

import (
	"ccloud/web/constant"
	"ccloud/web/entity"
	"ccloud/web/entity/cmd"
	"ccloud/web/http/util"
	"ccloud/web/log"
	"ccloud/web/service"
	"net/http"
)

type uploadfilehandler struct {
	pattern       string
	handler       http.Handler
	uploadservice service.UploadFileService
}

func NewUploadFileHandler() (HttpHandler, error) {
	uploadService, err := service.NewUploadFileService()
	if err != nil {
		return nil, err
	}

	upload := &uploadfilehandler{
		pattern:       "/file/upload",
		uploadservice: uploadService,
	}

	upload.handler = upload
	return upload, nil
}

func (upload uploadfilehandler) Pattern() string {
	return upload.pattern
}

func (upload uploadfilehandler) Handler() http.Handler {
	return upload.handler
}

func (upload *uploadfilehandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	file, header, err := r.FormFile("file")
	if err != nil {
		log.Logger.Error(err)
		resp := entity.Fail(constant.FileReadError)
		util.WriteJson(w, resp)
		return
	}

	createTime := r.FormValue("createTime")
	cmd := cmd.UploadCmd{
		File:       file,
		FileHeader: header,
		CreateTime: createTime,
		MediaType:  r.FormValue("mediaType"),
	}
	defer file.Close()
	userId := r.Header.Get("x-user-id")
	cmd.UserId = userId
	response := upload.uploadservice.Upload(&cmd)
	util.WriteJson(w, response)
}
