package logger

import (
	"bytes"
	"log"
	"strings"
	"testing"
)

func TestLogger(t *testing.T) {
	tests := []struct {
		testName     string
		logFunc      func(...interface{})
		initLogLevel string
		expected     string
		message      string
	}{
		{"DebugLog", Debug, "DEBUG", "DEBUG", "test debug message"},
		{"InfoLog", Info, "INFO", "INFO", "test info message"},
		{"ErrorLog", Error, "ERROR", "ERROR", "test error message"},
	}

	for _, tt := range tests {
		t.Run(tt.testName, func(t *testing.T) {
			var buf bytes.Buffer
			logger = log.New(&buf, "", 0)
			Init(tt.initLogLevel, logger)
			tt.logFunc(tt.message)

			// Use strings.Contains to check if the buffer contains the expected string
			logOutput := buf.String()
			if !strings.Contains(logOutput, tt.expected) {
				t.Errorf("expected log to contain %q, got %q", tt.expected, logOutput)
			}
		})
	}
}

func TestParseLogLevel(t *testing.T) {
	tests := []struct {
		input    string
		expected LogLevel
	}{
		{"DEBUG", DEBUG},
		{"INFO", INFO},
		{"ERROR", ERROR},
		{"FATAL", FATAL},
		{"UNKNOWN", INFO}, // Default to INFO on unknown input
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := parseLogLevel(tt.input)
			if result != tt.expected {
				t.Errorf("parseLogLevel(%q) = %v, want %v", tt.input, result, tt.expected)
			}
		})
	}
}
