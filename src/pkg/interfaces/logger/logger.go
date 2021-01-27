package logger

import (
	"fmt"
	"os"
	"strings"
	"time"
)

// Logger ...
type (
	Logger interface {
		Debug(f string, a ...interface{})
		Info(f string, a ...interface{})
		Warn(f string, a ...interface{})
		Error(f string, a ...interface{})
	}
	impl struct {
		level  Level
		header string
	}
)

// NewDefaultLogger ...
func NewDefaultLogger(level Level, header string) Logger {
	return &impl{
		level:  level,
		header: header,
	}
}

func (l *impl) Debug(f string, a ...interface{}) {
	l.print(LoggerLevelDebug, f, a...)
}

func (l *impl) Info(f string, a ...interface{}) {
	l.print(LoggerLevelInfo, f, a...)
}

func (l *impl) Warn(f string, a ...interface{}) {
	l.print(LoggerLevelWarn, f, a...)
}

func (l *impl) Error(f string, a ...interface{}) {
	l.print(LoggerLevelError, f, a...)
}

func (l *impl) print(level Level, f string, a ...interface{}) {
	if l.level > level {
		return
	}
	now := time.Now().Format(time.RFC3339)
	msg := fmt.Sprintf(f, a...)
	lvl := fmt.Sprintf("%-5s", level)
	log := strings.Join([]string{now, lvl, l.header, msg}, "|")
	fmt.Fprintln(os.Stderr, log)
}
