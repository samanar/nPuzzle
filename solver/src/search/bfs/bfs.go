package bfs

import (
	"fmt"
	"github.com/gorilla/websocket"
	"log"
	"nPuzzle"
)

type FIFO struct {
	queue []*nPuzzle.Node
}

type Reply struct {
	Title string `json:"title"`
}

func Solve(puzzle *nPuzzle.NPuzzle , conn *websocket.Conn) {
	visited := make(map[*nPuzzle.Node]bool)
	var fifo FIFO
	fifo.push(puzzle.Root)
	children := puzzle.Root.GenerateChildren(puzzle.RowSize , puzzle.ColSize)
	for _,n := range children {
		if !visited[n] {
			visited[n] = true
			fifo.push(n)
		}
	}
	reply := Reply{Title:"saman"}
	if err := conn.WriteJSON(reply); err != nil {
		log.Println(err)
		return
	}
	for _,n := range visited {
		fmt.Println(n)
	}
}

func (fifo *FIFO) Pop() *nPuzzle.Node{
	if len(fifo.queue) == 0 {
		return nil
	}

	head := fifo.queue[0]
	fifo.queue = fifo.queue[1:]
	return head
}

func (fifo *FIFO) push(node *nPuzzle.Node) {
	fifo.queue  = append(fifo.queue , node)
}

func (fifo *FIFO) Print() {
	for _ , n := range fifo.queue {
		fmt.Println(n)
	}
}

