package main

import (
	"fmt"
	"net/url"
	"strings"
)

func getEndpoint(s string) (string, error) {
	domain, err := resolveAPIEndpoint(s)
	if err != nil {
		return "", err
	}

	u := &url.URL{}
	u.Scheme = "https"
	u.Host = domain
	u.Path = "/v1/tool/ssh_ip_allow"

	return u.String(), nil
}

func resolveAPIEndpoint(s string) (string, error) {
	if strings.HasSuffix(s, ".xrea.com") {
		return "api.xrea.com", nil
	}

	return "", fmt.Errorf("invalid server_name: %s", s)
}
