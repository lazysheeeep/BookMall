package serializer

import "BookMall/pkg/e"

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

func BuildListResponse(items interface{}, total uint) Response {
	return Response{
		Status: 200,
		Msg:    e.GetMsg(200),
		Data: ListData{
			List:  items,
			Total: total,
		},
	}
}
