package main

import (
	"net/http"
	"encoding/json"
	"io/ioutil"
	"fmt"
	"golang.org/x/net/websocket"
)

type message struct {
	To string `json:"to"`
	Name string `json:"name"`
	Content string `json:"content"`
}

func newMessage(to, content string) message {
	return message{Name: *name, Content: content, To: to}
}

func getReceiveURL(onion string) string {
	return fmt.Sprint("http://", onion, receiveURL)
}

func receive(w http.ResponseWriter, r *http.Request) {
	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	var m message
	err = json.Unmarshal(b, &m)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}
	websocket.JSON.Send(conn, m)
}

func send(m message) ([]byte, error) {
	onion, err := getOnion(m.To)
	if err != nil {
		return nil, err
	}
	resp, err := sendRequest(http.MethodGet, getReceiveURL(onion), newMessage(m.To, m.Content))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return ioutil.ReadAll(resp.Body)
}