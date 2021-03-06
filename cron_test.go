package main

import (
	"reflect"
	"testing"
	"time"
)

func parseForTest(line string) CronJob {
	job, _ := Parse(line)
	return job
}

func TestMatch(t *testing.T) {
	jobs := []CronJob{
		parseForTest("15 18 * * * /tmp/hoge.sh"),
		parseForTest("*/5 * * * * /tmp/hoge.sh"),
		parseForTest("* * 4 * thu /tmp/hoge.sh"),
		parseForTest("* * 5 * wed /tmp/hoge.sh"),
		parseForTest("15 18 3-5 * * /tmp/fuga.sh"),
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

func TestMatch_schedule(t *testing.T) {
	jobs := []CronJob{
		parseForTest("15 18 * * * /tmp/hoge.sh"),
		parseForTest("*/5 * * * * /tmp/hoge.sh"),
		parseForTest("* * 4 * thu /tmp/hoge.sh"),
		parseForTest("* * 5 * wed /tmp/hoge.sh"),
		parseForTest("15 18 3-5 * * /tmp/fuga.sh"),
	}
	condition := time.Date(2017, 10, 4, 18, 15, 0, 0, time.UTC)

	for _, expected := range jobs {
		contains, actual := expected.match(condition)
		if !contains {
			t.Fatalf("This job should be match: %+v \n", expected)
		}
		if actual.schedule != condition {
			t.Fatalf("expected schedule: %+v but actual schedule: %+v\n", condition, actual.schedule)
		}
	}
}

func TestMatch_sunday(t *testing.T) {
	jobs := []CronJob{
		parseForTest("* * 1 * sun /tmp/hoge.sh"),
		parseForTest("* * 1 * 0 /tmp/hoge.sh"),
		parseForTest("* * 1 * 7 /tmp/hoge.sh"),
		parseForTest("* * 1 * fri-sun /tmp/hoge.sh"),
		parseForTest("* * 1 * 5,7 /tmp/hoge.sh"),
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
		parseForTest("1 19 4 10 * /tmp/hoge.sh"),
		parseForTest("* * 3 * thu /tmp/hoge.sh"),
		parseForTest("15 6 10-12 * * /tmp/fuga.sh"),
		parseForTest("0 19 3 * * /tmp/fuga.sh"),
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
