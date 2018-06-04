package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	m := len(s)/2
	go sum(s[:m], c)
	go sum(s[m:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x + y)
}
