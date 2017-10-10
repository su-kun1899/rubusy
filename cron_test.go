package main

import "testing"
import "reflect"

func TestNewCronJob(t *testing.T) {
	jobs := []CronJob{
		{line: "32 17 3 10 2 /tmp/hoge.sh", minute: []int{32}, hour: []int{17}, dayOfMonth: []int{3}, month: []int{10}, dayOfWeek: []int{2}},
		{line: "* * 3 10 2 /tmp/hoge.sh", minute: MinutesRange.all, hour: HourRange.all, dayOfMonth: []int{3}, month: []int{10}, dayOfWeek: []int{2}},
		{line: "4,7,9 17 3 10 2 /tmp/hoge.sh", minute: []int{4, 7, 9}, hour: []int{17}, dayOfMonth: []int{3}, month: []int{10}, dayOfWeek: []int{2}},
		{line: "3-6 17 3 10 2 /tmp/hoge.sh", minute: []int{3, 4, 5, 6}, hour: []int{17}, dayOfMonth: []int{3}, month: []int{10}, dayOfWeek: []int{2}},
		{line: "3-6/2 17 3 10 2 /tmp/hoge.sh", minute: []int{3, 5}, hour: []int{17}, dayOfMonth: []int{3}, month: []int{10}, dayOfWeek: []int{2}},
	}

	for _, expected := range jobs {
		actual := NewCronJob(expected.line)

		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("expected: %+v but actual: %+v\n", expected, actual)
		}
	}
}

func TestNewCronRange(t *testing.T) {
	r := newCronRange(1, 31)

	if len(r.all) != 31 {
		t.Fatalf("expected: %d but actual: %d\n", 31, len(r.all))
	}
	for index, actual := range r.all {
		if expected := index + 1; actual != expected {
			t.Fatalf("expected: %q but actual: %q\n", expected, actual)
		}
	}

}
