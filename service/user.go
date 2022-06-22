package service

import (
	"errors"

	"github.com/Dlimingliang/go-admin/global"
	"github.com/Dlimingliang/go-admin/model"
	"github.com/Dlimingliang/go-admin/model/request"
	"github.com/Dlimingliang/go-admin/utils"
)

type UserService struct{}

func (userService *UserService) GetUserList(page request.PageInfo) ([]model.User, int64, error) {
	db := global.GaDb.Model(model.User{})

	var total int64
	err := db.Count(&total).Error
	if err != nil {
		return nil, total, err
	}

	var userList []model.User
	err = db.Scopes(Paginate(page.Page, page.PageSize)).Find(&userList).Error
	return userList, total, err
}

func (userService UserService) RegisterAdmin(user model.User) (model.User, error) {
	//验证用户和电话是否存在
	if result := global.GaDb.Where(&model.User{Username: user.Username}).First(&user); result.RowsAffected > 0 {
		return user, errors.New("该用户名已被使用")
	}
	if result := global.GaDb.Where(&model.User{Mobile: user.Mobile}).First(&user); result.RowsAffected > 0 {
		return user, errors.New("该电话已被使用")
	}

	//生成用户密码
	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return user, err
	}
	user.Password = hashPassword
	//创建用户并返回用户信息
	err = global.GaDb.Create(&user).Error
	return user, err
}
