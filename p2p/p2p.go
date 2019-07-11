package p2p

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

type msg struct {
	event   string
	content string
}

func RunP2P() {

	res, err := http.Get("http://127.0.0.1:8080/peers") // 從種子獲得列表
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	addrs := []interface{}{}
	json.Unmarshal(b, &addrs)

	//加入種子8080
	addrs = append(addrs, "127.0.0.1:8080")
	fmt.Printf("%+v\n", addrs)
	for _, v := range addrs {
		u := url.URL{Scheme: "ws", Host: v.(string), Path: "/new", RawQuery: "port=" + Port}
		var dialer *websocket.Dialer

		conn, _, err := dialer.Dial(u.String(), nil)
		if err != nil {
			panic("p2p err: " + err.Error())
		}
		fmt.Println("new addr :" + conn.RemoteAddr().String())
		go func(conn *websocket.Conn) {
			conn.SetReadDeadline(time.Now().Add(100 * time.Minute))
			for {
				_, message, err := conn.ReadMessage()
				if err != nil {
					fmt.Println("read err :", err)
					return
				}
				fmt.Printf("received: %s\n", string(message))
			}
		}(conn)
	}

}

func BoardAddr(msg []byte) {
	for _, p := range Peers {
		p.send <- msg
	}
}
