package main

import (
	"fmt"
	"math"
)

//Carter Gray act2 7/1/22
func half_evenCheck(x int) (float64, bool) {
	if x%2 == 0 {
		return float64(x) / 2, true
	} else {
		return math.Floor(float64(x) / 2), false
	}
}

func main() {

	half, even := half_evenCheck(1)
	fmt.Println(half, even)

}
