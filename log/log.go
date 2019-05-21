package log

import (
	"os"

	"github.com/go-kit/kit/log"
)

// Logger defines a method for writing application logs
type Logger interface {
	Log(keyvals ...interface{}) error
}

// NewLogger initializes a new Logger
func NewLogger() Logger {
	return newLogger()
}

func newLogger() Logger {
	return log.NewJSONLogger(log.NewSyncWriter(os.Stdout))
}
