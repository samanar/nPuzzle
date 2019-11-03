package bfs

import (
	"fmt"
	"nPuzzle"
)

type FIFO struct {
	queue []*nPuzzle.Node
}

func Solve(puzzle *nPuzzle.NPuzzle) {
	var fifo FIFO
	fifo.push(puzzle.Root)
	children := puzzle.Root.GenerateChildren(puzzle.RowSize , puzzle.ColSize)
	for _,n := range children {
		fifo.push(n)
	}
	fifo.Print()
	fifo.Pop()
	fifo.Pop()
	fifo.Pop()
	fmt.Println("after poping")
	fifo.Print()
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

