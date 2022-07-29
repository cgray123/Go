package main

import (
	"fmt"
	"math"
	"time"
)
//gets the differents bewteen 2 dates
func main() {
	var fyear, fmonth, fday, syear, smonth, sday int

	fmt.Println("Input first time by year, month, day (ex: 2004 3 5)")
	fmt.Scanln(&fyear, &fmonth, &fday)
	fmt.Println("Input second time by year, month, day (ex: 2004 3 5) must be greater than the first")
	fmt.Scanln(&syear, &smonth, &sday)
	ftime := time.Date(fyear, time.Month(fmonth), fday, 0, 0, 0, 0, time.Local)
	stime := time.Date(syear, time.Month(smonth), sday, 0, 0, 0, 0, time.Local)

	ndiff := stime.Sub(ftime)
	days := math.Floor(ndiff.Hours() / 24)
	weeks := math.Floor(days / 7)
	months := math.Floor(weeks / 4)
	years := math.Floor(months / 12)
	temp := months - years*12

	fmt.Println(years, "Years", temp, "months", days-temp*7-years*12*7, "days")
	fmt.Println(months, "months", days-months*4*7, "days")
	fmt.Println(weeks, "weeks", days-weeks*7, "days")
	fmt.Println(days, "days")
	fmt.Println(ndiff.Hours(), "Hours")
	fmt.Println(ndiff.Hours()*60, "minutes")
}
