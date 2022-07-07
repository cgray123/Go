package sub

import "fmt"

func SliceSwap() {
	s := make([][]interface{}, 3)

	for j := 0; j < len(s); j++ {
		for x := 0; x < len(s); x++ {
			var input interface{}
			fmt.Scanln(&input)
			s[j][x] = input
		}
	}
	fmt.Println("would you like to swap rows or columns? input r or c")
	var y string
	fmt.Scanln(&y)
	if y == "c" {
		for i := 0; i < len(s); i++ {
			s[i][0], s[i][len(s)-1] = s[i][len(s)-1], s[i][0]
		}
	} else if y == "r" {
		for i := 0; i < len(s); i++ {
			s[0][i], s[len(s)-1][i] = s[len(s)-1][i], s[0][i]
		}
	}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}
