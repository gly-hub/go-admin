package logic

import (
	"github.com/gly-hub/go-admin/authorize/global"
	"github.com/gly-hub/go-admin/authorize/internal/dao"
	"github.com/gly-hub/go-admin/authorize/internal/enum"
	"github.com/gly-hub/go-admin/authorize/internal/model"
	"github.com/gly-hub/go-admin/authorize/tools/redisx"
	"github.com/gly-hub/go-dandelion/logger"
	"gorm.io/gorm"
	"time"
)

func init() {
	// 注册token过期订阅
	redisx.RegisterExpireFunc(dao.RedisKey.TokenKey(""), Auth.LoginExpire)
}

var Auth authLogic

type authLogic struct {
}

func (al *authLogic) Login(userId, password, ip string) (string, error) {
	userInfo, uErr := dao.AdminUser.GetAdminUserInfo(userId)
	if uErr == gorm.ErrRecordNotFound {
		return "", enum.LoginUserNameNotFound
	}

	if uErr != nil {
		logger.Error(uErr)
		return "", enum.DataBaseErr
	}

	if password != userInfo.Password {
		return "", enum.LoginPasswordErr
	}

	// 生成token
	var meta = model.AuthMeta{
		Id:       userInfo.Id,
		UserId:   userInfo.UserId,
		UserName: userInfo.Name,
	}
	var (
		token string
		err   error
	)
	if token, err = global.JwtGlobal.Token(meta); err != nil {
		logger.Error(err)
		return "", enum.SystemErr
	}
	// redis中存储token
	if err = dao.AdminAuth.InsertTokenToRedis(userId, token); err != nil {
		logger.Error(err)
		return "", enum.SystemErr
	}

	// 存储用户登录日志信息 TODO

	return token, nil
}

// CheckToken 校验token
func (al *authLogic) CheckToken(token string) (meta model.AuthMeta, newToken string, err error) {
	// 校验token
	if err = global.JwtGlobal.Parse(token, &meta); err != nil {
		logger.Error(err)
		return meta, "", enum.TokenExpire
	}
	// 校验和redis是否一致
	if redisToken, err := dao.AdminAuth.GetTokenFormRedis(meta.UserId); err != nil || redisToken != token {
		logger.Error(err)
		return meta, "", enum.TokenExpire
	}

	// token续期，token除jwt有效期外，
	// 增加请求刷新，当12小时内未有请求
	// 进行刷新，将清除redis中的缓存
	if expireAt, err := global.JwtGlobal.ExpireTime(token); err != nil {
		return meta, "", enum.TokenExpire
	} else if expireAt < time.Now().Add(time.Hour).Unix() {
		// 生成新token
		// 需要进行token自动刷新，token过
		// 期时间小于一小时，将会返回新的
		// token
		newToken, err = global.JwtGlobal.Token(meta)
		if err != nil {
			logger.Error(err)
			return meta, "", enum.TokenExpire
		}
		// redis插入新的缓存
		if err = dao.AdminAuth.InsertTokenToRedis(meta.UserId, newToken); err != nil {
			logger.Error(err)
			return meta, "", enum.TokenExpire
		}
	}

	return
}

func (al *authLogic) Logout(userId string) error {
	// 清除redis中的token缓存
	if err := dao.AdminAuth.DelTokenFormRedis(userId); err != nil {
		logger.Error(err)
		return enum.SystemErr
	}
	// 数据库插入用户手动注销记录 TODO
	return nil
}

func (al *authLogic) LoginExpire(key string) {
	// 数据库插入用户token过期，自动退出记录 TODO
	logger.Debug(key)
	return
}
