package config

import (
	"log"
	"os"
	"testing"
)

func TestConfigLoad(t *testing.T) {

	// Create a sample .env file for testing
	createSampleEnvFile()

	// Test LoadConfig function
	cfg, err := LoadConfig(".env.test")
	if err != nil {
		t.Fatalf("LoadConfig failed: %v", err)
	}

	// Test LogLevel
	if cfg.LogLevel != "INFO" {
		t.Errorf("Expected LogLevel to be 'info', got '%s'", cfg.LogLevel)
	}
}

func createSampleEnvFile() {
	// Create a sample .env file for testing
	file, err := os.Create(".env.test")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.WriteString("LOG_LEVEL=INFO\n")
	if err != nil {
		log.Fatal(err)
	}
}

// Delete the sample .env file after testing
func TestMain(m *testing.M) {
	// Run the tests
	exitCode := m.Run()

	// Delete the sample .env file
	err := os.Remove(".env.test")
	if err != nil {
		log.Fatal(err)
	}

	// Exit with exit code
	os.Exit(exitCode)
}
