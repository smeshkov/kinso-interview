package server

import (
	"context"
	"fmt"
	"net/http"
)

// AppError ...
type AppError struct {
	Err     error
	Message string
	Code    int
	Ctx     context.Context
}

func (e *AppError) Error() string {
	return fmt.Sprintf("application error [%d]: %s - %v", e.Code, e.Message, e.Err)
}

// StatusBadRequest - HTTP 400
func StatusBadRequest(c context.Context, message string, v ...interface{}) *AppError {
	return &AppError{
		Ctx:     c,
		Message: fmt.Sprintf(message, v...),
		Code:    http.StatusBadRequest,
	}
}
