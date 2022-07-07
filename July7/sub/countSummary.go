package sub

import (
	"fmt"
)

func Csummary() {
	var list []int
	var input int
reenter:
	fmt.Println("Input an amount of int between 1-100, enter '101' or higher when have entered all of your ints")
	fmt.Scanln(&input)

	if input < 101 {

		list = append(list, input)
		fmt.Println(list)
		goto reenter
	} else if input >= 101 {
		var g, l = list[0], list[0]
		var c int
		for x := 1; x < len(list); x++ {
			if g < list[x] {
				g = list[x]
			} else if l > list[x] {
				l = list[x]
			}
			c++
		}
		fmt.Println("the range is", l, "to", g, "and the count is:", c)
	}
}
