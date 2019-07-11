package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"

	"github.com/gorilla/websocket"
	"github.com/jou66jou/go-p2p-websocket/p2p"
)

func GetPeers(res http.ResponseWriter, req *http.Request) {

	var addrs []string
	for _, p := range p2p.Peers {
		addrs = append(addrs, p.Taget)
	}
	b, e := json.Marshal(addrs)
	if e != nil {
		fmt.Println(e)
	}
	res.Write(b)
}

func NewWS(res http.ResponseWriter, req *http.Request) {
	q := req.URL.Query()
	rPort, ok := q["port"]
	if !ok {
		fmt.Println("url value port is nil")
		http.NotFound(res, req)
		return
	}
	ip := strings.Split(req.RemoteAddr, ":")
	taget := ip[0] + ":" + rPort[0]
	// fmt.Println("new Peer target :" + taget)

	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		fmt.Println("new client error: " + err.Error())
		http.NotFound(res, req)
		return
	}

	// req帶有brdcst key則不進行廣播，brdcst代表req端是接收到廣播而發起websocket，避免廣播風暴
	v, ok := q["brdcst"]
	if !ok {
		if len(v) == 0 {
			// 廣播新結點
			p2p.BroadcastAddr(taget)
		}
	}

	// p2p
	newPeer := p2p.AppendNewPeer(conn, taget)
	go newPeer.Write()
	go newPeer.Read()
}
