package web

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/cry999/pm-projects/pkg/interfaces/logger"
)

type (
	impl struct {
		level  logger.Level
		header string
	}
)

// NewRequestLogger ...
func NewRequestLogger(r *http.Request, level logger.Level) logger.Logger {
	return &impl{
		level:  level,
		header: strings.Join([]string{r.Method, r.URL.Path}, "|"),
	}
}

// NewDefaultLogger ...
func NewDefaultLogger(level logger.Level) logger.Logger {
	return &impl{
		level:  level,
		header: "(default-logger)",
	}
}

func (l *impl) Debug(f string, a ...interface{}) {
	l.print(logger.LoggerLevelDebug, f, a...)
}

func (l *impl) Info(f string, a ...interface{}) {
	l.print(logger.LoggerLevelInfo, f, a...)
}

func (l *impl) Warn(f string, a ...interface{}) {
	l.print(logger.LoggerLevelWarn, f, a...)
}

func (l *impl) Error(f string, a ...interface{}) {
	l.print(logger.LoggerLevelError, f, a...)
}

func (l *impl) print(level logger.Level, f string, a ...interface{}) {
	if l.level > level {
		return
	}
	now := time.Now().Format(time.RFC3339)
	msg := fmt.Sprintf(f, a...)
	lvl := fmt.Sprintf("%-5s", level)
	log := strings.Join([]string{now, lvl, l.header, msg}, "|")
	fmt.Fprintln(os.Stderr, log)
}
