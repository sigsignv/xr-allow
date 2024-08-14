package main

import (
	"testing"
)

func TestParseConfig(t *testing.T) {
	toml := `
[[servers]]
account = "USERNAME"
server_name = "SERVERNAME.xrea.com"
api_secret_key = "YOUR_API_KEY"
`

	config, err := parseConfig([]byte(toml))
	if err != nil {
		t.Error(err)
	}

	if len(config.Servers) != 1 {
		t.Errorf("Array length is incorrect: %d\n", len(config.Servers))
	}

	for _, s := range config.Servers {
		if s.Account != "USERNAME" {
			t.Errorf("Account is incorrect: %s\n", s.Account)
		}
		if s.ServerName != "SERVERNAME.xrea.com" {
			t.Errorf("ServerName is incorrect: %s\n", s.ServerName)
		}
		if s.SecretKey != "YOUR_API_KEY" {
			t.Errorf("SecretKey is incorrect: %s\n", s.SecretKey)
		}
	}
}
