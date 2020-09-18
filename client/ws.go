package main

import (
	"golang.org/x/net/websocket"
)

var (
	conn *websocket.Conn
)

func wsHandler(c *websocket.Conn) {
	conn = c
	for {
		var m message
		err := websocket.JSON.Receive(conn, &m)
		if err != nil {
			continue
		}
		_, err = send(m)
		if err != nil {
			continue
		}
	}
}