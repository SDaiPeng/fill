package main

import (
	"fill/server/wsConn"
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
	"time"
)

var (
	upGrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHnadler(w http.ResponseWriter, r *http.Request) {
	var (
		data         []byte
		err          error
		wsConnection *websocket.Conn
		connection   *wsConn.Connection
	)
	if wsConnection, err = upGrader.Upgrade(w, r, nil); err != nil {
		return
	}
	if connection, err = wsConn.ConnectionInit(wsConnection); err != nil {
		goto Err
	}
	go func() {
		for {
			if err =  connection.WriteMessage([]byte("hello world")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}
	}()
	for {
		if data, err = connection.ReadMessage(); err != nil {
			return
		}
		fmt.Println(string(data))
	}
Err:
	connection.Close()
}

func main() {
	http.HandleFunc("/ws", wsHnadler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
