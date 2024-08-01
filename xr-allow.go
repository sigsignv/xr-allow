package main

import (
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
	Account   string `toml:"account"`
	Server    string `toml:"server_name"`
	SecretKey string `toml:"api_secret_key"`
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

	fmt.Printf("%v\n", config)
}

func getEndpoint(server string) (string, error) {
	if strings.HasSuffix(server, ".xrea.com") {
		return "api.xrea.com", nil
	}

	return "", fmt.Errorf("invalid server_name: %s", server)
}
