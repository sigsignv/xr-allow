package main

import (
	"fmt"
	"net/netip"
	"net/url"
	"strings"
)

func getAPIEndpoint(s string) (string, error) {
	host, err := resolveAPIEndpoint(s)
	if err != nil {
		return "", err
	}

	u := &url.URL{}
	u.Scheme = "https"
	u.Host = host
	u.Path = "/v1/tool/ssh_ip_allow"

	return u.String(), nil
}

func getParams(s Server, ip netip.Addr) url.Values {
	v := url.Values{}
	v.Set("account", s.Account)
	v.Set("server_name", s.ServerName)
	v.Set("api_secret_key", s.SecretKey)
	v.Set("param[addr]", ip.String())

	return v
}

func resolveAPIEndpoint(s string) (string, error) {
	if strings.HasSuffix(s, ".xrea.com") {
		return "api.xrea.com", nil
	}

	return "", fmt.Errorf("invalid server_name: %s", s)
}
