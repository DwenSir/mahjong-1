package router

type Reply struct {
	Code int `json:"code"`
	Error string `json:"error"`
	Data interface{} `json:"data,omitempty"`
}

const(
	ErrNoCode = "没有用户code"
	ErrSuccess = "成功"
)

const(
	CodeSuccess = 0
	CodeNoCode = 11
)

