package main

import "fmt"

//Carter Gray act3 7/1/22
func findGreatest(nums ...int) int {
	var greatest int = nums[0]
	for _, n := range nums {
		if greatest < n {
			greatest = n
		}
	}
	return greatest
}

func main() {

	fmt.Println(findGreatest(-10, -2, -8, -4, -1, -3))

}
