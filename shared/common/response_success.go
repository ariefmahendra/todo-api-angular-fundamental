package common

import (
	"encoding/json"
	"github.com/ariefmahendra/crud-api-article/model/dto"
	"net/http"
)

func ResponseSuccess(w http.ResponseWriter, statusCode int, status string, message string, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)

	response := dto.Response{
		Code:    statusCode,
		Status:  status,
		Message: message,
		Data:    data,
	}
	json.NewEncoder(w).Encode(response)
}
