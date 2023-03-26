package playground

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeMethod(t *testing.T) {
	// 현재날짜 -8d
	y, m, d := time.Now().AddDate(0, 0, -8).UTC().Date()
	fmt.Println(fmt.Sprintf("%d-%d-%d", y, m, d)) // 2023-3-9

}

func TestBaseTime(t *testing.T) {
	t1 := BaseTime(time.Now())

	fmt.Println(t1) // 2023-03-21 00:00:00 +0900 KST, 함수 시작하는 날짜 00:00:00
	fmt.Println(time.Now())
}

func BaseTime(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Seoul")
	year, month, day := t.In(loc).Date()
	onTime := time.Date(year, month, day, 0, 0, 0, 0, loc)
	return onTime
}

func TestStrToTime(t *testing.T) {
	s := "2023-03-23"
	parse, err := time.Parse("2006-01-02", s)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(parse)

	//localTime := GetLocalTime(time.Now())
	//fmt.Println(localTime.Format("2006-01-02"))

	fmt.Println(time.Now().Format("2006-01-02"))
}

var loc *time.Location

func GetLocalTime(t time.Time) time.Time {
	return t.In(loc)
}

func TestDateCompare(t *testing.T) {
	today := BaseTime(time.Now())
	tomorrow := today.AddDate(0, 0, 1)

	fmt.Println("today", today)
	fmt.Println("tomorrow", tomorrow)

	fmt.Println(today.After(tomorrow))
	fmt.Println(today.Before(tomorrow))
}
