// +--------------------------------------+
// |               Logger                 |
// +--------------------------------------+
// |                                      |
// +--------------------------------------+
// | + Next(logger Logger)                |
// | + Log(message string, logType LogType)|
// +--------------------------------------+

// +--------------------------------------+
// |            BaseLogger                |
// +--------------------------------------+
// | - next: Logger                       |
// +--------------------------------------+
// | + Next(logger Logger)                |
// | + Log(message string, logType LogType)|
// +--------------------------------------+

// +--------------------------------------+
// |            InfoLogger                |
// +--------------------------------------+
// | - BaseLogger                         |
// +--------------------------------------+
// | + Log(message string, logType LogType)|
// +--------------------------------------+

// +--------------------------------------+
// |         WarningLogger                |
// +--------------------------------------+
// | - BaseLogger                         |
// +--------------------------------------+
// | + Log(message string, logType LogType)|
// +--------------------------------------+

// +--------------------------------------+
// |          ErrorLogger                 |
// +--------------------------------------+
// | - BaseLogger                         |
// +--------------------------------------+
// | + Log(message string, logType LogType)|
// +--------------------------------------+

// +--------------------------------------+
// |            Enumerations              |
// +--------------------------------------+
// | - LogType                            |
// +--------------------------------------+

package loggingSystem

func App() {
	errorLogger := &ErrorLogger{}
	warningLogger := &WarningLogger{}
	infoLogger := &InfoLogger{}

	warningLogger.Next(errorLogger)
	infoLogger.Next(warningLogger)

	infoLogger.Log("This is an information", Info)
	infoLogger.Log("This is a warning", Warning)
	infoLogger.Log("This is an error", Error)
}
