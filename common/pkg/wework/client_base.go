package wework

import (
	"net/http"
)

type wework struct {
	corpId     string // 企业ID
	corpSecret string // 应用的凭证密钥，注意应用需要是启用状态
}

type Client struct {
	http        *http.Client // http
	accessToken *accessToken // access
	wework      *wework      // 企微基础配置
}

type requestInfo struct {
	url    string
	method string
}
