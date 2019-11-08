package nPuzzle

import (
	"fmt"
	
	"nPuzzle"
)

type LIST struct {
	List [] *nPuzzle.Node
}

func (list *LIST) IsEmpty() bool {
	return len(list.List) == 0
}

func (list *LIST) InsertPriority(node *nPuzzle.Node) {
	if list.IsEmpty() {
		list.List = append(list.List, node)
		return
	}
	for index, val := range list.List {
		if val.Weight > node.Weight {
			insert(list, index, node)
			return
		}
	}
	list.List = append(list.List, node)
}

func (list *LIST) PopPriority() *nPuzzle.Node {
	tail := list.List[0]
	list.List = list.List[1:]
	return tail
}

func insert(list *LIST, index int, node *nPuzzle.Node) {
	list.List = append(list.List, node)
	copy(list.List[index+1:], list.List[index:])
	list.List[index] = node
	// list.List = append(list.List[:index], append([]*nPuzzle.Node{node}, list.List[index:]...)...)
}

func (list *LIST) Print() {
	for _, n := range list.List {
		fmt.Println(n.Numbers, n.Weight)
	}
}
