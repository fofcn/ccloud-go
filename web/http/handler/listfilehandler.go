package handler

import (
	"ccloud/web/entity/cmd"
	"ccloud/web/http/util"
	"ccloud/web/service"
	"net/http"
	"strconv"
)

type listfilehandler struct {
	pattern         string
	handler         http.Handler
	listfileservice service.ListFileService
}

func NewListFileHandler() (HttpHandler, error) {
	listfileservice, err := service.NewListFileService()
	if err != nil {
		return nil, err
	}

	handler := &listfilehandler{
		pattern:         "/file/list",
		listfileservice: listfileservice,
	}

	handler.handler = handler
	return handler, nil
}

func (listfile listfilehandler) Pattern() string {
	return listfile.pattern
}

func (listfile listfilehandler) Handler() http.Handler {
	return listfile.handler
}

func (listfile *listfilehandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()
	pageNo := params["pageNo"][0]
	pageSize := params["pageSize"][0]
	orderBy := params["orderBy"][0]
	userId := r.Header.Get("x-user-id")

	iPageNo, _ := strconv.Atoi(pageNo)
	iPageSize, _ := strconv.Atoi(pageSize)
	iUserId, _ := strconv.ParseInt(userId, 10, 64)

	cmd := cmd.ListFileCmd{
		PageRequest: cmd.PageRequest{
			PageNo:   iPageNo,
			PageSize: iPageSize,
		},
		UserId:  iUserId,
		OrderBy: orderBy,
	}

	response := listfile.listfileservice.List(&cmd)
	util.WriteJson(w, response)
}
