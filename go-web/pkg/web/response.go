package web

import "strconv"

type Response struct {
	Code  string      `json:"code"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewResponse(code int, data interface{}, err string) Response {
	statusCode := strconv.FormatInt(int64(code), 10)
	if code < 300 {
		return Response{
			Code:  statusCode,
			Data:  data,
			Error: "",
		}
	}

	return Response{
		Code:  statusCode,
		Data:  nil,
		Error: "",
	}
}
