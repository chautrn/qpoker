package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type client struct {
	ws *websocket.Conn
}

func addClient(w http.ResponseWriter, r *http.Request, id string) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }

	ws, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}

	var c = client{ws: ws}
	fmt.Printf("Connection received, client %d created\n", id)
	fmt.Println(c)
}
