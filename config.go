package main

import "github.com/pelletier/go-toml/v2"

type Config struct {
	Servers []Server
}

type Server struct {
	Account    string `toml:"account"`
	ServerName string `toml:"server_name"`
	SecretKey  string `toml:"api_secret_key"`
}

func parseConfig(b []byte) (*Config, error) {
	var config Config

	err := toml.Unmarshal(b, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}
