package main

import (
	"time"

	"github.com/kataras/iris/websocket"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

type Connection struct {
	wsConn      *websocket.Conn
	emitClose   func(*Connection)
	emitMessage func([]byte)

	send chan []byte
}

// NewConnection creates a connection and returns it
func NewConnection(websocketConnection *websocket.Conn, onClose func(*Connection), onMessage func([]byte)) *Connection {
	c := &Connection{
		wsConn:      websocketConnection,
		send:        make(chan []byte, 256),
		emitClose:   onClose,
		emitMessage: onMessage,
	}
	return c
}

func (c *Connection) Listen() {
	go c.startWriter()
	c.startReader()
}

func (c *Connection) write(messageType int, payload []byte) error {
	c.wsConn.SetWriteDeadline(time.Now().Add(writeWait))
	return c.wsConn.WriteMessage(messageType, payload)
}

// startWriter sends messages from the hub to the websocket client/connection
func (c *Connection) startWriter() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.wsConn.Close()
	}()

	for {
		select {
		case message, ok := <-c.send:
			if !ok {
				c.write(websocket.CloseMessage, []byte{})
				return
			}

			if err := c.write(websocket.TextMessage, message); err != nil {
				return
			}
		case <-ticker.C:
			if err := c.write(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func (c *Connection) startReader() {
	defer func() {
		// if this breaks means that an error happen
		c.emitClose(c) // send notification to the hub that this hub is closed by an error or timeout
		c.wsConn.Close()
	}()

	c.wsConn.SetReadLimit(maxMessageSize)
	c.wsConn.SetReadDeadline(time.Now().Add(pongWait))
	c.wsConn.SetPongHandler(func(string) error {
		c.wsConn.SetReadDeadline(time.Now().Add(pongWait)) // on pong just continue the connection by extend its life
		return nil
	})

	for {
		_, message, err := c.wsConn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway) {
				println(err.Error())
			}
			break
		}
		c.emitMessage(message) //send the message to the hub
	}
}
