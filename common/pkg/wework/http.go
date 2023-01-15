package wework

import (
	"bytes"
	"errors"
	jsoniter "github.com/json-iterator/go"
	"io/ioutil"
	"net/http"
)

func (client Client) doAction(info requestInfo, payload interface{}, reply interface{}) (err error) {
	var (
		request  *http.Request
		response *http.Response
	)
	if info.method == "GET" {
		request, err = http.NewRequest(info.method, info.url, nil)
	} else {
		var data string
		data, err = jsoniter.MarshalToString(payload)
		if err != nil {
			return
		}
		body := new(bytes.Buffer)
		body.Write([]byte(data))
		request, err = http.NewRequest(info.method, info.url, body)
	}
	response, err = client.http.Do(request)
	if err != nil {
		return
	}
	defer response.Body.Close()

	var respBody []byte
	respBody, err = ioutil.ReadAll(response.Body)
	if err != nil {
		return err
	}

	var errResp = &Error{}
	err = jsoniter.Unmarshal(respBody, errResp)
	if err != nil {
		return
	}

	if errResp.ErrCode != 0 {
		return errors.New(errResp.ErrMsg)
	}

	err = jsoniter.Unmarshal(respBody, reply)
	if err != nil {
		return
	}

	return nil
}
