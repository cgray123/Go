package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var wait chan bool

func main() {
	wg.Add(2)
	wait = make(chan bool)
	go func() {
		for i := 0; i < 101; i++ {
			if ok := <-wait; ok {
				fmt.Println(i)
				time.Sleep(time.Millisecond * 45)
			}

		}
		wg.Done()
	}()
	go func() {
		for i := 100; i >= 0; i-- {

			fmt.Println(i)
			wait <- true
			time.Sleep(time.Millisecond * 45)

		}
		wg.Done()
	}()
	wg.Wait()
}
