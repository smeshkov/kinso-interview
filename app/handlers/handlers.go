package handlers

import (
	"context"
	"fmt"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"

	"github.com/smeshkov/kinso-interview/app/config"
	"github.com/smeshkov/kinso-interview/app/consumer"
	"github.com/smeshkov/kinso-interview/app/ctx"
	"github.com/smeshkov/kinso-interview/app/event"
	"github.com/smeshkov/kinso-interview/app/listener"
	"github.com/smeshkov/kinso-interview/app/server"
)

// New ...
func New(c context.Context, cfg *config.Config, run *config.RuntimeConfig, log *slog.Logger) (http.Handler, error) {
	env := cfg.GetEnv(run.EnvName)
	if env == nil {
		return nil, fmt.Errorf("environment configuration not found for [%s]", run.EnvName)
	}

	r := chi.NewRouter()

	var err error

	SetupMiddleware(r, run, "kinso-interview")

	// start queue listener in the background
	go listener.Run(c, env.QueueAddr, func(events []*event.EventDTO) error {
		return consumer.Consume(c, events)
	})

	// ------------------------ API ------------------------
	r.Route("/", func(root chi.Router) {
		// Shows current version of the App
		root.Get("/info", server.AppHandlerFunc(infoHandler(run)))

		root.Route("/api/v1/events", func(r chi.Router) {
			r.Get("/{userID}", server.AppHandlerFunc(getUserEvents))
		})
	})

	return r, err
}

func infoHandler(run *config.RuntimeConfig) func(http.ResponseWriter, *http.Request) *server.AppError {
	return func(w http.ResponseWriter, r *http.Request) *server.AppError {
		return server.WriteResponse(r.Context(), w, map[string]any{
			"version":       run.Version,
			"env":           run.EnvName,
			"instanceGroup": run.InstanceGroup,
		})
	}
}

func getUserEvents(w http.ResponseWriter, r *http.Request) *server.AppError {
	c := r.Context()

	uid := chi.URLParam(r, "userID")
	if uid == "" {
		return server.StatusBadRequest(c, "User ID is required.")
	}

	events := ctx.DB.GetByUserID(uid)

	return server.WriteResponse(c, w, map[string]any{
		"data": map[string]any{
			"events": event.ToDTOs(events),
		},
	})
}
