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

func AppendNewPeer(conn *websocket.Conn, target string) Peer {
	p := GetPeer(conn, target)
	Peers = append(Peers, p)
	return p
}

func GetPeer(conn *websocket.Conn, target string) Peer {
	return Peer{conn, make(chan []byte), target}
}

func (c *Peer) Read() {
	defer func() {
		c.socket.Close()
	}()

	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			c.socket.Close()
			break
		}
		m := msg{}
		err = json.Unmarshal(message, &m)
		if err != nil {
			fmt.Println("Peer Read() err : " + err.Error())
			continue
		}
		if m.event == "new addr" { // 新節點事件
			ConnectionToAddr(m.content, true)
		}
	}
}

func (c *Peer) Write() {
	defer func() {
		c.socket.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.socket.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			c.socket.WriteMessage(websocket.TextMessage, message)
		}
	}
}
