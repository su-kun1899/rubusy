package main

import (
	"reflect"
	"testing"
	"time"
)

func TestMatch(t *testing.T) {
	jobs := []CronJob{
		Parse("15 9 * * * /tmp/hoge.sh"),
	}
	condition := targetTime{
		from: time.Date(2017, 10, 4, 18, 0, 0, 0, time.UTC),
		to:   time.Date(2017, 10, 4, 19, 0, 0, 0, time.UTC),
	}

	for _, expected := range jobs {
		contains, actual := expected.match(condition)
		if !contains {
			t.Fatalf("This job should be mathch: %+v \n", expected)
		}
		if !reflect.DeepEqual(&expected, actual) {
			t.Fatalf("expected: %+v but actual: %+v\n", expected, actual)
		}
	}
}
