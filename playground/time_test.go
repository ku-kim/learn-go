package playground

import (
	"fmt"
	"testing"
	"time"
)

func TestTimemethod(t *testing.T) {
	// 현재날짜 -8d
	y, m, d := time.Now().AddDate(0, 0, -8).UTC().Date()
	fmt.Println(fmt.Sprintf("%d-%d-%d", y, m, d)) // 2023-3-9

}
