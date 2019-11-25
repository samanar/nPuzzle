package dfs

import (
	"strconv"

	"github.com/cespare/xxhash"
	"github.com/gorilla/websocket"
	"nPuzzle"
	nPuzzle2 "structure/nPuzzle"
	"utils"
)

func NPuzzleSolve(puzzle *nPuzzle.NPuzzle, conn *websocket.Conn) {
	stack := nPuzzle2.STACK{}
	visited := make(map[uint64]bool)
	logFilter := 100
	count := 0
	logOutput := make(chan utils.Reply)
	go utils.SendLogs(logOutput, conn)
	stack.Push(puzzle.Root)
	visited[xxhash.Sum64(puzzle.Root.Numbers)] = true
	for !stack.IsEmpty() {
		target := stack.Pop()
		for _, n := range target.GenerateChildren(puzzle.RowSize, puzzle.ColSize) {
			hash := xxhash.Sum64(n.Numbers)
			if !visited[hash] {
				visited[hash] = true
				if n.IsGoal() {
					path := n.GeneratePath()
					utils.SendBfsLogs(logOutput, "goal", "goal found ", utils.ConvertArrayToString(path))
					close(logOutput)
					return
				}
				count++
				logFilter = utils.UpdateFilter(count)
				if count%logFilter == 0 {
					utils.SendBfsLogs(logOutput, "log", strconv.Itoa(count), []string{})
				}
				stack.Push(n)
			}
		}
	}
	close(logOutput)
}
