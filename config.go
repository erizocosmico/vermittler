package main

import (
	"encoding/json"
	"fmt"
	"os"
)

// Config contains the application settings
type Config struct {
	// Enable or disable the cache
	CacheEnabled bool `json:"cache_enabled"`
	// Path where the cache is stored, without trailing slash
	CachePath string `json:"cache_path"`
	// Port where the app will run e.g: "8888"
	Port string `json:"port"`
	// Log all operations
	Verbose bool `json:"verbose"`
}

// NewConfig returns a Config struct for the given json file
func NewConfig(file string) *Config {
	f, err := os.Open(file)
	if err != nil {
		fmt.Println("Error: Unable to find the config file " + file)
		os.Exit(2)
	}

	decoder := json.NewDecoder(f)
	config := &Config{}
	decoder.Decode(&config)

	return config
}
