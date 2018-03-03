package util

// +build test

import (
	"flag"
	"log"
)

// These are functions for shortening testing boilerplate.

// ReadTestConfig parses command line and reads configuration.
func ReadTestConfig(conf interface{}) {
	fconfig := flag.String(
		"config", "dappctrl-test.config.json", "Configuration file")
	flag.Parse()

	if err := ReadJSONFile(*fconfig, conf); err != nil {
		log.Fatalf("failed to read configuration: %s\n", err)
	}
}

// NewTestLogger creates a new logger.
func NewTestLogger(conf *LogConfig) *Logger {
	logger, err := NewLogger(conf)
	if err != nil {
		log.Fatalf("failed to create logger: %s\n", err)
	}
	return logger
}
