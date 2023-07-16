package response

import (
	"net/http"

	"github.com/doffy007/rest-api-makanan/internal/constants"
)

type BaseResponse struct {
	StatusCode int
	Message    string
	Data       interface{}
	Error      []string
}

func (b BaseResponse) InternalServerError(err string) BaseResponse {
	b.StatusCode = http.StatusInternalServerError
	b.Message = constants.INTERNAL_SERVER_ERROR
	b.Data = false
	b.Error = []string{err}

	return b
}

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Content interface{} `json:"content"`
}

func NewResponse(status int, message string, content interface{}) *Response {
	return &Response{
		Status:  status,
		Message: message,
		Content: content,
	}
}
