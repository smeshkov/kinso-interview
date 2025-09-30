package config

import "os"

const (
	// Instance groups names
	WebService = "web-service"

	// Environment names
	EnvLocal = "local"
	EnvStg   = "stg"
	EnvProd  = "prod"
)

// RuntimeConfig ...
type RuntimeConfig struct {
	Version       string
	EnvName       string
	InstanceGroup string
}

func NewRuntime(
	version,
	envName,
	instanceGroup string,
) *RuntimeConfig {
	run := &RuntimeConfig{
		Version:       version,
		EnvName:       envName,
		InstanceGroup: instanceGroup,
	}
	if run.EnvName == "" {
		run.EnvName = EnvLocal
	}
	if run.InstanceGroup == "" {
		run.InstanceGroup = WebService
	}
	return run
}

func (run *RuntimeConfig) ApplyEnvVars() {
	// App
	if os.Getenv("ENV_NAME") != "" {
		run.EnvName = os.Getenv("ENV_NAME")
	}
	if run.EnvName == "" {
		run.EnvName = EnvProd
	}
	if val := os.Getenv("INSTANCE_GROUP"); val != "" {
		run.InstanceGroup = val
	}
}

func (run *RuntimeConfig) IsWebService() bool {
	return run.InstanceGroup == WebService
}

func (run *RuntimeConfig) IsLocal() bool {
	return IsLocalEnv(run.EnvName)
}

func (run *RuntimeConfig) IsStg() bool {
	return IsStgEnv(run.EnvName)
}

func (run *RuntimeConfig) IsProd() bool {
	return IsProdEnv(run.EnvName)
}

func IsLocalEnv(envName string) bool {
	return envName == EnvLocal
}

func IsStgEnv(envName string) bool {
	return envName == EnvStg
}

func IsProdEnv(envName string) bool {
	return envName == EnvProd
}
