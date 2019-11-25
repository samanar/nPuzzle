package aStar

import (
	"container/list"
	"strconv"
	"utils"

	"github.com/cespare/xxhash"
	"github.com/gorilla/websocket"
	nPuzzle2 "nPuzzle"
	"structure/nPuzzle"
)

func replaceIfBetter(list *list.List) {

}

func NPuzzleSolve(puzzle *nPuzzle2.NPuzzle, conn *websocket.Conn) {
	logFilter := 100
	count := 0
	logOutput := make(chan utils.Reply)
	go utils.SendLogs(logOutput, conn)

	var list nPuzzle.LIST
	visited := make(map[uint64]bool)
	visited[xxhash.Sum64(puzzle.Root.Numbers)] = true
	list.InsertPriority(puzzle.Root)
	visited[xxhash.Sum64(puzzle.Root.Numbers)] = true

	for !list.IsEmpty() {
		target := list.PopPriority()
		if target.IsGoal() {
			path := target.GeneratePath()
			utils.SendBfsLogs(logOutput, "goal", "goal found ", utils.ConvertArrayToString(path))
			close(logOutput)
			return
		}
		for _, n := range target.GenerateChildren(puzzle.RowSize, puzzle.ColSize) {
			hash := xxhash.Sum64(n.Numbers)
			if !visited[hash] {
				visited[hash] = true
				count++
				logFilter = utils.UpdateFilter(count)
				if count%logFilter == 0 {
					utils.SendBfsLogs(logOutput, "log", strconv.Itoa(count), []string{})
				}
				n.Height = target.Height + 1
				n.Weight = target.Height + n.CalculateManhattanDistance(puzzle.RowSize, puzzle.ColSize)
				list.InsertPriority(n)
			} else {
				// check what to do with repeated ones
				// add replace for list here
			}
		}
	}

}
