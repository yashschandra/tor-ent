package main

import (
	"flag"
	"fmt"
	"net/http"
	"tor-ent/common"
	"golang.org/x/net/websocket"
)

const (
	receiveURL = "/receive"
	port = ":17667"
)

var (
	serverAddress *string
	name *string
)

func register(name string, address string) (error) {
	e := common.Entry{Name: name, Address: address}
	url := fmt.Sprintf("%s/add", *serverAddress)
	_, err := sendRequest(http.MethodGet, url, e)
	if err != nil {
		return err
	}
	return nil
}

func start() {
	serverAddress = flag.String("server", "http://localhost:17666", "")
	name = flag.String("name", "", "")
	address := flag.String("address", "localhost:17667", "")
	flag.Parse()
	err := register(*name, *address)
	if err != nil {
		panic(err)
	}
	http.Handle("/ws", websocket.Handler(wsHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.HandleFunc(receiveURL, receive)
	if err := http.ListenAndServe(port, nil); err != nil {
		panic(err)
	}
}

func main() {
	start()
}