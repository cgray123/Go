package main

import "fmt"

func main() {
	RecursiveFib(0, 1, 50)
}
func RecursiveFib(f, s, n int64) {
	var nFib int64
	if n > -1 {
		fmt.Println(f)
		nFib = f + s
		f = s
		s = nFib
		RecursiveFib(f, s, n-1)
	}
}
