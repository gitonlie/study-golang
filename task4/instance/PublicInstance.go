package instance

type SuccessResponse struct {
	Code    int         `json:"code" example:"200"`   // 状态码
	Message string      `json:"message" example:"成功"` // 消息
	Data    interface{} `json:"data"`
}

type FailedResponse struct {
	Code    int    `json:"code" example:"400"`   // 状态码
	Message string `json:"message" example:"失败"` // 消息
}

type UnauthorizedResponse struct {
	Code    int    `json:"code" example:"401"`     // 状态码
	Message string `json:"message" example:"权限不足"` // 消息
}

type ErrorResponse struct {
	Code    int    `json:"code" example:"500"`                      // 状态码
	Message string `json:"message" example:"Internal Server Error"` // 消息
	Err     string `json:"error"`
}
