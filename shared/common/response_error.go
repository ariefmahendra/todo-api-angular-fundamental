package common

import (
	"encoding/json"
	"github.com/ariefmahendra/crud-api-article/model/dto"
	"net/http"
)

func ResponseError(w http.ResponseWriter, statusCode int, status string, message string) {
	w.Header().Set("Content-Type", "application/json")

	response := dto.Response{
		Code:    statusCode,
		Status:  status,
		Message: message,
		Data:    nil,
	}

	json.NewEncoder(w).Encode(response)
}
