package main

import "testing"
import "reflect"

func TestNewCronJob(t *testing.T) {
	jobs := []CronJob{
		{line: "32 17 3 10 2 /tmp/hoge.sh", minute: []int{32}, hour: []int{17}, dayOfMonth: []int{3}, month: []int{10}, dayOfWeek: []int{2}},
		{line: "* 17 3 10 2 /tmp/hoge.sh", minute: []int{32}, hour: []int{17}, dayOfMonth: []int{3}, month: []int{10}, dayOfWeek: []int{2}},
	}

	for _, expected := range jobs {
		actual := NewCronJob(expected.line)

		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("expected: %+v but actual: %+v\n", expected, actual)
		}
	}
}

func TestCronRange(t *testing.T) {
	r := CronRange{from: 1, to: 31}.parse()

	if len(r) != 31 {
		t.Fatalf("expected: %d but actual: %d\n", 31, len(r))
	}
	for index, actual := range r {
		if expected := index + 1; actual != expected {
			t.Fatalf("expected: %q but actual: %q\n", expected, actual)
		}
	}

}
