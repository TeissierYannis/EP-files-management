package config

import (
    "github.com/joho/godotenv"
    "os"
)

/*
 * Config struct
 *
 * @description Struct to hold configuration values for the application \n LogLevel can be INFO, DEBUG, ERROR, FATAL \n FilePath is the path to the directory where files will be handled
 * @var string LogLevel
 * @var string FilePath
 * @example
 *   cfg := &Config{
 *       LogLevel: "info",
 *       FilePath: "/tmp/files",
 *   }
 *   fmt.Println(cfg.LogLevel)
 *   // Output: info
 *   fmt.Println(cfg.FilePath)
 *   // Output: /tmp/files
 */
type Config struct {
	LogLevel string
}

func LoadConfig(file string) (*Config, error) {
	if file == "" {
		file = ".env"
	}
	// load .env if exists
	godotenv.Load(file)

	return &Config{
		LogLevel: os.Getenv("LOG_LEVEL"),
	}, nil
}
