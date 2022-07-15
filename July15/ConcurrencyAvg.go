package main

import "fmt"

func avg(a []int, c chan float64, d float64) {
	sum := float64(0)
	for _, v := range a {
		sum += float64(v)
	}
	c <- sum / d 
}

func main() {
	s := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}

	c := make(chan float64)
	d := float64(len(s))
	go avg(s[:5], c, d)
	go avg(s[5:10], c, d)
	go avg(s[10:15], c, d)
	go avg(s[15:], c, d)
	v, x, y, z := <-c, <-c, <-c, <-c 

	fmt.Println(v + x + y + z)
}
