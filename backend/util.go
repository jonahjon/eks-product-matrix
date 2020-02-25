package main

import (
	"errors"
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
