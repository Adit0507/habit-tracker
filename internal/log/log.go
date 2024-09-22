package log

import (
	"io"
	"log"
	"sync"
)

// log messages
type Logger struct {
	mutex sync.Mutex
	logger *log.Logger
}

func New(output io.Writer) *Logger {
	return &Logger{
		logger: log.New(output, "", log.Ldate|log.Ltime),
	}
}

func (l*Logger) Logf(format string, args ...any) {
	l.mutex.Lock()
	defer l.mutex.Unlock()
	l.logger.Printf(format, args...)
}