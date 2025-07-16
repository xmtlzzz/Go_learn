package main

import "fmt"

type Weekday int

const (
	Sunday Weekday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Satday
)

func echoWeekDay(w Weekday) {
	fmt.Printf("现在是星期%d\n", w)
}

var a int = 1

func main() {
	echoWeekDay(1)
	echoWeekDay(Monday)
}
