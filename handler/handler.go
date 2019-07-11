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
	fmt.Println(p2p.Port + " server GetPeers : ")
	fmt.Printf("%+v\n", addrs)
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

	conn, err := (&websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}).Upgrade(res, req, nil)
	if err != nil {
		fmt.Println("new client error: " + err.Error())
		http.NotFound(res, req)
		return
	}
	ip := strings.Split(conn.RemoteAddr().String(), ":")
	taget := ip[0] + ":" + rPort[0]
	fmt.Println("new Peer target :" + taget)

	//廣播新結點
	p2p.SendToPeer([]byte(taget))

	// 新節點監聽
	newPeer := p2p.AppendNewPeer(conn, taget)
	go newPeer.Write()
	go newPeer.Read()
}
