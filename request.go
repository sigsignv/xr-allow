package main

import (
	"fmt"
	"strings"
)

func getApiServer(s string) (string, error) {
	if strings.HasSuffix(s, ".xrea.com") {
		return "api.xrea.com", nil
	}

	return "", fmt.Errorf("invalid server_name: %s", s)
}
