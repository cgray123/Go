package main

import (
	"fmt"
	"math"
)

func inc(s string) chan int {
	out := make(chan int)
	go func() {

		for i := 0; i <= 100; i++ {
			out <- i

		}
		close(out)
	}()

	return out
}
func puller(c chan int) chan int {
	out := make(chan int)
	go func() {

		for n := range c {
			flag := true
			for x := 2; x <= int(math.Floor(float64(n)/2)); x++ {
				if n%x == 0 {

					flag = false
					break
				}
			}
			if flag && n > 1 {

				out <- n
			}

		}

		close(out)
	}()
	return out
}

func main() {
	n1 := inc("Genarate1")

	p1 := puller(n1)
	for n := range p1 {
		fmt.Println(n, "is prime")
	}

}
