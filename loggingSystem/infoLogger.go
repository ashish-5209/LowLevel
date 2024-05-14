package loggingSystem

// InfoLogger represents the info logger.
type InfoLogger struct {
	BaseLogger
}

// Log logs the message with info log type.
func (il *InfoLogger) Log(message string, logType LogType) {
	if logType == Info {
		il.BaseLogger.Log(message, logType)
	} else {
		il.BaseLogger.Log(message, logType)
	}
}
