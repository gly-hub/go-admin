package dao

import (
	"bytes"
	"github.com/gly-hub/go-admin/authorize/internal/enum"
)

var RedisKey redisKeyDao

type redisKeyDao struct {
}

func (r redisKeyDao) join(str ...string) string {
	var bt bytes.Buffer
	for _, s := range str {
		bt.WriteString(s)
	}
	return bt.String()
}

func (r redisKeyDao) TokenKey(userId string) string {
	return r.join(enum.RedisPrefix, "token:", userId)
}
