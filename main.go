package main

import (
	"flag"
	"log"

	"github.com/jou66jou/go-p2p-websocket/p2p"
	"github.com/jou66jou/go-p2p-websocket/router"
)

var (
	port string
	seed string
)

func main() {
	initFlag()
	p2p.MyPort = port
	if ("127.0.0.1:" + port) != seed {
		go p2p.ConnectionToAddr(seed, false)
	}
	log.Fatal(router.RunHTTP(port))
}

func initFlag() {
	flag.StringVar(&port, "p", "8080", "listen port")               // 8080
	flag.StringVar(&seed, "seed", "127.0.0.1:8080", "seed ip:port") // 127.0.0.1:8080
	flag.Parse()
}
