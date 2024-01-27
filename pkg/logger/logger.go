package logger

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"
)

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	ERROR
	FATAL
)

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

func Init(level string) {
	logger = log.New(os.Stdout, "", 0)
	logLevel = parseLogLevel(level)
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