package main

import "fmt"

//Carter Gray act4 7/1/22
func makeOddGenerator() func() int {
	i := 1
	return func() (x int) {
		x = i
		i += 2
		return
	}
}
func main() {

	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}
