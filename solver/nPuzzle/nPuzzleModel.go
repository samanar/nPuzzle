package nPuzzle

type NPuzzle struct {
	Numbers []int
	RowSize int
	ColSize int
	parent *NPuzzle
	children []*NPuzzle
}

func (nPuzzle *NPuzzle) convertIndexToCoordinates(index int) (int , int) {
	i :=  index / nPuzzle.RowSize
	j :=  index - (i * nPuzzle.RowSize)
	if j < 0 {
		j = j * -1
	}
	return i , j
}

func (nPuzzle *NPuzzle) findZero() int {
	for index ,n := range nPuzzle.Numbers {
		if n == 0 {
			return index
		}
	}
	return -1
}

func (nPuzzle *NPuzzle)  GenerateChildren() {
	zeroIndex := nPuzzle.findZero()
	i , j := nPuzzle.convertIndexToCoordinates(zeroIndex)
}