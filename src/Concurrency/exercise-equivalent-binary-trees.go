package main

import (
	"fmt"
	"../../../../Go/misc/tour/src/golang.org/x/tour/tree"
)

func Walk(t *tree.Tree, ch chan int) {
	if t != nil {
		_walk(t, ch)
	}
	close(ch)
}

func _walk(t*tree.Tree, c chan int) {
	if t != nil {
		_walk(t.Left, c)
		c <- t.Value
		_walk(t.Right, c)
	}
}

func Same(t1, t2*tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	for v1 := range c1 {
		v2 := <- c2
		if v1 != v2{
			return false
		}
	}

	_, ok := <- c2
	if ok {
		return false
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))
}
