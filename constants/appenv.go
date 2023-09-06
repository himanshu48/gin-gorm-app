package constants

import (
	"fmt"
	"log"
	"strings"
)

type AppEnvironment uint32

func (env AppEnvironment) String() string {
	if b, err := env.marshalText(); err == nil {
		return string(b)
	} else {
		return "unknown"
	}
}

func ParseAppEnviroment(env string) AppEnvironment {
	switch strings.ToLower(env) {
	case "development", "dev":
		return DevelopmentEnv
	case "test":
		return TestEnv
	case "production", "prod":
		return ProductionEnv
	}

	log.Fatalf("not a valid app environment: %q", env)

	var l AppEnvironment
	return l
}

func (environment AppEnvironment) marshalText() ([]byte, error) {
	switch environment {
	case DevelopmentEnv:
		return []byte("development"), nil
	case TestEnv:
		return []byte("test"), nil
	case ProductionEnv:
		return []byte("production"), nil
	}

	return nil, fmt.Errorf("not a valid environment %d", environment)
}

var AllAppEnvironments = []AppEnvironment{
	DevelopmentEnv,
	TestEnv,
	ProductionEnv,
}

const (
	DevelopmentEnv AppEnvironment = iota
	TestEnv
	ProductionEnv
)
