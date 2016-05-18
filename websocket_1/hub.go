package main

type (
	Hub struct {
		// Register adds a connection from the list
		Register chan *Connection
		// Unregister removes a connection from the list
		Unregister chan *Connection
		// Connections the list of registed connectons
		Connections map[*Connection]bool
		// Broadcast messages from the connections
		Broadcast chan []byte
	}
)

func NewHub() *Hub {
	return &Hub{
		Register:    make(chan *Connection),
		Unregister:  make(chan *Connection),
		Connections: make(map[*Connection]bool),
		Broadcast:   make(chan []byte),
	}
}

func (h *Hub) Run() {
	for {
		select {
		case c := <-h.Register:
			h.Connections[c] = true
		case c := <-h.Unregister:
			if _, ok := h.Connections[c]; ok {
				delete(h.Connections, c)
				close(c.send)
			}

		case message := <-h.Broadcast:
			for c := range h.Connections {
				select {
				case c.send <- message:
				default:
					close(c.send)
					delete(h.Connections, c)
				}
			}
		}
	}
}
