package config

import (
	"encoding/json"
	"log"
	"os"
)

// Config holds the application configuration
type Config struct {
	LogLevel    string `json:"log_level"`
	ServerPort  string `json:"server_port"`
	APIKey      string `json:"api_key"`
	RegistryURI string `json:"registry_uri"`
}

// AppConfig is a global variable that holds the loaded configuration
var AppConfig Config

// LoadConfig loads configuration from a JSON file
func LoadConfig(configFile string) {
	log.Println("", "Loading Static Configuration...")

	file, err := os.Open(configFile)
	if err != nil {
		log.Fatalf("Failed to open config file: %v", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	if err := decoder.Decode(&AppConfig); err != nil {
		log.Fatalf("Failed to decode config file: %v", err)
	}

	// Override with environment variables if available
	overrideWithEnv()
	log.Printf("Static Configuration: \n\t - Log level: %s \n\t - Server port: %s \n\t - API key: %s \n\t - Registry URI: %s",
		AppConfig.LogLevel, AppConfig.ServerPort, AppConfig.APIKey, AppConfig.RegistryURI)

	log.Println("", "Configuration loaded successfully.")
}

// overrideWithEnv checks for environment variables and overrides config values
func overrideWithEnv() {
	if port := os.Getenv("WORKER_SERVER_PORT"); port != "" {
		AppConfig.ServerPort = port
	}
	if logLevel := os.Getenv("WORKER_LOG_LEVEL"); logLevel != "" {
		AppConfig.LogLevel = logLevel
	}
	if apiKey := os.Getenv("WORKER_API_KEY"); apiKey != "" {
		AppConfig.APIKey = apiKey
	}
	if registryUri := os.Getenv("REGISTRY_URI"); registryUri != "" {
		AppConfig.RegistryURI = registryUri
	}
}
