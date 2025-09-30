package handlers

import (
	"log/slog"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/httplog/v2"
	"github.com/go-chi/httprate"

	"github.com/smeshkov/kinso-interview/app/config"
)

func SetupMiddleware(r chi.Router, run *config.RuntimeConfig, appName string) {
	// ------------------------ [START] Middleware ------------------------
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(httplog.RequestLogger(newLog(run, appName)))
	r.Use(middleware.Recoverer)

	// Set a timeout value on the request context (ctx), that will signal
	// through ctx.Done() that the request has timed out and further
	// processing should be stopped.
	r.Use(middleware.Timeout(60 * time.Second))

	// Enable httprate request limiter of 100 requests per minute.
	//
	// In the code example below, rate-limiting is bound to the request IP address
	// via the LimitByIP middleware handler.
	//
	// To have a single rate-limiter for all requests, use httprate.LimitAll(..).
	//
	// Please see _example/main.go for other more, or read the library code.
	var allowedRate int
	if run.IsLocal() {
		allowedRate = 1000
	} else {
		allowedRate = 100
	}
	r.Use(httprate.LimitByIP(allowedRate, time.Minute))

	// Respond to App Engine and Compute Engine health checks.
	// Indicate the server is healthy.
	r.Use(middleware.Heartbeat("/_ah/health"))
	r.Use(middleware.Heartbeat("/ping"))
	// ------------------------ [END] Middleware ------------------------
}

var (
	quietDownRoutes = []string{
		"/",
		"/info",
		"/ping",
		"/_ah/health",
	}
)

func newLog(run *config.RuntimeConfig, appName string) *httplog.Logger {
	var logLevel slog.Level
	if run.IsProd() {
		logLevel = slog.LevelInfo
	} else {
		logLevel = slog.LevelDebug
	}
	// Logger
	return httplog.NewLogger(appName, httplog.Options{
		JSON:             !run.IsLocal(),
		LogLevel:         logLevel,
		Concise:          run.IsLocal(),
		RequestHeaders:   true,
		MessageFieldName: "message",
		// TimeFieldFormat: time.RFC850,
		Tags: map[string]string{
			"version":        run.Version,
			"env":            run.EnvName,
			"instance_group": run.InstanceGroup,
		},
		QuietDownRoutes: quietDownRoutes,
		QuietDownPeriod: 10 * time.Second,
		// SourceFieldName: "source",
	})
}
