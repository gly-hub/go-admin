package redisx

import (
	"github.com/gly-hub/go-dandelion/database/redigo"
	"github.com/gly-hub/go-dandelion/logger"
	"github.com/gomodule/redigo/redis"
	"unsafe"
)

type EventType string

const (
	ExpireEvent EventType = "__keyevent@0__:expired"
)

type PSubscribeCallback func(pattern, channel, device string)

type ExpireKeyCallBack func(key string)

var (
	cbMap       map[string]PSubscribeCallback
	expireFuncs = make(map[string]ExpireKeyCallBack)
)

func PSubscribeListen(conn *redigo.Client) (err error) {
	cbMap = make(map[string]PSubscribeCallback)
	err = conn.Subscribe(func(c redis.PubSubConn) (err error) {
		// 订阅过期事件
		err = c.PSubscribe(ExpireEvent)
		if err != nil {
			logger.Error("redis Subscribe error.")
		}
		cbMap[string(ExpireEvent)] = timeoutEventCallBack
		go func() {
			for {
				switch res := c.Receive().(type) {
				case redis.Message:
					pattern := (*string)(unsafe.Pointer(&res.Pattern))
					channel := (*string)(unsafe.Pointer(&res.Channel))
					message := (*string)(unsafe.Pointer(&res.Data))
					if cb, ok := cbMap[*channel]; ok {
						cb(*pattern, *channel, *message)
					}
				case redis.Subscription:
					logger.Info("%s: %s %d", res.Channel, res.Kind, res.Count)
				case error:
					//logger.Error("error handle...")
					continue
				}
			}
			defer c.Close()
		}()
		return
	})
	return
}

func timeoutEventCallBack(pattern, channel, device string) {
	// 获取注册
	for key, function := range expireFuncs {
		if len(device) > len(key) && key == device[:len(key)] {
			function(device)
		}
	}
}

func RegisterExpireFunc(keyPrefix string, function ExpireKeyCallBack) {
	expireFuncs[keyPrefix] = function
}
