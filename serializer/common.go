package serializer

type Response struct {
	Status int         `json:"status"`
	Err    string      `json:"err"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

type TokenData struct {
	User  UserVO `json:"user"`
	Token string `json:"token"`
}

type ListData struct {
	List  interface{} `json:"list"`
	Total uint        `json:"total"`
}
