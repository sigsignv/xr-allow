package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Servers []Server
}

type Server struct {
	Account    string `toml:"account" json:"account"`
	ServerName string `toml:"server_name" json:"server_name"`
	SecretKey  string `toml:"api_secret_key" json:"api_secret_key"`
}

func main() {
	doc, err := os.ReadFile("./conf.toml")
	if err != nil {
		log.Fatal("Failed: os.ReadFile()")
	}

	var config Config
	err = toml.Unmarshal(doc, &config)
	if err != nil {
		log.Fatal("Failed: toml.Unmarshal()")
	}

	for _, s := range config.Servers {
		err := request(s, "192.168.1.1")
		if err != nil {
			log.Fatal(err)
		}
	}
}

func getEndpoint(serverName string) (string, error) {
	if strings.HasSuffix(serverName, ".xrea.com") {
		return "api.xrea.com", nil
	}

	return "", fmt.Errorf("invalid server_name: %s", serverName)
}

func request(s Server, addr string) error {
	endpoint, err := getEndpoint(s.ServerName)
	if err != nil {
		return err
	}

	json, err := json.Marshal(s)
	if err != nil {
		return err
	}

	fmt.Printf("Endpoint: %s\n", endpoint)
	fmt.Println(string(json))

	return nil
}
