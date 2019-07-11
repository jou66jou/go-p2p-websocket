package p2p

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/jou66jou/go-p2p-websocket/common"
)

type msg struct {
	event   string
	content string
}

var (
	Peers  []Peer
	MyPort string
)

func ConnectionToAddr(addr string, isBrdcst bool) {
	rawQ := "port=" + MyPort
	if isBrdcst { //是否為接收到廣播而發起連線
		rawQ += ";brdcst=1"
	}
	u := url.URL{Scheme: "ws", Host: addr, Path: common.RouteName["newWS"], RawQuery: rawQ}
	// u := url.URL{Scheme: "ws", Host: addr, Path: "/new", RawQuery: rawQ}
	var dialer *websocket.Dialer

	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("ConnectionToAddr err: " + err.Error())
		return
	}

	// 加入新節點
	newPeer := AppendNewPeer(conn, addr)
	go newPeer.Write()
	go newPeer.Read()

}

func BroadcastAddr(tgt string) {
	m := msg{"new addr", tgt}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("BroadcastAddr json error : " + err.Error())
		return
	}
	for _, p := range Peers {
		p.send <- b
	}
}
