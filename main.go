package main

import (
	"fmt"
	"log"
	"net/netip"
	"os"
	"time"

	flag "github.com/spf13/pflag"
)

var (
	VERSION = "0.2.0-dev"
)

func main() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: xr-allow [options...] 192.0.2.0\n")
		flag.PrintDefaults()
	}
	help := flag.BoolP("help", "h", false, "Show help")
	version := flag.BoolP("version", "v", false, "Show version")
	flag.Parse()

	if *help {
		flag.Usage()
		os.Exit(0)
	}

	if *version {
		fmt.Println(VERSION)
		os.Exit(0)
	}

	if flag.NArg() == 0 {
		flag.Usage()
		os.Exit(1)
	}

	addr := flag.Arg(0)
	ip, err := netip.ParseAddr(addr)
	if err != nil {
		log.Fatalf("invalid IP address: %s", addr)
	}
	if !ip.Is4() {
		log.Fatalf("IPv4 address required: %s", ip.String())
	}

	config, err := loadConfig("./conf.toml")
	if err != nil {
		log.Fatal(err)
	}

	for i, s := range config.Servers {
		// 1 req/sec
		if i != 0 {
			time.Sleep(1000 * time.Millisecond)
		}

		u, err := getAPIEndpoint(s.ServerName)
		if err != nil {
			log.Printf("[%s@%s] %s\n", s.Account, s.ServerName, err)
			continue
		}

		if err := request(u, getParams(s, ip)); err != nil {
			log.Printf("[%s@%s] %s\n", s.Account, s.ServerName, err)
			continue
		}

		fmt.Printf("%s@%s: success\n", s.Account, s.ServerName)
	}
}
