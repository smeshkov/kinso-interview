package logger

import (
	"log/slog"
	"os"

	"github.com/smeshkov/kinso-interview/app/config"
)

func NewLog(env, version, instanceGroup string) *slog.Logger {
	isLocal := env == config.EnvLocal
	isProd := env == config.EnvProd

	opts := &slog.HandlerOptions{}
	if isProd {
		opts.Level = slog.LevelInfo
	} else {
		opts.Level = slog.LevelDebug
	}

	var log *slog.Logger
	if isLocal {
		log = slog.New(slog.NewTextHandler(os.Stdout, opts))
	} else {
		log = slog.New(slog.NewJSONHandler(os.Stdout, opts))
	}
	log = log.With("env", env, "version", version, "instance_group", instanceGroup)

	slog.SetDefault(log)

	return log
}
