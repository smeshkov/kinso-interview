package app

import (
	"context"
	"flag"
	"fmt"

	"github.com/smeshkov/kinso-interview/app/config"
	"github.com/smeshkov/kinso-interview/app/ctx"
	"github.com/smeshkov/kinso-interview/app/handlers"
	"github.com/smeshkov/kinso-interview/app/logger"
	"github.com/smeshkov/kinso-interview/app/server"
)

var (
	cfgFile = flag.String("config", "_resources/config.yml", "configuration YAML")
)

func Run(run *config.RuntimeConfig) error {
	// Runtime configuration
	run.ApplyEnvVars()

	// Global logger
	log := logger.NewLog(run.EnvName, run.Version, run.InstanceGroup)

	flag.Parse()

	// Static configuration
	cfg, err := config.Load(*cfgFile)
	if err != nil {
		return fmt.Errorf("error in loading configuration [%s]: %w", *cfgFile, err)
	}

	// Global context
	c, cancel := context.WithCancel(context.Background())
	defer cancel()
	log.Info(config.InfoString(&cfg, run))

	// App context
	ctx.Setup()

	// Server
	handler, err := handlers.New(c, &cfg, run, log)
	if err != nil {
		return fmt.Errorf("error in setting up handlers: %w", err)
	}
	if err := server.Run(&cfg, handler); err != nil {
		return fmt.Errorf("error in running server: %w", err)
	}

	return nil
}
