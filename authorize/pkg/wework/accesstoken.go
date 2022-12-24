package wework

import (
	"fmt"
	"time"
)

type accessToken struct {
	accessKey string    // access_token
	createdAt time.Time // 创建时间
	expiresIn int32     // 有效期 单位s
}

func (at *accessToken) verifyTime() bool {
	if at.accessKey == "" || at.expiresIn == 0 {
		return false
	}

	if at.createdAt.Add(time.Duration(at.expiresIn) * time.Second).After(time.Now()) {
		return true
	}
	return false
}

// getAccessTokenURL 获取请求access_token url
func (client *Client) getAccessTokenURL() requestInfo {
	return requestInfo{
		url:    fmt.Sprintf("%s/gettoken?corpid=%s&corpsecret=%s", BaseUrl, client.wework.corpId, client.wework.corpSecret),
		method: GET,
	}
}

func (client *Client) checkAccessToken() error {
	if client.accessToken == nil {
		client.accessToken = &accessToken{
			accessKey: "",
			createdAt: time.Now(),
			expiresIn: 0,
		}
	}
	// 验证成功则不用重新获取
	if client.accessToken.verifyTime() {
		return nil
	}

	// 验证未通过需要重新获取
	var respData = &getAccessTokenResp{}
	err := client.doAction(client.getAccessTokenURL(), nil, respData)
	if err != nil {
		return err
	}

	client.accessToken.createdAt = time.Now()
	client.accessToken.expiresIn = respData.ExpiresIn
	client.accessToken.accessKey = respData.AccessToken
	return nil
}
