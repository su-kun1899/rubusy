package main

import (
	"reflect"
	"testing"
	"time"
)

func TestMatch(t *testing.T) {
	jobs := []CronJob{
		Parse("15 18 * * * /tmp/hoge.sh"),
		Parse("*/5 * * * * /tmp/hoge.sh"),
		Parse("* * 4 * thu /tmp/hoge.sh"),
		Parse("* * 5 * wed /tmp/hoge.sh"),
	}
	condition := time.Date(2017, 10, 4, 18, 15, 0, 0, time.UTC)

	for _, expected := range jobs {
		contains, actual := expected.match(condition)
		if !contains {
			t.Fatalf("This job should be match: %+v \n", expected)
		}
		if !reflect.DeepEqual(&expected, actual) {
			t.Fatalf("expected: %+v but actual: %+v\n", expected, actual)
		}
	}
}

func TestMatch_sunday(t *testing.T) {
	jobs := []CronJob{
		Parse("* * 1 * sun /tmp/hoge.sh"),
		Parse("* * 1 * 0 /tmp/hoge.sh"),
		Parse("* * 1 * 7 /tmp/hoge.sh"),
		Parse("* * 1 * fri-sun /tmp/hoge.sh"),
		Parse("* * 1 * 5,7 /tmp/hoge.sh"),
	}
	condition := time.Date(2017, 10, 8, 18, 15, 0, 0, time.UTC)

	for _, expected := range jobs {
		contains, actual := expected.match(condition)
		if !contains {
			t.Fatalf("This job should be match: %+v \n", expected)
		}
		if !reflect.DeepEqual(&expected, actual) {
			t.Fatalf("expected: %+v but actual: %+v\n", expected, actual)
		}
	}
}

func TestUnMatch(t *testing.T) {
	jobs := []CronJob{
		Parse("1 19 4 10 * /tmp/hoge.sh"),
		Parse("* * 3 * thu /tmp/hoge.sh"),
	}
	condition := time.Date(2017, 10, 4, 19, 0, 0, 0, time.UTC)

	for _, expected := range jobs {
		contains, actual := expected.match(condition)
		if contains {
			t.Fatalf("This job should be unmatch: %+v \n", expected)
		}
		if !reflect.DeepEqual(&expected, actual) {
			t.Fatalf("expected: %+v but actual: %+v\n", expected, actual)
		}
	}
}
