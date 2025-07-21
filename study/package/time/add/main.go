package main

import "time"

func main() {

}
func GetBeforeMonthDate(date string, month int) (string, error) {

	timeDate, err := time.ParseInLocation("2006-01-02", date, time.Local)
	if err != nil {
		return "", err
	}
	newTime := getBeforeMonthTime(timeDate, month)

	return newTime.Format("2006-01-02"), nil

}
func getBeforeMonthTime(t time.Time, month int) time.Time {
	beforeMonth := t.AddDate(0, -month, 0) // 前一个月的日期
	return beforeMonth
}
