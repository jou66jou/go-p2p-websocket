package p2p

import (
	"encoding/json"
	"fmt"

	"github.com/gorilla/websocket"
)

type Peer struct {
	socket *websocket.Conn
	send   chan []byte
	Taget  string
}

func GetPeer(conn *websocket.Conn, target string) Peer {
	return Peer{conn, make(chan []byte), target}
}

func (p *Peer) Read() {
	defer func() {
		p.socket.Close()
	}()

	for {
		_, message, err := p.socket.ReadMessage()
		if err != nil {
			p.socket.Close()
			break
		}
		m := msg{}
		err = json.Unmarshal(message, &m)
		if err != nil {
			fmt.Println("Peer Read() err : " + err.Error())
			continue
		}
		if m.Event == "new addr" { // 新節點事件
			ConnectionToAddr(m.Content, true)
		}
	}
}

func (p *Peer) Write() {
	defer func() {
		p.socket.Close()
	}()

	for {
		select {
		case message, ok := <-p.send:
			if !ok {
				p.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			p.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
