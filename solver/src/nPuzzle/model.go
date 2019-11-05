package nPuzzle

import (
	"log"
)

type NPuzzle struct {
	RowSize byte
	ColSize byte
	Root    *Node
}

type Node struct {
	Numbers []byte
	Parent  *Node
}

func convertIndexToCoordinates(index byte, rowSize, colSize byte) (byte, byte) {
	i := byte(index / rowSize)
	j := byte(index - (i * rowSize))
	if j < 0 {
		j = -j
	}
	return i, j
}

func convertCoordinatesToIndex(i, j, rowSize byte) byte {
	return i*rowSize + j
}

func (node *Node) findZero() (bool , byte) {
	for index, n := range node.Numbers {
		if n == 0 {
			return false , byte(index)
		}
	}
	return true , 0
}

func (node *Node) availableMoves(rowSize, colSize byte) (byte, byte, [4]string) {
	var moves [4]string
	err , zeroIndex := node.findZero()
	if err {
		log.Fatal("error when trying to find zero index")
	}
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

func swapNumbers(numbers []byte, first, second byte) []byte {
	result := make([]byte, len(numbers))
	copy(result, numbers)
	result[first], result[second] = result[second], result[first]
	return result
}

func (node Node) generateChild(direction string, zeroI, zeroJ, rowSize, colSize byte) *Node {
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

func (node Node) GenerateChildren(rowSize, colSize byte) []*Node {
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
	misMatch := false
	var i byte
	for i = 0; i < byte(len(node.Numbers)) - 1; i++ {
		if node.Numbers[i] != i + 1 {
			misMatch = true
			break
		}
	}
	return !misMatch
}

func (node *Node) GeneratePath() [][] byte{
	var path [][] byte
	parent := node

	for parent != nil {
		path = append([][] byte {parent.Numbers} , path...)
		parent = parent.Parent
	}

	return path
}
