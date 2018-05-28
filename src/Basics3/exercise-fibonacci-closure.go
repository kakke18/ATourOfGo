package main

import "fmt"

func fibonacci() func(int) int {
	num, sum := 1, 1
	return func(x int) int {
		if(x == 0 || x == 1) {
			return sum
		}
		else {
			sum += num
			return sum
		}
	}
}

func main() {
	f :- fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f(i))
	}
}
