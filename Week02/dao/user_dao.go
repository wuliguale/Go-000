package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

type UserModel struct {
	Id uint `sql:"id" json:"id"`
	Name string `sql:"name" json:"name"`
	Age uint `sql:"age" json:"age"`
	State int `sql:"state" json:"state"`
	CreateAt uint `sql:"create_at" json:"create_at"`
	UpdateAt uint `sql:"update_at" json:"update_at"`
}


type UserDao struct {
	db *sql.DB
}

func NewUserDao(db *sql.DB) *UserDao{
	return &UserDao{db:db}
}

func (u *UserDao) GetByName(name string) (*UserModel, error) {
	userModel := &UserModel{}

	err := u.db.QueryRow("select * from cyjc_user where name=?", name).Scan(&userModel.Id, &userModel.Name, &userModel.Age, &userModel.State, &userModel.CreateAt, &userModel.UpdateAt)

	if err == nil {
		return userModel, nil
	}

	//没有找到，算是特殊错误
	if err == sql.ErrNoRows {
		return nil, errors.Wrap(err, "用户名不存在")
	}

	return nil, errors.Wrap(err, "根据用户名查找用户出错")
}
