package main

import (
	"errors"
	"os"
	"strings"
)

// isRequired
func isRequired(list map[string]string) error {
	var err error = nil

	for k, v := range list {
		if len(v) == 0 {
			err = errors.New(strings.Title(k) + " required.")
			return err
		}
	}
	return err
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if len(value) == 0 {
		return defaultValue
	}
	return value
}
