package server

import (
	"context"
	"fmt"
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
