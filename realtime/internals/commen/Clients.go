package commen

import "github.com/gorilla/websocket"

type Client struct {
	ID    int32
	Inbox chan []byte
	conn  *websocket.Conn
}

func NewClient(id int32, conn *websocket.Conn) *Client {
	return &Client{
		ID:    id,
		Inbox: make(chan []byte),
		conn:  conn,
	}
}

func (c *Client) WritePump() {
	defer func() {
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.Inbox:
			if !ok {
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}
			c.conn.WriteMessage(websocket.TextMessage, message)
		}
	}
}
