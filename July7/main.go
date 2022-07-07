package main

import (
	"fmt"
	"mymod/sub"
)

func main() {
	var mapOfPeople map[int]sub.People
	p1 := sub.People{Name: "Joe", Age: 18}
	p2 := sub.People{Name: "Bob", Age: 28}
	p3 := sub.People{Name: "Steve", Age: 40}
	mapOfPeople = map[int]sub.People{1: p1, 2: p2, 3: p3}
	fmt.Println(mapOfPeople)
	//sub.MapSearch()
}
