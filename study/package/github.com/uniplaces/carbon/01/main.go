package main

import (
	"fmt"
	"github.com/uniplaces/carbon"
)

func main() {
	fmt.Printf("tomorrow time is %+v\n", carbon.Now().AddDay().AddWeek())
	fmt.Printf("now time is %s\n", carbon.Now().DateTimeString())
	today, err := carbon.NowInLocation("America/New_York")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("japan now is %s\n", today)
	fmt.Println("weekday?", carbon.Now().IsWeekday())
}
