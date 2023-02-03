package dao

import (
	"github.com/gly-hub/go-admin/authorize/global"
	"github.com/gly-hub/go-dandelion/application"
	"github.com/gomodule/redigo/redis"
)

var AdminAuth adminAuthDao

type adminAuthDao struct {
	application.DB
	application.Redis
}

// InsertTokenToRedis 插入token
func (ad *adminAuthDao) InsertTokenToRedis(userId, token string) (err error) {
	key := RedisKey.TokenKey(userId)
	_, err = ad.GetRedis().Execute(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("set", key, token, "EX", int(global.BaseConfig.Jwt.Expire)*60)
	})
	return
}

// GetTokenFormRedis 获取token
func (ad *adminAuthDao) GetTokenFormRedis(userId string) (token string, err error) {
	key := RedisKey.TokenKey(userId)
	token, err = ad.GetRedis().String(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("get", key)
	})
	return
}

// DelTokenFormRedis 获取token
func (ad *adminAuthDao) DelTokenFormRedis(userId string) (err error) {
	key := RedisKey.TokenKey(userId)
	_, err = ad.GetRedis().Execute(func(c redis.Conn) (res interface{}, err error) {
		return c.Do("del", key)
	})
	return
}
