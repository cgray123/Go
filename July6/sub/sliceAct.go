package sub

import (
	"fmt"
	"strings"
)

func Input() string {
	fmt.Println("Input word")
	var x string
	fmt.Scanln(&x)
	return strings.ToLower(x)
}
func SliceActMain() {
	var wordArr [10]string
	var total int
	amount := len(wordArr)
	for i := 0; i < amount; i++ {
		wordArr[i] = Input()
		total += len(wordArr[i])
	}
	total /= amount
	var less, more []string
	for x := 0; x < amount; x++ {
		if len(wordArr[x]) > total {
			less = append(less, wordArr[x])
		} else if len(wordArr[x]) < total {
			more = append(more, wordArr[x])
		}
	}
	fmt.Println("All of the inputted words:", wordArr)
	fmt.Println("All of the shorter than the average words:", less)
	fmt.Println("All of the greater than the average words:", more)
}
