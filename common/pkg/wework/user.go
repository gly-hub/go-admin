package wework

import "fmt"

// getUserListURL 读取成员URL
func (client Client) getUserURL(userId string) requestInfo {
	return requestInfo{
		url:    fmt.Sprintf("%s/user/get?access_token=%s&userid=%s", BaseUrl, client.accessToken.accessKey, userId),
		method: GET,
	}
}

// getUser 读取成员
func (client Client) getUser(userId string) (info *User, err error) {
	if err = client.checkAccessToken(); err != nil {
		return
	}

	info = &User{}
	err = client.doAction(client.getUserURL(userId), nil, info)
	if err != nil {
		return
	}
	return
}

// getUserSimpleListURL 获取部门成员URL
func (client Client) getUserSimpleListURL(departmentId int) requestInfo {
	return requestInfo{
		url:    fmt.Sprintf("%s/user/simplelist?access_token=%s&department_id=%d", BaseUrl, client.accessToken.accessKey, departmentId),
		method: GET,
	}
}

// getUserSimpleList 获取部门成员列表
func (client Client) getUserSimpleList(departmentId int) (list []UserSimple, err error) {
	if err = client.checkAccessToken(); err != nil {
		return
	}

	var response = &getUserSimpleListResp{}
	err = client.doAction(client.getUserSimpleListURL(departmentId), nil, response)
	if err != nil {
		return
	}
	list = response.UserList
	return
}

// getUserListURL 获取部门成员详情URL
func (client Client) getUserListURL(departmentId int) requestInfo {
	return requestInfo{
		url:    fmt.Sprintf("%s/user/list?access_token=%s&department_id=%d", BaseUrl, client.accessToken.accessKey, departmentId),
		method: GET,
	}
}

// getUserList 获取部门成员详情列表
func (client Client) getUserList(departmentId int) (list []User, err error) {
	if err = client.checkAccessToken(); err != nil {
		return
	}

	var response = &getUserListResp{}
	err = client.doAction(client.getUserSimpleListURL(departmentId), nil, response)
	if err != nil {
		return
	}
	list = response.UserList
	return
}
