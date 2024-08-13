package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/netip"
	"net/url"
	"os"
	"strings"
	"time"

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

type Result struct {
	StatusCode int    `json:"status_code"`
	Message    string `json:"message"`
}

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: xr-allow [options...] 192.0.2.0\n")
		flag.PrintDefaults()
	}
	flag.Parse()

	addr, err := netip.ParseAddr(flag.Arg(0))
	if err != nil {
		log.Fatal(err)
	}

	if !addr.Is4() {
		log.Fatalf("IPv4 address required: %s", addr.String())
	}

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
		err := request(s, addr.String())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s@%s: success\n", s.Account, s.ServerName)
		time.Sleep(1000 * time.Millisecond)
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
	u := &url.URL{}
	u.Scheme = "https"
	u.Host = domain
	u.Path = "/v1/tool/ssh_ip_allow"

	return u.String(), nil
}

func parseResult(r io.Reader) (*Result, error) {
	b, err := io.ReadAll(r)
	if err != nil {
		return nil, err
	}

	var resp Result
	err = json.Unmarshal(b, &resp)
	if err != nil {
		return nil, err
	}

	return &resp, nil
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

	resp, err := http.PostForm(u, v)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return fmt.Errorf("request is failed: %s", resp.Status)
	}

	result, err := parseResult(resp.Body)
	if err != nil {
		return err
	}

	if result.StatusCode != 200 {
		return fmt.Errorf("result is failed: %s", result.Message)
	}

	return nil
}
