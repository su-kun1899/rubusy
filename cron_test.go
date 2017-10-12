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

func TestUnMatch(t *testing.T) {
	jobs := []CronJob{
		Parse("1 19 4 10 * /tmp/hoge.sh"),
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
