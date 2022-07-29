package main

import (
	"fmt"
	"regexp"
)

func main() {
	str1 := "aaoeu1056"
	str := "in on io inn innn"
	//act1
	re1 := regexp.MustCompile(`^[a-zA-Z0-9]+$`)
	//act2
	re2 := regexp.MustCompile("in*")
	//act3
	re3 := regexp.MustCompile("in+")
	//act4
	re4 := regexp.MustCompile("i(n{1}|n{2})[^n]")
	//act5
	re5 := regexp.MustCompile("in{3}")

	mat := re1.FindAllString(str1, -1)
	for _, ind := range mat {
		fmt.Println("1", ind)
	}
	mat2 := re2.FindAllString(str, -1)
	for _, ind := range mat2 {
		fmt.Println("2", ind)
	}
	mat3 := re3.FindAllString(str, -1)
	for _, ind := range mat3 {
		fmt.Println("3", ind)
	}
	mat4 := re4.FindAllString(str, -1)
	for _, ind := range mat4 {
		fmt.Println("4", ind)
	}
	mat5 := re5.FindAllString(str, -1)
	for _, ind := range mat5 {
		fmt.Println("5", ind)
	}
}
