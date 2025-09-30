package config

import (
	"fmt"
	"os"
	"time"

	"gopkg.in/yaml.v2"
)

// Config ...
type Config struct {
	Server struct {
		Name         string
		Addr         string
		ReadTimeout  time.Duration
		WriteTimeout time.Duration
		IdleTimeout  time.Duration
	}
	Env map[string]*EnvConfig
}

// Load loads configuration from file.
func Load(file string) (cfg Config, err error) {
	cfg.Server.Addr = ":8080"
	cfg.Server.ReadTimeout = 5 * time.Second
	cfg.Server.WriteTimeout = 5 * time.Second
	cfg.Server.IdleTimeout = 5 * time.Second

	data, err := os.ReadFile(file)
	if err != nil {
		return
	}

	if err = yaml.Unmarshal(data, &cfg); err != nil {
		return
	}

	return
}

func (cfg *Config) GetEnv(envName string) *EnvConfig {
	if v, ok := cfg.Env[envName]; ok {
		return v
	}
	return cfg.Env[EnvProd]
}

// InfoString returns configuration details as a string.
func InfoString(cfg *Config, run *RuntimeConfig) string {
	return fmt.Sprintf("server info: name = %q, version = %q, env name = %q, instance group = %q, addr = %q, read timeout = %q, write timeout = %q, idle timeout = %q",
		cfg.Server.Name, run.Version, run.EnvName, run.InstanceGroup, cfg.Server.Addr, cfg.Server.ReadTimeout, cfg.Server.WriteTimeout, cfg.Server.IdleTimeout)
}

type EnvConfig struct {
	EnvName   string
	QueueAddr string
}
