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
	if port != "8080" { //8080為種子
		go p2p.ConnectionToAddr(seed)
	}
	log.Fatal(router.RunHTTP(port))
}

func initFlag() {
	flag.StringVar(&port, "p", "8081", "listen port")               // 8081
	flag.StringVar(&seed, "seed", "127.0.0.1:8080", "seed ip:port") // 127.0.0.1:8080
	flag.Parse()
}
