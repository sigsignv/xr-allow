package main

import "testing"

func TestResolveAPIEndpoint(t *testing.T) {
	host, err := resolveAPIEndpoint("SERVERNAME.xrea.com")
	if err != nil {
		t.Error(err)
	}
	if host != "api.xrea.com" {
		t.Errorf("invalid API Endpoint server: %s\n", host)
	}
}

func TestResolveAPIEndpointFailure(t *testing.T) {
	_, err := resolveAPIEndpoint("invalid.example.com")
	if err == nil {
		t.Error("should return an error")
	}
}

func TestGetAPIEndpoint(t *testing.T) {
	u, err := getAPIEndpoint("SERVERNAME.xrea.com")
	if err != nil {
		t.Error(err)
	}
	if u != "https://api.xrea.com/v1/tool/ssh_ip_allow" {
		t.Errorf("invalid API Endpoint: %s\n", u)
	}
}

func TestParseResult(t *testing.T) {
	json := "{\"status_code\":200}"
	result, err := parseResult([]byte(json))
	if err != nil {
		t.Error(err)
	}
	if result.StatusCode != 200 {
		t.Errorf("invalid status code: %d\n", result.StatusCode)
	}
}
