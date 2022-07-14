package main

import (
	"fmt"
	"math"
	"os"
	"time"
)

var (
	printNow chan bool
	i        int
)

func main() {
	printNow = make(chan bool)
	go sender()
	go printer()

	for {
	}
}
func printer() {
	for {
		if _, ok := <-printNow; ok {
			flag := true
			for x := 2; x <= int(math.Floor(float64(i)/2)); x++ {
				if i%x == 0 {
					flag = false
				}
			}
			if flag && i > 1 {
				fmt.Print(" is prime\n")
			} else {
				fmt.Print(" is not prime\n")
			}

		}

	}
}
func sender() {
	for {
		for i = 0; i <= 100; i++ {
			fmt.Print(i)

			printNow <- true

			time.Sleep(2 * time.Millisecond)
		}
		os.Exit(0)
	}
}
