package log

import (
	"os"

	"github.com/go-kit/kit/log"
)

// Logger defines a method for writing application logs
type Logger interface {
	Log(keyvals ...interface{}) error
}

// New initializes a new Logger
func New() Logger {
	return newLogger()
}

func newLogger() Logger {
	return log.NewLogfmtLogger(log.NewSyncWriter(os.Stdout))
}
