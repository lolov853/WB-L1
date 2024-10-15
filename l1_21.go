package main

import "fmt"

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct{}

func (cl *ConsoleLogger) LogMessage(message string) {
	fmt.Println("ConsoleLogger:", message)
}

type LoggerAdapter struct {
	consoleLogger *ConsoleLogger
}

func (la *LoggerAdapter) Log(message string) {
	la.consoleLogger.LogMessage(message)
}

func main() {

	consoleLogger := &ConsoleLogger{}

	loggerAdapter := &LoggerAdapter{consoleLogger: consoleLogger}

	loggerAdapter.Log("This is a log message.")
}
