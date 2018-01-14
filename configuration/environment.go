package configuration

import (
	"os"
	"strconv"
)

// SetEnvironmentVariables changes config settings when certain environment variables are found
func SetEnvironmentVariables(conf *Config) {
	interval := os.Getenv("MOONFOLIO_INTERVAL_SECONDS")
	if interval != "" {
		iValue, err := strconv.Atoi(interval)
		if err != nil {
			conf.IntervalSeconds = iValue
		}
	}
}
