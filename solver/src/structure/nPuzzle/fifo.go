package nPuzzle

import (
	"fmt"
	
	"nPuzzle"
)

type FIFO struct {
	queue [] *nPuzzle.Node
}

func (fifo *FIFO) IsEmpty() bool {
	return len(fifo.queue) == 0
}

func (fifo *FIFO) Pop() *nPuzzle.Node {
	if len(fifo.queue) == 0 {
		return nil
	}
	
	head := fifo.queue[0]
	fifo.queue = fifo.queue[1:]
	return head
}

func (fifo *FIFO) Push(node *nPuzzle.Node) {
	fifo.queue = append(fifo.queue, node)
}

func (fifo *FIFO) Print() {
	for _, n := range fifo.queue {
		fmt.Println(n)
	}
}
