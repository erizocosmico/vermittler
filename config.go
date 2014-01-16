package main

// App configuration.
type Config struct {
	// Enable or disable the cache
	CacheEnabled bool
	// Path where the cache is stored, without trailing slash
	CachePath string
	// Port where the app will run preceded e.g: "8888"
	Port string
}

// Returns a fancy and very hardcoded config for our app.
func NewConfig() *Config {
	return &Config{
		CacheEnabled: true,
		CachePath:    "/path/to/cache/dir",
		Port:         "8888",
	}
}
