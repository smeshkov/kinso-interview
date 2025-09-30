package server

import (
	"fmt"
	"net/http"

	"github.com/smeshkov/kinso-interview/app/config"
)

func Run(cfg *config.Config, handler http.Handler) error {
	srv := &http.Server{
		ReadHeaderTimeout: cfg.Server.ReadTimeout,
		IdleTimeout:       cfg.Server.IdleTimeout,
		ReadTimeout:       cfg.Server.ReadTimeout,
		WriteTimeout:      cfg.Server.WriteTimeout,
		Addr:              cfg.Server.Addr,
		Handler:           handler,
	}
	if err := srv.ListenAndServe(); err != nil {
		return fmt.Errorf("error in listen and serve: %w", err)
	}
	return nil
}
