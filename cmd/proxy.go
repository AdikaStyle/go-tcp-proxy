package cmd

import (
	"fmt"
	"inet.af/tcpproxy"
	"log"
)

func BuildRoutes(cfg Config) *tcpproxy.Proxy {
	var p tcpproxy.Proxy

	for _, route := range cfg.Routes {
		log.Printf("Loaded Rule: %s (%d -> %s)", route.Name, route.FromPort, route.ToAddress)
		p.AddRoute(fmt.Sprintf(":%d", route.FromPort), tcpproxy.To(route.ToAddress))
	}

	return &p
}
