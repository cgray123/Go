package main

import (
	"fmt"
	"regexp"
)

func main() {
	str := "Hello I am 20 years 10 month and 20 days old cat cot  my DOB is 2000-11-05 my telephone number is +11235551234 my IP address is 192.166.255 192.170.1"
	//regex pattern for finding a ip in the range 192.160.1 to 192.170.1
	re1 := regexp.MustCompile(`192\.1(6[0-9]\.(2[0-5][0-5]|1[0-9][0-9]|[1-9][0-9]|[1-9])|70\.1)`)
	//regex pattern for finding words that start with c and end with t
	re3 := regexp.MustCompile("c([a-z]+)t")

	mat := re3.FindAllString(str, -1)
	for _, ind := range mat {
		fmt.Println(ind)
	}
	mat2 := re1.FindAllString(str, -1)
	for _, ind := range mat2 {
		fmt.Println(ind)
	}
}
