package bfs

import (
	"bytes"
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
	var visited [][]byte
	logOutput := make(chan Reply)
	go sendLogs(logOutput, conn)
	fifo := nPuzzle2.FIFO{}
	count := 0
	fifo.Push(puzzle.Root)
	for !fifo.IsEmpty() {
		child := fifo.Pop()
		children := child.GenerateChildren(puzzle.RowSize, puzzle.ColSize)
		for _, n := range children {
			if n.IsGoal() {
				path := n.GeneratePath()
				logOutput <- Reply{Title: "log", Body: "goal found "}
				for _ , pathItem := range path {
					fmt.Println("path to goal" , pathItem)
				}
				close(logOutput)
				return
			}
			if !isRepeated(visited, n.Numbers) {
				visited = append(visited, n.Numbers)
				count++
				logOutput <- Reply{Title: "log", Body: convertArrayToString(n.Numbers)}
				fifo.Push(n)
			}
		}
	}
	fmt.Println("goal not found")
	close(logOutput)
}

func sendLogs(logOutput chan Reply, conn *websocket.Conn) {
	for {
		select {
		case val, ok := <-logOutput:
			if ok {
				fmt.Println(val)
				err := conn.WriteJSON(val)
				if err != nil {
					fmt.Println("something wen wrong trying to send data to socket ", err)
				}
			} else {
				return
			}
		}
	}
}

func convertArrayToString(array []byte) string {
	return strings.Trim(strings.Join(strings.Fields(fmt.Sprint(array)), "-"), "[]")
}

func isRepeated(visited [][]byte, target []byte) bool {
	for _, n := range visited {
		if bytes.Equal(n, target) {
			return true
		}
	}
	return false
}
