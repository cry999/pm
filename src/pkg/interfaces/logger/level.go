package logger

// Level ...
type Level int

// LoggerLevel constant variables
const (
	LoggerLevelDebug Level = iota
	LoggerLevelInfo
	LoggerLevelWarn
	LoggerLevelError
)

// String ...
func (ll Level) String() string {
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
