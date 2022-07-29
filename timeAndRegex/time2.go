package main

import (
	"fmt"
	"time"
)

func main() {
	var fyear, fmonth, fday, fhour, fmin, fsec, sday, shour, smin, ssec int
	var ndiff time.Time
	fmt.Println("Input first time by year, month, days, hours, mins, secs (ex: 2004 3 5 5 15 0)")
	fmt.Scanln(&fyear, &fmonth, &fday, &fhour, &fmin, &fsec)
	fmt.Println("Input time to add or subtract by  days, hours, mins, secs (ex: 3 5 15 0)")
	fmt.Scanln(&sday, &shour, &smin, &ssec)
	fmt.Println("Input add or subtract(a or s)")
	var input string
	fmt.Scanln(&input)
	ftime := time.Date(fyear, time.Month(fmonth), fday, fhour, fmin, fsec, 0, time.Local)

	if input == "s" {
		ndiff = ftime.Add(time.Hour*time.Duration(-(shour+sday*24)) + time.Minute*time.Duration(-smin) + time.Second*time.Duration(-ssec))
	} else if input == "a" {

		ndiff = ftime.Add(time.Hour*time.Duration(shour+sday*24) + time.Minute*time.Duration(smin) + time.Second*time.Duration(ssec))

	}
	fmt.Println(ndiff)
}
