package main

import (
	"time"

	"github.com/gorilla/websocket"
)

const (
	// Time allowed to write message to peer
	writeWait = 10 * time.Second

	// Time allowed to write message to peer
	pongWait = 60 * time.Second

	// Send pings to peer, must be less than pongWait
	pingPeriod = (pongWait * 9) / 10

	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Client struct {
	room *Room
	conn *websocket.Conn
	// Buffered channel of outbound messages
	send chan []byte
}

// readPump pumps messages from the websocket connection to the room
func(c *Client) readPump() {
    defer func() {
        c.room.unregister <- c
        c.conn.Close()
    }()
}
