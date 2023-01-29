package dao

import (
	"ccloud/web/entity/model"
	"ccloud/web/store"
	"errors"
)

type MediaFileDao interface {
	SaveFile(model.MediaModel) (int64, error)
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
