package main

import (
	"go-tcp-proxy/cmd"
	"log"
	"os"
)

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("missing config file")
	}

	cfg, err := cmd.ParseConfig(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	routes := cmd.BuildRoutes(cfg)
	log.Fatal(routes.Run())
}
