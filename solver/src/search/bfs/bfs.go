package bfs

import (
	"fmt"
	"strings"
	
	"github.com/gorilla/websocket"
	"nPuzzle"
	nPuzzle2 "structure/nPuzzle"
)

type Reply struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

func Solve(puzzle *nPuzzle.NPuzzle, conn *websocket.Conn) {
	visited := make(map[*nPuzzle.Node]bool)
	output := make(chan Reply)
	goalFound := false
	go sendLogs(output, conn)
	fifo := nPuzzle2.FIFO{}
	count := 0
	fifo.Push(puzzle.Root)
	for !fifo.IsEmpty() {
		child := fifo.Pop()
		if child.IsGoal() {
			goalFound = true
			// trace path to route
			break
		}
		for _, n := range child.GenerateChildren(puzzle.RowSize, puzzle.ColSize) {
			if !visited[n] {
				visited[n] = true
				output <- Reply{Title: "log", Body: convertArrayToString(n.Numbers)}
				count++
				fmt.Println(count)
				fifo.Push(n)
			}
		}
	}
	if goalFound {
		output <- Reply{Title: "Found", Body: "goal was found"}
		
	} else {
	
	}
	
	close(output)
}

func sendLogs(output chan Reply, conn *websocket.Conn) {
	for {
		select {
		case val, ok := <-output:
			if ok {
				reply := Reply{Title: val}
				err := conn.WriteJSON(reply)
				if err != nil {
					fmt.Println("something wen wrong trying to send data to socket ", err)
				}
			} else {
				return
			}
		}
	}
}

func convertArrayToString(array []int) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), "-"), "[]")
}
