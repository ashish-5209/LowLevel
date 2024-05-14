package loggingSystem

// LogType represents the type of log.
type LogType int

const (
	// Info log type.
	Info LogType = iota
	// Warning log type.
	Warning
	// Error log type.
	Error
)

// Logger interface defines the method to handle log messages.
type Logger interface {
	Next(logger Logger)
	Log(message string, logType LogType)
}
