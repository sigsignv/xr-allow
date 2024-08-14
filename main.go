package main

import (
	"flag"
	"fmt"
	"log"
	"net/netip"
	"os"
	"time"
)

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

	for i, s := range config.Servers {
		// 1 req/sec
		if i != 0 {
			time.Sleep(1000 * time.Millisecond)
		}

		u, err := getAPIEndpoint(s.ServerName)
		if err != nil {
			log.Println(err)
			continue
		}
		v := getParams(s, addr)
		err = request(u, v)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s@%s: success\n", s.Account, s.ServerName)
	}
}
