package common

type Response struct {
	Code int32 `json:"-"`
	Msg string `json:"-"`
}
