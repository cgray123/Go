package main

import "fmt"

type person struct {
	name string
	age  int
	rent int
}

func main() {
	c := make(chan person)

	go func() {
		c <- person{"Bob", 21, 1865}
		c <- person{"Joe", 54, 1586}
		c <- person{"Steve", 85, 4866}
		c <- person{"Jill", 18, 1568}
		c <- person{"Frank", 9, 2588}
		c <- person{"Tom", 27, 2577}
		close(c)
	}()
	var tot int
	for n := range c {
		tot += n.rent
	}
	fmt.Println(tot)
}
