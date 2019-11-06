package bfs

import (
	"strconv"
	
	"github.com/cespare/xxhash"
	"github.com/gorilla/websocket"
	"nPuzzle"
	nPuzzle2 "structure/nPuzzle"
	"utils"
)

func Solve(puzzle *nPuzzle.NPuzzle, conn *websocket.Conn) {
	visited := make(map[uint64]bool)
	logOutput := make(chan utils.Reply)
	var count = 0
	logFilter := 100
	go utils.SendLogs(logOutput, conn)
	fifo := nPuzzle2.FIFO{}
	fifo.Push(puzzle.Root)
	for !fifo.IsEmpty() {
		child := fifo.Pop()
		children := child.GenerateChildren(puzzle.RowSize, puzzle.ColSize)
		for _, n := range children {
			if n.IsGoal() {
				path := n.GeneratePath()
				utils.SendBfsLogs(logOutput, "goal", "goal found ", utils.ConvertArrayToString(path))
				close(logOutput)
				return
			}
			hash := xxhash.Sum64(n.Numbers)
			if !visited[hash] {
				count++
				visited[hash] = true
				logFilter = utils.UpdateFilter(count)
				if count%logFilter == 0 {
					utils.SendBfsLogs(logOutput, "log", strconv.Itoa(count), []string{})
				}
				fifo.Push(n)
			}
		}
	}
	close(logOutput)
}
