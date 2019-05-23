package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	current := 0
	return func() int {
		arg := current
		current++
		return fibonacci_impl(arg)
	}
}

func fibonacci_impl(i int) int {
	if i == 0 {
		return 0
	}
	if i == 1 {
		return 1
	}
	return fibonacci_impl(i-1) + fibonacci_impl(i-2)
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

