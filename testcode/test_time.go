package main

import (
	"fmt"
	"time"
)

func LastMonth(t time.Time) time.Time {
	deltaDay := 0
	for {
		lastMonth := t.AddDate(0, -1, deltaDay)
		if lastMonth.Month() == t.Month() {
			deltaDay--
			continue
		}
		return lastMonth
	}
	return time.Time{}
}
func main() {
	//fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
	p := fmt.Println

	// 当前时间
	// p(time.Now().Format("2006-01-02 15:04:05"))

	// t := time.Now().UnixNano()
	// p(t)
	p(1.0 + 1)
	// t := time.Now()
	// newT := t.AddDate(0, -1, 0).Format("2006-01-02")
	// p(newT)
	// p(LastMonth(time.Now().AddDate(0, -1, 26)).Format("2006-01-02"))
}
