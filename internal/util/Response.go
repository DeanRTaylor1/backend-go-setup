package util

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/deanrtaylor1/backend-go/internal/config"
	"github.com/go-playground/validator/v10"
)

type APIErrorResponse struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Stack   string      `json:"stack,omitempty"`
}

type APISuccessResponse struct {
	Status  bool        `json:"status"`
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

func NewAPIErrorResponse(status bool, code int, message string, data interface{}, err error, config config.EnvConfig) APIErrorResponse {
	stack := ""
	if config.IsDevelopment {
		if err != nil {
			stack = fmt.Sprintf("%+v", err)
		}
	}
	return APIErrorResponse{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    data,
		Stack:   stack,
	}
}

func NewAPISuccessResponse(status bool, code int, message string, data interface{}) APISuccessResponse {
	return APISuccessResponse{
		Status:  status,
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func SendResponse(w http.ResponseWriter, statusCode int, response interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func ValidateStruct(validate *validator.Validate, params interface{}, config config.EnvConfig) (error, APIErrorResponse) {
	if err := validate.Struct(params); err != nil {
		errors := err.(validator.ValidationErrors)

		errorMessages := make([]string, len(errors))
		for i, err := range errors {
			errorMessages[i] = fmt.Sprintf("Field %s failed validation (%s)", err.Field(), err.Tag())
		}

		// Create a map with the error messages
		res := NewAPIErrorResponse(false, http.StatusBadRequest, "Validation Failed.", errorMessages, err, config)

		return err, res
	}

	return nil, APIErrorResponse{}
}
