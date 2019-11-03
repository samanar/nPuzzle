package nPuzzle

type NPuzzle struct {
	RowSize int
	ColSize int
	Root    *Node
}

type Node struct {
	Numbers []int
	Parent  *Node
	Test    int
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

func swapNumbers(numbers []int, first, second int) []int {
	result := make([]int, len(numbers))
	copy(result, numbers)
	result[first], result[second] = result[second], result[first]
	return result
}

func (node Node) generateChild(direction string, zeroI, zeroJ, rowSize, colSize int) *Node {
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
	result := new(Node)
	result.Numbers = swapNumbers(
		node.Numbers,
		convertCoordinatesToIndex(zeroI, zeroJ, rowSize),
		convertCoordinatesToIndex(secondI, secondJ, rowSize))
	result.Parent = &node
	return result
}

func (node Node) GenerateChildren(rowSize, colSize int) []*Node {
	var children []*Node
	zeroI, zeroJ, moves := node.availableMoves(rowSize, colSize)
	for _, n := range moves {
		if len(n) == 0 {
			continue
		}
		child := node.generateChild(n, zeroI, zeroJ, rowSize, colSize)
		children = append(children, child)
	}
	
	return children
}

func (node Node) IsGoal() bool {
	isGoal := true
	for i := 1; i <= len(node.Numbers); i++ {
		if node.Numbers[i-1] != i {
			isGoal = false
		}
	}
	return isGoal
}
