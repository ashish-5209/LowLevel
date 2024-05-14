package loggingSystem

import "fmt"

// BaseLogger represents the base logger in the chain.
type BaseLogger struct {
	next Logger
}

// Next sets the next logger in the chain.
func (bl *BaseLogger) Next(logger Logger) {
	bl.next = logger
}

// Log logs the message.
func (bl *BaseLogger) Log(message string, logType LogType) {
	fmt.Printf("[%s] %s\n", logType, message)
	if bl.next != nil {
		bl.next.Log(message, logType)
	}
}
