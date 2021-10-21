package model

import (
	"crud/chrono"

	val "github.com/go-playground/validator/v10"
)

type errorResponse struct {
	TimeStamp    chrono.ISO8601 `json:"timestamp"`
	ErrorCode    string         `json:"error_code"`
	ErrorMessage string         `json:"error_message"`
	ErrorDetail  interface{}    `json:"error_details,omitempty"`
}

type fieldError struct {
	Name  string
	Tag   string
	Value interface{}
}

func ErrorResponse(code string, message string, detail interface{}) errorResponse {
	return errorResponse{
		TimeStamp:    chrono.Now(),
		ErrorCode:    code,
		ErrorMessage: message,
		ErrorDetail:  detail,
	}
}

func FieldErrors(errors val.ValidationErrors) map[string]string {
	messages := make(map[string]string)
	for _, err := range errors {
		messages[err.Field()] = err.Error()
	}
	return messages
}
