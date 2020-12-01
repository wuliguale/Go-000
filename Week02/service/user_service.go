package service

import (
	"Go-000/Week02/dao"
	"database/sql"
)

type UserService struct {
	userDao *dao.UserDao
}

func NewUserService(db *sql.DB) *UserService{
	userDao := dao.NewUserDao(db)
	return &UserService{userDao:userDao}
}


func (userService *UserService) GetByName(name string) (*dao.UserModel, error) {
	userModel, err := userService.userDao.GetByName(name)

	if err != nil {
		return nil, err
	}

	return userModel, nil
}
