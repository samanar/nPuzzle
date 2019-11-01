package nPuzzle

type NPuzzle struct {
	RowSize int
	ColSize int
	Root    *Node
}

type Node struct {
	Numbers  []int
	parent   *Node
	children []*Node
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

func (node *Node) availableMoves(rowSize, colSize int) (int, int, [4]bool) {
	var moves [4]bool
	zeroIndex := node.findZero()
	zeroI, zeroJ := convertIndexToCoordinates(zeroIndex , rowSize , colSize)
	if zeroI != 0 {
		moves[0] = true
	}
	if zeroJ != colSize-1 {
		moves[1] = true
	}
	if zeroI != rowSize-1 {
		moves[2] = true
	}
	if zeroJ != 0 {
		moves[3] = true
	}

	return zeroI, zeroJ, moves
}

func SwapNumbers(numbers *[]int, first, second int) {
	(*numbers)[first], (*numbers)[second] = (*numbers)[second], (*numbers)[first]
}

func (node Node) generateChild(direction string, zeroI, zeroJ int) *Node {
	numbers := node.Numbers
	switch direction {
	case "up":
		//numbers
	}

}

func (node *Node) GenerateChildren() {

}
