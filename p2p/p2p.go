package p2p

import (
	"encoding/json"
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/jou66jou/go-p2p-websocket/common"
)

// 節點通信格式
type msg struct {
	Event   string `json:"event"`   // 事件
	Content string `json:"content"` // 內容
}

var (
	Peers  []Peer
	MyPort string //本機port
)

// 向指定位置發出websocket請求
func ConnectionToAddr(addr string, isBrdcst bool) {

	rawQ := "port=" + MyPort // 將本機port發送給目標節點

	if isBrdcst { // 是否為接收到廣播而發起連線，可避免廣播風暴
		rawQ += ";brdcst=1"
	}

	// 建立websocket連線
	u := url.URL{Scheme: "ws", Host: addr, Path: common.RouteName["newWS"], RawQuery: rawQ}
	var dialer *websocket.Dialer

	conn, _, err := dialer.Dial(u.String(), nil)
	if err != nil {
		fmt.Println("ConnectionToAddr err: " + err.Error())
		return
	}

	// 向本機加入新節點
	newPeer := AppendNewPeer(conn, addr)
	go newPeer.Write()
	go newPeer.Read()

}

// 發出新節點事件廣播
func BroadcastAddr(tgt string) {
	m := &msg{"new addr", tgt}
	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("BroadcastAddr json error : " + err.Error())
		return
	}

	// 遍歷本機節點將訊息傳入channel
	for _, p := range Peers {
		p.send <- b
	}
}

// 加入新節點
func AppendNewPeer(conn *websocket.Conn, target string) Peer {
	p := NewPeer(conn, target)
	Peers = append(Peers, p)
	return p
}
