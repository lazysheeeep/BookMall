package serializer

type Response struct {
	Status int         `json:"status"`
	Err    string      `json:"err"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}
