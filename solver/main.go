package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	
	"nPuzzle"
	
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type Message struct {
	Title   string `json:"title"`
	Type    string `json:"type"`
	Body    string `json:"body"`
	Numbers []int  `json:"numbers"`
	RowSize int    `json:"rowSize"`
	ColSize int    `json:"colSize"`
}

func homePage(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "home page")
	if err != nil {
		log.Fatal(err)
	}
}

func wsEndpoint(w http.ResponseWriter, r *http.Request) {
	upgrader.CheckOrigin = func(r *http.Request) bool { return true }
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("client connected successfully")
	socketHandler(conn)
	
}

func socketHandler(conn *websocket.Conn) {
	for {
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		
		var message Message
		err = json.Unmarshal([]byte(string(p)), &message)
		if err != nil {
			log.Fatal("error when parsing json", err)
			return
		}
		
		if len(message.Type) == 0 {
			log.Fatal("type needs to be specified")
			return
		}
		
		switch message.Type {
		case "nPuzzle":
			NPuzzleHandler(conn, &message)
		}
		
		if err := conn.WriteMessage(messageType, p); err != nil {
			log.Println(err)
			return
		}
	}
}

func NPuzzleHandler(conn *websocket.Conn, message *Message) {
	switch message.Title {
	case "init":
		puzzle := nPuzzle.NPuzzle{
			RowSize: message.RowSize,
			ColSize: message.ColSize,
			Root:    &nPuzzle.Node{Numbers: message.Numbers},
		}
		// fmt.Println(puzzle)
		// fmt.Println(*puzzle.Root)
		children := (*puzzle.Root).GenerateChildren(puzzle.RowSize, puzzle.ColSize)
		for _ , n := range children{
			fmt.Println(*n)
		}
	}
}

func setUpRoutes() {
	http.HandleFunc("/", homePage);
	http.HandleFunc("/ws", wsEndpoint);
}

func main() {
	fmt.Println("server started on localhost:8000")
	
	setUpRoutes()
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
