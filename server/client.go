package main

import (
	"log"

	"github.com/gorilla/websocket"
)

var NextID int

type client struct {
	id   int
	conn *websocket.Conn
	ch   chan counter
	quit chan struct{}
}

func NewClient(conn *websocket.Conn) *client {
	NextID++
	return &client{
		id:   NextID,
		conn: conn,
		quit: make(chan struct{}),
		ch:   make(chan counter),
	}
}

func (c *client) Close() {
	close(c.quit)
}

func (c *client) handle() {
	for {
		select {
		case <-c.quit:
			if err := c.conn.Close(); err != nil {
				log.Printf("client %d connection close error %v\n", c.id, err)
			}
			return
		case n := <-c.ch:
			jmsg := map[string]interface{}{
				"counter":  n.v,
				"clientId": c.id,
			}
			if err := c.conn.WriteJSON(jmsg); err != nil {
				log.Println("ws write error:", err)
				return
			}
		}
	}
}
