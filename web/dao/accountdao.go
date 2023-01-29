package dao

import (
	"ccloud/web/entity/model"
	"ccloud/web/store"
	"errors"
	"time"
)

type AccountDao interface {
	SelectByUsername(string) (model.UserModel, error)
	InsertUser(userModel model.UserModel) (int64, error)
}

type accountdaoimpl struct {
	store store.SqlStore
}

func NewAccountDao() (AccountDao, error) {
	store, err := store.NewSingleSqliteStore(".", "db.sql")
	if err != nil {
		return nil, err
	}
	return &accountdaoimpl{
		store: store,
	}, nil
}

func (impl accountdaoimpl) SelectByUsername(username string) (model.UserModel, error) {
	// 根据用户名查询数据库
	sql := "select * from userinfo where username=?"
	data, err := impl.store.Query(sql, username)
	if err != nil {
		return model.UserModel{}, err
	}

	if len(data) == 0 {
		return model.UserModel{}, errors.New("user not found")
	}

	record := data[0]
	// 组装数据返回
	return model.UserModel{
		Id:         record["id"].(int64),
		Username:   record["username"].(string),
		Password:   record["password"].(string),
		CreateTime: record["create_time"].(time.Time),
	}, nil
}

func (impl accountdaoimpl) InsertUser(userModel model.UserModel) (int64, error) {
	sql := "insert into userinfo (username, password,create_time)values(?,?,?)"
	return impl.store.Insert(sql, userModel.Username, userModel.Password, userModel.CreateTime)
}
