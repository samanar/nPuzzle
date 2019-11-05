package nPuzzle

import (
	"fmt"
	
	"nPuzzle"
)

type FIFO struct {
	Queue [] *nPuzzle.Node
}

func (fifo *FIFO) IsEmpty() bool {
	return len(fifo.Queue) == 0
}

func (fifo *FIFO) Pop() *nPuzzle.Node {
	if len(fifo.Queue) == 0 {
		return nil
	}
	
	head := fifo.Queue[0]
	fifo.Queue = fifo.Queue[1:]
	return head
}

func (fifo *FIFO) Push(node *nPuzzle.Node) {
	fifo.Queue = append(fifo.Queue, node)
}

func (fifo *FIFO) Print() {
	for _, n := range fifo.Queue {
		fmt.Println(n)
	}
}
