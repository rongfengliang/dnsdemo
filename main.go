package main

import (
	"net"
	"os"

	"github.com/docker/dnsserver"
)

func main() {
	dnsserver := dnsserver.NewDNSServer("docker")
	dnsserver.SetA("dalong", net.ParseIP("127.0.0.2"))
	if len(os.Args) > 1 {
		if err := dnsserver.Listen(os.Args[1]); err != nil {
			panic(err)
		}
	} else {
		if err := dnsserver.Listen("0.0.0.0:53"); err != nil {
			panic(err)
		}
	}
}
