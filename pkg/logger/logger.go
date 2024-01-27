package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type LogLevel int

var (
	logger    *log.Logger
	logLevel  LogLevel
	logLevels = map[string]LogLevel{
		"DEBUG": DEBUG,
		"INFO":  INFO,
		"ERROR": ERROR,
		"FATAL": FATAL,
	}
)

const (
	DEBUG LogLevel = iota
	INFO
	ERROR
	FATAL
)

// Init initializes the logging system. It now accepts an optional logger parameter.
// If no logger is provided, it defaults to creating a logger that writes to os.Stdout.
func Init(level string, customLoggers ...*log.Logger) {
	// Determine the log level
	logLevel = parseLogLevel(level)

	// If a custom logger is provided, use it. Otherwise, create a default logger.
	if len(customLoggers) > 0 {
		logger = customLoggers[0]
	} else {
		logger = log.New(os.Stdout, "", 0)
	}
}

func Debug(v ...interface{}) {
	if logLevel <= DEBUG {
		output(DEBUG, v...)
	}
}

func Info(v ...interface{}) {
	if logLevel <= INFO {
		output(INFO, v...)
	}
}

func Error(v ...interface{}) {
	if logLevel <= ERROR {
		output(ERROR, v...)
	}
}

func Fatal(v ...interface{}) {
	if logLevel <= FATAL {
		output(FATAL, v...)
		os.Exit(1)
	}
}

func output(level LogLevel, v ...interface{}) {
	prefix := fmt.Sprintf("[%s][%s] ", time.Now().Format("2006-01-02 15:04:05"), levelToString(level))
	logger.SetPrefix(prefix)

	// Using fmt.Sprintf to format the message
	formattedMessage := fmt.Sprint(v...)
	logger.Println(formattedMessage)
}

func parseLogLevel(level string) LogLevel {
	l, ok := logLevels[strings.ToUpper(level)]
	if !ok {
		return INFO // Default log level
	}
	return l
}

func levelToString(level LogLevel) string {
	for k, v := range logLevels {
		if v == level {
			return k
		}
	}
	return ""
}
