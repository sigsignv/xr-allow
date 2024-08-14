package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/netip"
	"net/url"
	"strings"
)

type Result struct {
	StatusCode int `json:"status_code"`
}

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

func parseResult(b []byte) (*Result, error) {
	var result Result

	err := json.Unmarshal(b, &result)
	if err != nil {
		return nil, err
	}

	return &result, nil
}

func request(url string, data url.Values) error {
	resp, err := http.PostForm(url, data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// XREA API always returns 200
	if resp.StatusCode != 200 {
		return fmt.Errorf("server is probably unavailable: %s", resp.Status)
	}

	b, err := io.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	r, err := parseResult(b)
	if err != nil {
		return err
	}

	if r.StatusCode != 200 {
		return fmt.Errorf("API request failed: %s", string(b))
	}

	return nil
}

func resolveAPIEndpoint(s string) (string, error) {
	if strings.HasSuffix(s, ".xrea.com") {
		return "api.xrea.com", nil
	}

	return "", fmt.Errorf("invalid server_name: %s", s)
}
