package server

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// WriteResponse ...
func WriteResponse(c context.Context, rw http.ResponseWriter, object any) *AppError {
	rw.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(rw).Encode(object)
	if err != nil {
		rw.Header().Set("Content-Type", "text/plain; charset=utf-8")
		return &AppError{
			Err:     fmt.Errorf("error in writing JSON response: %w", err),
			Message: fmt.Sprintf("error in response write: %v", err),
			Code:    http.StatusInternalServerError,
			Ctx:     c,
		}
	}
	return nil
}
