package timeparse

import (
	"fmt"
	"testing"
)

type TestTime struct {
	time string
	ok   bool
}

func TestParseTime(t *testing.T) {
	table := make([]TestTime, 5)
	for i := 0; i < 5; i++ {
		fmt.Scanln(&table[i].time)
		fmt.Scanln(&table[i].ok)
	}
	for _, value := range table {
		_, err := parseTimne(value.time)
		if value.ok && err != nil {
			t.Errorf("%v %v, err should be nill", value.time, value.ok)
		}
	}
}

/*table := []struct {
	time string
	ok   bool
}{
	{"19:00:00", true},
	{"1:3:44", true},
	{"bad", false},
	{"1:-3:44", false},
	{"0:59:59", true},
	{"", false},
	{"11:22", false},
	{"aa:bb:cc", false},
	{"11:22:", false},
}*/
