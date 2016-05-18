package main

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/websocket"
)

type Ws struct {
	upgrader    websocket.Upgrader
	requestPath string
	hub         *Hub
}

func New(requestPath string) *Ws {
	w := &Ws{
		requestPath: requestPath,
		hub:         NewHub(),
	}
	w.upgrader = websocket.New(w.HandleConnection)
	return w
}

func (w *Ws) PostListen(station *iris.Iris) {
	station.Get(w.requestPath, func(ctx *iris.Context) { // or.Any
		if err := w.upgrader.Upgrade(ctx); err != nil {
			station.Logger().Printf("Error on websocket connection upgrade, reason: %s", err.Error())
			return
		}
	})

	w.hub.Run() // start the hub
}

func (w *Ws) HandleConnection(websocketConn *websocket.Conn) {
	conn := NewConnection(websocketConn, w.connectionClosed, w.messageReceived)

	w.hub.Register <- conn
	conn.Listen()
}

func (w *Ws) connectionClosed(conn *Connection) {
	w.hub.Unregister <- conn
}

func (w *Ws) messageReceived(message []byte) {
	w.hub.Broadcast <- message
}
