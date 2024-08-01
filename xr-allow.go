package main

import (
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Servers []Server
}

type Server struct {
	Account    string `toml:"account"`
	ServerName string `toml:"server_name"`
	SecretKey  string `toml:"api_secret_key"`
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

func getApiServer(s string) (string, error) {
	if strings.HasSuffix(s, ".xrea.com") {
		return "api.xrea.com", nil
	}

	return "", fmt.Errorf("invalid server_name: %s", s)
}

func getEndpoint(s string) (string, error) {
	domain, err := getApiServer(s)
	if err != nil {
		return "", err
	}
	u, err := url.Parse(fmt.Sprintf("https://%s", domain))
	if err != nil {
		return "", err
	}
	u.Path = "/v1/tool/ssh_ip_allow"

	return u.String(), nil
}

func request(s Server, addr string) error {
	u, err := getEndpoint(s.ServerName)
	if err != nil {
		return err
	}

	v := url.Values{}
	v.Set("account", s.Account)
	v.Set("server_name", s.ServerName)
	v.Set("api_secret_key", s.SecretKey)
	v.Set("param[addr]", addr)

	fmt.Printf("Endpoint: %s\n", u)
	fmt.Printf("%v\n", v)

	return nil
}
