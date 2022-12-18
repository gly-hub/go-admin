package logic

import (
	"github.com/gly-hub/go-admin/authorize/internal/dao"
	"github.com/gly-hub/go-admin/authorize/internal/enum"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/jinzhu/gorm"
)

var Auth authLogic

type authLogic struct {

}

func (al *authLogic) Login(userName, password string)(string, error){
	userInfo, uErr := dao.AdminUser.GetAdminUserInfo(userName)
	if uErr == gorm.ErrRecordNotFound {
		return "", enum.LoginUserNameNotFound
	}

	if uErr != nil{
		logger.Error(uErr)
		return "", enum.DataBaseErr
	}

	if password != userInfo.Password {
		return "", enum.LoginPasswordErr
	}

	return "1", nil
}

func (al *authLogic) Logout(userName string) error {
	return nil
}
