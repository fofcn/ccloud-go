package dao

import (
	"ccloud/web/entity/cmd"
	"ccloud/web/entity/model"
	"ccloud/web/store"
	"errors"

	"github.com/goinggo/mapstructure"
)

type MediaFileDao interface {
	SaveFile(model.MediaModel) (int64, error)
	GetFileByPage(userId int64, orderBy string, pageRequest cmd.PageRequest) ([]model.MediaModel, error)
}

type MediaFileDaoImpl struct {
	store store.SqlStore
}

func NewMediaFileDao() (MediaFileDao, error) {
	store, err := store.NewSingleSqliteStore(".", "db.sql")
	if err != nil {
		return nil, err
	}
	return &MediaFileDaoImpl{
		store: store,
	}, nil
}

func (impl MediaFileDaoImpl) SaveFile(media model.MediaModel) (int64, error) {
	sql := "insert into `media_file` (user_id, file_name, store_path, file_create_time, media_type, create_time) values (?, ?,?,?,?,?)"
	rowId, err := impl.store.Insert(sql, media.UserId, media.FileName, media.StorePath, media.FileCreateTime, media.MediaType, media.CreateTime)
	if err != nil {
		return 0, err
	}

	if rowId < 1 {
		return 0, errors.New("insert data failed")
	}

	return rowId, nil
}

func (impl MediaFileDaoImpl) GetFileByPage(userId int64, orderBy string, pageRequest cmd.PageRequest) ([]model.MediaModel, error) {
	// 计算Page
	offset := (pageRequest.PageNo - 1) * pageRequest.PageSize
	sql := "select * from `media_file` where user_id=? order by ? limit ?,? "
	fileList, err := impl.store.Query(sql, userId, orderBy, offset, pageRequest.PageSize)
	if err != nil {
		return nil, err
	}

	var result []model.MediaModel = make([]model.MediaModel, 0)
	for _, file := range fileList {
		var mediaFile model.MediaModel
		mapstructure.Decode(file, &mediaFile)
		result = append(result, mediaFile)
	}

	return result, nil
}
