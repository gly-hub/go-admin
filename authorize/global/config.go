package global

import "time"

var BaseConfig *CustomConfig

// CustomConfig 自定义配置
type CustomConfig struct {
	Jwt *Jwt `json:"jwt" yaml:"jwt"`
}

type Jwt struct {
	Key    string        `json:"key" yaml:"key"`
	Expire time.Duration `json:"expire" yaml:"expire"`
}
