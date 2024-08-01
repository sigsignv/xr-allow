package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pelletier/go-toml/v2"
)

type Config struct {
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

	fmt.Println("Success")
}
