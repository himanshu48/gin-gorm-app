package constants

import (
	"fmt"
	"log"
	"strings"
)

type LogLevel uint32

func (level LogLevel) String() string {
	if b, err := level.marshalText(); err == nil {
		return string(b)
	} else {
		return "unknown"
	}
}

func ParseLogLevel(lvl string) LogLevel {
	switch strings.ToLower(lvl) {
	case "panic":
		return PanicLevel
	case "fatal":
		return FatalLevel
	case "error":
		return ErrorLevel
	case "warn", "warning":
		return WarnLevel
	case "info":
		return InfoLevel
	case "debug":
		return DebugLevel
	case "trace":
		return TraceLevel
	}
	log.Fatalf("not a valid log Level: %q", lvl)
	var l LogLevel
	return l
}

func (level LogLevel) marshalText() ([]byte, error) {
	switch level {
	case TraceLevel:
		return []byte("trace"), nil
	case DebugLevel:
		return []byte("debug"), nil
	case InfoLevel:
		return []byte("info"), nil
	case WarnLevel:
		return []byte("warning"), nil
	case ErrorLevel:
		return []byte("error"), nil
	case FatalLevel:
		return []byte("fatal"), nil
	case PanicLevel:
		return []byte("panic"), nil
	}

	return nil, fmt.Errorf("not a valid log level %d", level)
}

var AllLogLevels = []LogLevel{
	PanicLevel,
	FatalLevel,
	ErrorLevel,
	WarnLevel,
	InfoLevel,
	DebugLevel,
	TraceLevel,
}

const (
	PanicLevel LogLevel = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
	TraceLevel
)
