package main
import (
	"fmt"
	"math"
	"time"
)

func main() {
	var fweek, fday, fhour, fmin, fsec, sweek, sday, shour, smin, ssec int
	var days, times float64
	fmt.Println("Input first time by weeks, days, hours, mins, secs (ex: 4 3 5 15 0)")
	fmt.Scanln(&fweek, &fday, &fhour, &fmin, &fsec)
	fmt.Println("Input second time by weeks, days, hours, mins, secs (ex: 4 3 5 15 0) must be greater than the first")
	fmt.Scanln(&sweek, &sday, &shour, &smin, &ssec)
	fmt.Println("Input add or subtract(a or s)")
	var input string
	fmt.Scanln(&input)
	fday += fweek * 7
	sday += sweek * 7
	ftime := time.Date(2020, 07, fday, fhour, fmin, fsec, 0, time.Local)
	stime := time.Date(2020, 07, sday, shour, smin, ssec, 0, time.Local)
	diff := ftime.Sub(stime)
	if input == "s" {
		days = math.Floor(-diff.Hours() / 24)
		fmt.Println(days, "days", -diff.Hours()-days*24, "hours", -diff.Minutes()-days*24*60, "minutes", -diff.Seconds()-days*24*60*60, "seconds")
		times = -diff.Hours()
		fmt.Println(times/24, "days")
		fmt.Println(times, "Hours")
		fmt.Println(times*60, "Minutes")
		fmt.Println(times*60*60, "Seconds")
	} else if input == "a" {

		ndiff := ftime.Add(time.Hour*time.Duration(shour+sday*24) + time.Minute*time.Duration(smin) + time.Second*time.Duration(ssec))

		fmt.Println(ndiff.Day(), "days")
		fmt.Println(ndiff.Hour(), "Hours")
		fmt.Println(ndiff.Minute(), "Minutes")
		fmt.Println(ndiff.Second(), "Seconds")
	}

}
