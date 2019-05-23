package main

import (
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	DoWalk(t, ch)
	close(ch)
}

func DoWalk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}

	DoWalk(t.Left, ch)
	ch <- t.Value
	DoWalk(t.Right, ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	return EqualSlices(FetchElements(t1), FetchElements(t2))
}

func EqualSlices(s1, s2 []int) bool {
	if len(s1) != len(s2) {
		return false
	}
	for i, v := range s1 {
		if s2[i] != v {
			return false
		}
	}

	return true
}

func FetchElements(t *tree.Tree) (elems []int) {
	ch := make(chan int, 10)
	go Walk(t, ch)
	elems = make([]int, 0, 10)
	for v := range ch {
		elems = append(elems, v)
	}
	return
}

func main() {
	if !EqualSlices(FetchElements(tree.New(1)), []int{1,2,3,4,5,6,7,8,9,10}) {
		panic("!EqualSlices(FetchElements(tree.New(1),[]int{1,2,3,4,5,6,7,8,9,10})")
	}

	if !Same(tree.New(1), tree.New(1)) {
		panic("!Same(tree.New(1), tree.New(1))")
	}

	if Same(tree.New(1), tree.New(2)) {
		panic("Same(tree.New(1), tree.New(2))")
	}
}

