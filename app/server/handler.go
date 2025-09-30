package server

import (
	"log/slog"
	"net/http"

	"github.com/go-chi/httplog/v2"
)

// AppHandler ...
// http://blog.golang.org/error-handling-and-go
type AppHandler func(http.ResponseWriter, *http.Request) *AppError

// ServeHTTP ...
func (fn AppHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if e := fn(w, r); e != nil { // e is *appError, not os.Error.
		c := e.Ctx
		log := httplog.LogEntry(c)
		log = log.With("http_code", e.Code, "message", e.Message)

		if e.Code < 400 || e.Code >= 500 {
			log.Error(e.Error())
		} else if log.Enabled(c, slog.LevelDebug) {
			log.Warn(e.Error())
		}

		http.Error(w, e.Message, e.Code)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func AppHandlerFunc(h AppHandler) http.HandlerFunc {
	return h.ServeHTTP
}
