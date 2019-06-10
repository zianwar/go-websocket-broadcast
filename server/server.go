package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type counter struct {
	v int
}

func getWsUpgrader() *websocket.Upgrader {
	return &websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
		CheckOrigin:     func(r *http.Request) bool { return true },
	}
}

func main() {
	gin.DefaultWriter = ioutil.Discard
	r := gin.Default()

	counterCh := make(chan counter)
	globalQuit := make(chan struct{})
	hub := NewHub(counterCh, globalQuit)

	defer close(globalQuit)

	go hub.start()
	go updateCounterEvery(time.Second, counterCh)

	r.GET("/status", wsHandler(getWsUpgrader(), globalQuit, hub))
	r.Run("localhost:7777")
}

func wsHandler(wsupgrader *websocket.Upgrader, globalQuit chan struct{}, hub *hub) gin.HandlerFunc {
	return func(c *gin.Context) {
		conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Fatalf("failed to upgrade websocket: %v\n", err)
		}

		client := NewClient(conn)
		go client.handle()

		hub.Register(client)
	}
}

func updateCounterEvery(d time.Duration, counterCh chan counter) {
	c := counter{}
	ticker := time.NewTicker(d)
	for {
		select {
		case <-ticker.C:
			c.v++
			counterCh <- c
		}
	}
}
