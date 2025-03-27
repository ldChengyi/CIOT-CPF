package vo

// 通用响应结构
type Response struct {
	Code int         `json:"code"`           // 状态码，200 为成功，其他为失败
	Msg  string      `json:"msg"`            // 消息
	Data interface{} `json:"data,omitempty"` // 可选字段，为空时不包含在 JSON 中
}

// 返回成功的响应
func Success(msg string, data interface{}) Response {
	return Response{
		Code: 200, // SUCCESS 状态码
		Msg:  msg,
		Data: data, // 如果 data 为 nil 或空，JSON 中不返回该字段
	}
}

// 返回失败的响应
func Fail(msg string) Response {
	return Response{
		Code: 400, // FAILURE 状态码
		Msg:  msg,
		Data: nil, // 失败时一般没有 Data
	}
}
