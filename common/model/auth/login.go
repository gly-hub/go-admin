package auth

type LoginParams struct {
	UserName string `json:"user_name"`
	Password string `json:"password"`
}

type LoginResp struct {
	Code int
	Msg string
	Data interface{}
}
