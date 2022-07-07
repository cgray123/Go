package main

import "fmt"
//Carter Gray act1 7/1/22
func sum(x []int) int {
	total := 0
	for _, num := range x {
		total += num
	}
	return total
}

func main() {
	var a = []int{1, 2, 3, 4}
	fmt.Println(sum(a))

}
