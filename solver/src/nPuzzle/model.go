package nPuzzle

import "fmt"

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

func (node Node) generateChild(direction string, zeroI, zeroJ, rowSize, colSize int) *Node {
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
	
	result := new(Node)
	result.Numbers = numbers
	result.Parent = &node
	return result
}

func copyAndAppend(i []*Node, item *Node) []*Node {
	j := make([]*Node, len(i), len(i)+1)
	copy(j, i)
	return append(j, item)
}

// func (node Node) GenerateChildren(rowSize, colSize int) [4]Node {
// 	var children [4]Node
// 	var child Node
// 	zeroI, zeroJ, moves := node.availableMoves(rowSize, colSize)
// 	for index, n := range moves {
// 		if len(n) == 0 {
// 			continue
// 		}
// 		child = node.generateChild(n, zeroI, zeroJ, rowSize, colSize)
// 		children[index] = child
// 		for _, test := range children {
// 			fmt.Println(test)
// 		}
// 		fmt.Println("-----------------------")
// 	}
// 	// adding generated children to node children ??
// 	return children
// }

func (node Node) GenerateChildren(rowSize, colSize int) []*Node {
	testing := make([]*Node, 4)
	// var child Node
	zeroI, zeroJ, _ := node.availableMoves(rowSize, colSize)
	child1 := node.generateChild("up", zeroI, zeroJ, rowSize, colSize)
	child2 := node.generateChild("right", zeroI, zeroJ, rowSize, colSize)
	fmt.Println("child 1", child1)
	fmt.Println("child 2", child2)
	
	testing[0] = child1
	testing[1] = child2
	
	// adding generated children to node children ??
	return testing
}
