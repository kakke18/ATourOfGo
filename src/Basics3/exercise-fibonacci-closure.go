package main

import "fmt"

func fibonacci() func(int) int {
	i, j := 1, 0
	return func(x int) int {
		if x == 0 {
			return j
		} else if x == 1 {
			return i
		} else {
			i, j = i + j, i
			return i
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
