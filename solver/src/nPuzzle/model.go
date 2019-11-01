package nPuzzle

import (
	"fmt"
)

type NPuzzle struct {
	RowSize int
	ColSize int
	Root    *Node
}

type Node struct {
	Numbers  []int
	Parent   *Node
	Children []*Node
	Test     int
}

func convertIndexToCoordinates(index int, rowSize, colSize int) (int, int) {
	i := index / rowSize
	j := index - (i * rowSize)
	if j < 0 {
		j = j * -1
	}
	return i, j
}

func convertCoordinatesToIndex(i, j, rowSize int) int {
	return i*rowSize + j
}

func (node *Node) findZero() int {
	for index, n := range node.Numbers {
		if n == 0 {
			return index
		}
	}
	return -1
}

func (node *Node) availableMoves(rowSize, colSize int) (int, int, [4]string) {
	var moves [4]string
	zeroIndex := node.findZero()
	zeroI, zeroJ := convertIndexToCoordinates(zeroIndex, rowSize, colSize)
	if zeroI != 0 {
		moves[0] = "up"
	}
	if zeroJ != colSize-1 {
		moves[1] = "right"
	}
	if zeroI != rowSize-1 {
		moves[2] = "down"
	}
	if zeroJ != 0 {
		moves[3] = "left"
	}
	
	return zeroI, zeroJ, moves
}

func swapNumbers(numbers *[]int, first, second int) {
	(*numbers)[first], (*numbers)[second] = (*numbers)[second], (*numbers)[first]
	
}

func (node Node) generateChild(direction string, zeroI, zeroJ, rowSize, colSize int) Node {
	numbers := node.Numbers
	secondI := zeroI
	secondJ := zeroJ
	switch direction {
	case "up":
		secondI--
	case "down":
		secondI++
	case "left":
		secondJ--
	case "right":
		secondJ++
	}
	
	swapNumbers(
		&numbers,
		convertCoordinatesToIndex(zeroI, zeroJ, rowSize),
		convertCoordinatesToIndex(secondI, secondJ, rowSize))
	
	return Node{
		Numbers: numbers,
		Parent:  &node,
	}
}

func copyAndAppend(i []*Node, item *Node) []*Node {
	j := make([]*Node, len(i), len(i)+1)
	copy(j, i)
	return append(j, item)
}

func (node *Node) GenerateChildren(rowSize, colSize int) []*Node {
	var children []*Node
	zeroI, zeroJ, moves := node.availableMoves(rowSize, colSize)
	for _, n := range moves {
		if len(n) == 0 {
			continue
		}
		child := node.generateChild(n, zeroI, zeroJ, rowSize, colSize)
		children = append(children, &child)
		for _, n := range children {
			fmt.Println(*n)
		}
		fmt.Println("-----------------------")
	}
	// adding generated children to node children ??
	return children
}
