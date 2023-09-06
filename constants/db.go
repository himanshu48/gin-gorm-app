package constants

import (
	"fmt"
	"log"
	"strings"
)

type DatabaseType uint32

func (env DatabaseType) String() string {
	if b, err := env.marshalText(); err == nil {
		return string(b)
	} else {
		return "unknown"
	}
}

func ParseDatabaseType(env string) DatabaseType {
	switch strings.ToLower(env) {
	case "mysql":
		return MySql
	case "postgres", "pg":
		return Postgres
	case "sqlite3", "sqlite":
		return Sqlite
	}

	log.Fatalf("not a valid database type: %q", env)

	var l DatabaseType
	return l
}

func (dbType DatabaseType) marshalText() ([]byte, error) {
	switch dbType {
	case MySql:
		return []byte("mysql"), nil
	case Postgres:
		return []byte("postgres"), nil
	case Sqlite:
		return []byte("sqlite3"), nil
	}

	return nil, fmt.Errorf("not a valid database type %d", dbType)
}

var AllDatabaseTypes = []DatabaseType{
	MySql,
	Postgres,
	Sqlite,
}

const (
	MySql DatabaseType = iota
	Postgres
	Sqlite
)
