package main

import "fmt"

func main() {
	RecursiveFib(0, 1, 50)
}
func RecursiveFib(f, s, n int64) {
	var nFib int64
	if n > 0 {
		nFib = f + s
		f = s
		s = nFib
		RecursiveFib(f, s, n-1)
	} else {
		fmt.Println(f)
	}
}
