package main

import "testing"

func TestResolveAPIEndpoint(t *testing.T) {
	host, err := resolveAPIEndpoint("SERVERNAME.xrea.com")
	if err != nil {
		t.Error(err)
	}
	if host != "api.xrea.com" {
		t.Errorf("invalid API Endpoint: %s\n", host)
	}
}

func TestResolveAPIEndpointFailure(t *testing.T) {
	_, err := resolveAPIEndpoint("invalid.example.com")
	if err == nil {
		t.Error("should return an error")
	}
}
