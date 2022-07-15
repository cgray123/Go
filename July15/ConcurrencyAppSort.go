package main

import (
	"fmt"
	"math/rand"
	"os"
	"sort"
	"time"
)

var (
	printNow chan bool
	l        []int
)

func main() {
	printNow = make(chan bool)
	go gen()
	go Sort()
	for {

	}
}
func gen() {
	for i := 0; i < 10; i++ {
		l = append(l, rand.Intn(100))
		fmt.Println("appended:", l)
		printNow <- true
		time.Sleep(time.Millisecond)
	}
	os.Exit(0)
}
func Sort() {
	for {
		if _, ok := <-printNow; ok {
			sort.Ints(l)
			fmt.Println("sorted:", l)
		}
	}
}
