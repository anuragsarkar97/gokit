package env

import (
	"os"
	"strconv"
	"strings"
	"time"
)

// Helper function to read an environment or return a default value
func AsString(key string, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// Helper function to read an environment variable into integer or return a default value
// Return Default value in case of an error
func AsInt(name string, defaultVal int) int {
	valueStr := AsString(name, "")
	if value, err := strconv.Atoi(valueStr); err == nil {
		return value
	}
	return defaultVal
}

// Helper to read an environment variable into a bool or return default value
// Return Default value in case of an error
func AsBool(name string, defaultVal bool) bool {
	valStr := AsString(name, "")
	if val, err := strconv.ParseBool(valStr); err == nil {
		return val
	}
	return defaultVal
}

// Helper to read an environment variable into time.Duration or return default value
// Return Default value in case of an error
func AsMillisecondDuration(name string, defaultVal time.Duration) time.Duration {
	valInt := AsInt(name, -1)
	if valInt == -1 {
		return defaultVal
	}
	return time.Millisecond * time.Duration(valInt)
}

// Helper to read an environment variable into time.Duration or return default value
// Return Default value in case of an error
func AsSecondDuration(name string, defaultVal time.Duration) time.Duration {
	valInt := AsInt(name, -1)
	if valInt == -1 {
		return defaultVal
	}
	return time.Second * time.Duration(valInt)
}

// Helper to read an environment variable into a string array or return default value
// Return Default value in case of an error
func AsStringArray(key string, separator string, defaultVal []string) []string {
	valueStr := AsString(key, "")
	if valueStr == "" {
		return defaultVal
	}
	valueStr = strings.ReplaceAll(valueStr, " ", "")
	return strings.Split(valueStr, separator)
}
