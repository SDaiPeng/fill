package wsConn

import (
	"github.com/gorilla/websocket"
	"sync"
)

type Connection struct {
	wsConnection *websocket.Conn
	inChan       chan []byte
	outChan      chan []byte
	closeChan    chan byte
	mutex        sync.Mutex
	isClosed     bool
}

func ConnectionInit(wsConnection *websocket.Conn) (conn *Connection, err error) {
	conn = &Connection{
		wsConnection: wsConnection,
		inChan:       make(chan []byte, 1000),
		outChan:      make(chan []byte, 1000),
		closeChan:    make(chan byte, 1),
	}
	go conn.readLoop()

	go conn.writeLoop()
	return
}

func (conn *Connection) ReadMessage() (data []byte, err error) {
	select {
	case data = <-conn.inChan:
	case <-conn.closeChan:
	}
	return
}
func (conn *Connection) WriteMessage(data []byte) (err error) {
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
	}
	return
}

func (conn *Connection) Close() (err error) {
	conn.wsConnection.Close()
	conn.mutex.Lock()
	if !conn.isClosed {
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
	return
}

func (conn *Connection) readLoop() {
	var (
		data []byte
		err  error
	)
	for {
		if _, data, err = conn.wsConnection.ReadMessage(); err != nil {
			goto ERR
		}
		select {
		case conn.inChan <- data:
		case <-conn.closeChan:
			goto ERR
		}
	}
ERR:
	conn.Close()
}

func (conn *Connection) writeLoop() {
	var (
		data []byte
		err  error
	)
	for {
		select {
		case data = <-conn.outChan:
		case <-conn.closeChan:
			goto ERR
		}
		if err = conn.wsConnection.WriteMessage(websocket.TextMessage, data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
}
