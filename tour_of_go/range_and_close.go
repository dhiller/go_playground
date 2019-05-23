package main

import (
	"fmt"
)

type fib struct {
	x, fibOfX int
}

func (f fib) String() (string) {
	return fmt.Sprintf("fib(%v) = %v", f.x, f.fibOfX)
}

func fibonacci(n int, c chan fib) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- fib{i, x}
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan fib, 30)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
