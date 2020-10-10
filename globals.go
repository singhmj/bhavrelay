package main

// NOTE: don't use these variables directly, please use their getter/setter methods
var (
	cache  *Cache
	config *SystemConfig
)

// SetCache : Sets the global cache variable for you
func SetCache(c *Cache) {
	// NOTE: add locks here if you think that it can be accessed in multiple routines
	cache = c
}

// GetCache : returns the global cache variable for you
func GetCache() *Cache {
	// NOTE: add locks here if you think that it can be accessed in multiple routines
	return cache
}

// SetConfig : Sets the global system config variable for you
func SetConfig(c *SystemConfig) {
	// NOTE: add locks here if you think that it can be accessed in multiple routines
	config = c
}

// GetConfig : Gets the global system config variable for you
func GetConfig() *SystemConfig {
	// NOTE: add locks here if you think that it can be accessed in multiple routines
	return config
}
