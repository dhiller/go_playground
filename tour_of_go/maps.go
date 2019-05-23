package main

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	r := make(map[string]int)
	words := strings.Split(s, " ")
	for _, w := range words {
		r[w] = r[w] + 1
	}
	return r
}

func main() {
	wc.Test(WordCount)
}

