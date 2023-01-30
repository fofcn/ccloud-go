package service

import (
	"ccloud/web/constant"
	"ccloud/web/dao"
	"ccloud/web/entity"
	"ccloud/web/entity/cmd"
	"ccloud/web/entity/dto"
	"ccloud/web/log"
	"sync"
)

type ListFileService interface {
	List(cmd *cmd.ListFileCmd) entity.Response
}

type listfileservice struct {
	mediafiledao dao.MediaFileDao
}

var once sync.Once
var service ListFileService

func NewListFileService() (ListFileService, error) {
	mediafiledao, err := dao.NewMediaFileDao()
	if err != nil {
		return nil, err
	}
	once.Do(func() {
		service = &listfileservice{
			mediafiledao: mediafiledao,
		}
	})

	return service, nil
}

func (service listfileservice) List(cmd *cmd.ListFileCmd) entity.Response {
	fileList, err := service.mediafiledao.GetFileByPage(cmd.UserId, cmd.MediaType, "id desc", cmd.PageRequest)
	if err != nil {
		log.Logger.Errorf("query file list error using user id: %v", cmd.UserId, err)
		return entity.Fail(constant.ListUserFileError)
	}

	var result []dto.ListFileDto = make([]dto.ListFileDto, 0)
	for _, file := range fileList {
		listFileDto := dto.ListFileDto{
			Id:         file.Id,
			FileName:   file.FileName,
			FileUrl:    file.StorePath,
			CreateTime: file.CreateTime,
		}

		result = append(result, listFileDto)
	}

	// 组装Page对象
	mediaFilePage := dto.Page{
		PageNo:    cmd.PageRequest.PageNo,
		PageSize:  len(result),
		HasResult: len(result) != 0,
		Result:    result,
	}
	return entity.OKWithData(mediaFilePage)
}
