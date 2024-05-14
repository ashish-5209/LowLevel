package loggingSystem

// ErrorLogger represents the error logger.
type ErrorLogger struct {
	BaseLogger
}

// Log logs the message with error log type.
func (el *ErrorLogger) Log(message string, logType LogType) {
	if logType == Error {
		el.BaseLogger.Log(message, logType)
	} else {
		el.BaseLogger.Log(message, logType)
	}
}
