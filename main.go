package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/netip"
	"os"
	"time"
)

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

	config, err := loadConfig("./conf.toml")
	if err != nil {
		log.Fatal(err)
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
	u, err := getAPIEndpoint(s.ServerName)
	if err != nil {
		return err
	}

	v := getParams(s, addr)

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
