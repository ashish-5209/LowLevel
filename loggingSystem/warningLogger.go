package loggingSystem

// WarningLogger represents the warning logger.
type WarningLogger struct {
	BaseLogger
}

// Log logs the message with warning log type.
func (wl *WarningLogger) Log(message string, logType LogType) {
	if logType == Warning {
		wl.BaseLogger.Log(message, logType)
	} else {
		wl.BaseLogger.Log(message, logType)
	}
}
