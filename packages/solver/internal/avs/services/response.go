package services

type Response struct {
	Data    any `json:"data"`
	Error   any `json:"error"`
	Message any `json:"message"`
}

func NewResponse(data any, message string) Response {
	return Response{Data: data, Error: false, Message: message}
}
