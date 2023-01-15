package wework

import "fmt"

// getDepartmentURL 获取单个部门详情URL
func (client Client) getDepartmentURL(departmentId int) requestInfo {
	return requestInfo{
		url:    fmt.Sprintf("%s/department/get?access_token=%s&id=%d", BaseUrl, client.accessToken.accessKey, departmentId),
		method: GET,
	}
}

// getDepartment 获取单个部门详情
func (client Client) getDepartment(departmentId int) (info *Department, err error) {
	if err = client.checkAccessToken(); err != nil {
		return
	}

	info = &Department{}
	err = client.doAction(client.getDepartmentURL(departmentId), nil, info)
	if err != nil {
		return
	}
	return
}

// getDepartmentListURL 获取部门列表URL
func (client Client) getDepartmentListURL(departmentId int) requestInfo {
	return requestInfo{
		url:    fmt.Sprintf("%s/department/list?access_token=%s&id=%d", BaseUrl, client.accessToken.accessKey, departmentId),
		method: GET,
	}
}

// getDepartmentList 获取部门列表
func (client Client) getDepartmentList(departmentId int) (list []Department, err error) {
	if err = client.checkAccessToken(); err != nil {
		return
	}

	var response = &getDepartmentListResp{}
	err = client.doAction(client.getDepartmentURL(departmentId), nil, response)
	if err != nil {
		return
	}
	list = response.Department
	return
}

// getDepartmentSimpleListURL 获取子部门ID列表URL
func (client Client) getDepartmentSimpleListURL(departmentId int) requestInfo {
	return requestInfo{
		url:    fmt.Sprintf("%s/department/simplelist?access_token=%s&id=%d", BaseUrl, client.accessToken.accessKey, departmentId),
		method: GET,
	}
}

// getDepartmentSimpleList 获取子部门ID列表
func (client Client) getDepartmentSimpleList(departmentId int) (list []DepartmentSimple, err error) {
	if err = client.checkAccessToken(); err != nil {
		return
	}

	var response = &getDepartmentSimpleListResp{}
	err = client.doAction(client.getDepartmentURL(departmentId), nil, response)
	if err != nil {
		return
	}
	list = response.Department
	return
}
