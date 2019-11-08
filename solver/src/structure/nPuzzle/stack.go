package nPuzzle

import (
	"fmt"
	
	"nPuzzle"
)

type STACK struct {
	Stack [] *nPuzzle.Node
}

func (stack *STACK) IsEmpty() bool {
	return len(stack.Stack) == 0
}

func (stack *STACK) Pop() *nPuzzle.Node {
	if len(stack.Stack) == 0 {
		return nil
	}
	
	tailIndex := len(stack.Stack) - 1
	tail := stack.Stack[tailIndex]
	stack.Stack = stack.Stack[:tailIndex]
	return tail
}

func (stack *STACK) Push(node *nPuzzle.Node) {
	stack.Stack = append(stack.Stack, node)
}

func (stack *STACK) Print() {
	for _, n := range stack.Stack {
		fmt.Println(n)
	}
}
