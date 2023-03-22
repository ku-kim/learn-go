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
	t1 := baseTime(time.Now())

	fmt.Println(t1) // 2023-03-21 00:00:00 +0900 KST, 함수 시작하는 날짜 00:00:00
	fmt.Println(time.Now())
}

func baseTime(t time.Time) time.Time {
	loc, _ := time.LoadLocation("Asia/Seoul")
	year, month, day := t.In(loc).Date()
	onTime := time.Date(year, month, day, 0, 0, 0, 0, loc)
	return onTime
}
