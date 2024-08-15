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

		if err := request(u, getParams(s, addr)); err != nil {
			log.Printf("[%s@%s] %s\n", s.Account, s.ServerName, err)
			continue
		}

		fmt.Printf("%s@%s: success\n", s.Account, s.ServerName)
	}
}
