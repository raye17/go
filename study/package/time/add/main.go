package main

import (
	"fmt"
	"time"
)

func main() {
	s, err := GetBeforeMonthDate("2025-07-22", 6)
	if err != nil {
		panic(err)
	}
	fmt.Println(s)
}
func GetBeforeMonthDate(date string, month int) (string, error) {

	timeDate, err := time.ParseInLocation("2006-01-02", date, time.Local)
	if err != nil {
		return "", err
	}
	newTime := timeDate.AddDate(0, -month, 0)

	return newTime.Format("2006-01-02"), nil

}
