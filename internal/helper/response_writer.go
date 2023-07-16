package helper

import (
	"encoding/json"
	"net/http"

	"github.com/doffy007/rest-api-makanan/internal/response"
)

func ResponseWriter(res http.ResponseWriter, statusCode int, message string, data interface{}) error {
	res.WriteHeader(statusCode)
	httpResponse := response.NewResponse(statusCode, message, data)
	err := json.NewEncoder(res).Encode(httpResponse)
	return err
}
