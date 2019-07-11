package main

import (
	"flag"
	"log"

	"github.com/jou66jou/go-p2p-websocket/p2p"
	"github.com/jou66jou/go-p2p-websocket/router"
)

func main() {
	var port string
	flag.StringVar(&port, "p", "", "listen port")
	flag.Parse()
	p2p.Port = port
	if port != "8080" { //8080為種子
		go p2p.RunP2P()
	}
	log.Fatal(router.RunHTTP(port))
}
