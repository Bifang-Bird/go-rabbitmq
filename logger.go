package rabbitmq

import (
	"fmt"
	"log"

	"go-rabbitmq/internal/logger"
)

// Logger is describes a logging structure. It can be set using
// WithPublisherOptionsLogger() or WithConsumerOptionsLogger().
type Logger logger.Logger

const loggingPrefix = "gorabbit"

type stdDebugLogger struct{}

func (l stdDebugLogger) Tracef(s string, i ...interface{}) {
	//TODO implement me
	panic("implement me")
}

// Fatalf -
func (l stdDebugLogger) Fatalf(format string, v ...interface{}) {
	log.Fatalf(fmt.Sprintf("%s FATAL: %s", loggingPrefix, format), v...)
}

// Errorf -
func (l stdDebugLogger) Errorf(format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("%s ERROR: %s", loggingPrefix, format), v...)
}

// Warnf -
func (l stdDebugLogger) Warnf(format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("%s WARN: %s", loggingPrefix, format), v...)
}

// Infof -
func (l stdDebugLogger) Infof(format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("%s INFO: %s", loggingPrefix, format), v...)
}

// Debugf -
func (l stdDebugLogger) Debugf(format string, v ...interface{}) {
	log.Printf(fmt.Sprintf("%s DEBUG: %s", loggingPrefix, format), v...)
}
