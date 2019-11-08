package aStar

import (
	"fmt"
	
	"github.com/cespare/xxhash"
	"github.com/gorilla/websocket"
	nPuzzle2 "nPuzzle"
	"structure/nPuzzle"
)

func NPuzzleSolve(puzzle *nPuzzle2.NPuzzle, conn *websocket.Conn) {
	fmt.Println("inside astar")
	
	var list nPuzzle.LIST
	visited := make(map[uint64]bool)
	visited[xxhash.Sum64(puzzle.Root.Numbers)] = true
	list.InsertPriority(puzzle.Root)
	list.Print()
	
	for !list.IsEmpty() {
		target := list.PopPriority()
		children := target.GenerateChildren(puzzle.RowSize , puzzle.ColSize)
		for i, n := range children {
			n.Weight = i;
			list.InsertPriority(n)
		}
		list.Print()
	}
	
	
}
