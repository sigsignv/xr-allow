package main

import "testing"

func TestResolveAPIEndpoint(t *testing.T) {
	domain, err := resolveAPIEndpoint("SERVERNAME.xrea.com")
	if err != nil {
		t.Error(err)
	}
	if domain != "api.xrea.com" {
		t.Errorf("invalid API Endpoint: %s\n", domain)
	}
}

func TestResolveAPIEndpointFailure(t *testing.T) {
	_, err := resolveAPIEndpoint("invalid.example.com")
	if err == nil {
		t.Error("should return an error")
	}
}
