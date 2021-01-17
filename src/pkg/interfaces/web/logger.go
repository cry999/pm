package web

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

type (
	// Logger ...
	Logger interface {
		Debug(f string, a ...interface{})
		Info(f string, a ...interface{})
		Warn(f string, a ...interface{})
		Error(f string, a ...interface{})
	}
	logger struct {
		level  LoggerLevel
		header string
	}
	// LoggerLevel should be printed out
	LoggerLevel int
)

// LoggerLevel constant variables
const (
	LoggerLevelDebug LoggerLevel = iota
	LoggerLevelInfo
	LoggerLevelWarn
	LoggerLevelError
)

// String ...
func (ll LoggerLevel) String() string {
	switch ll {
	case LoggerLevelDebug:
		return "DEBUG"
	case LoggerLevelInfo:
		return "INFO"
	case LoggerLevelWarn:
		return "WARN"
	case LoggerLevelError:
		return "ERROR"
	}
	if ll < LoggerLevelDebug {
		return "DEBUG"
	}
	return "ERROR"
}

// NewRequestLogger ...
func NewRequestLogger(r *http.Request, level LoggerLevel) Logger {
	return &logger{
		level:  level,
		header: strings.Join([]string{r.Method, r.URL.Path}, "|"),
	}
}

// NewDefaultLogger ...
func NewDefaultLogger(level LoggerLevel) Logger {
	return &logger{
		level:  level,
		header: "(default-logger)",
	}
}

func (l *logger) Debug(f string, a ...interface{}) {
	l.print(LoggerLevelDebug, f, a...)
}

func (l *logger) Info(f string, a ...interface{}) {
	l.print(LoggerLevelInfo, f, a...)
}

func (l *logger) Warn(f string, a ...interface{}) {
	l.print(LoggerLevelWarn, f, a...)
}

func (l *logger) Error(f string, a ...interface{}) {
	l.print(LoggerLevelError, f, a...)
}

func (l *logger) print(level LoggerLevel, f string, a ...interface{}) {
	if l.level > level {
		return
	}
	now := time.Now().Format(time.RFC3339)
	msg := fmt.Sprintf(f, a...)
	lvl := fmt.Sprintf("%-5s", level)
	log := strings.Join([]string{now, lvl, msg}, "|")
	fmt.Fprintln(os.Stderr, log)
}
