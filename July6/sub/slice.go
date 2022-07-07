package sub

import "fmt"

func Slice() {
	fmt.Println("input n where nxn is the size of the matix")
	var z int
	fmt.Scanln(&z)
	s := make([][]int, z)
	for j := 0; j < len(s); j++ {
		s[j] = RowSlice(z)
	}
	var count int
	for i := 0; i < len(s); i++ {
		count++
		for x := 0; x < count; x++ {
			fmt.Print(fmt.Sprint(s[i][x]) + " ")
		}
		fmt.Println()
	}
}
func RowSlice(z int) []int {
	s := make([]int, z)
	var input int
	fmt.Printf("input a int then hit enter, repeat %d more times", z-1)
	fmt.Println()
	for x := 0; x < len(s); x++ {
		fmt.Scanln(&input)
		s[x] = input
	}
	return s
}
