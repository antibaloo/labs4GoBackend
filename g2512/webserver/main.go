package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

func main() {
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/ws", WebSocketHandler)

	http.ListenAndServe(":8888", nil)

}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	indexHTML, _ := os.Open("index.html")
	defer indexHTML.Close()
	indexData, _ := io.ReadAll(indexHTML)
	fmt.Fprint(w, string(indexData))
}

var upgrader = websocket.Upgrader{}

func WebSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, _ := upgrader.Upgrade(w, r, nil)
	defer conn.Close()
	//получить сообщение от фронтенда
	mType, bts, _ := conn.ReadMessage()
	//вывести сообщение в консоль
	fmt.Println(string(bts))
	//отпрвить сообщение фронтенду
	conn.WriteMessage(mType, []byte("сообощение с бэкенда"))
}
