package main

import "testing"
import "reflect"

func TestNewCronJob(t *testing.T) {
	jobs := []CronJob{
		{line: "32 17 3 10 2 /tmp/hoge.sh", minute: []int{32}, hour: []int{17}, dayOfMonth: []int{3}, month: []int{10}, dayOfWeek: []int{2}},
		{line: "* * * * * /tmp/hoge.sh", minute: minutesRange.all, hour: hourRange.all, dayOfMonth: dayOfMonthRange.all, month: monthRange.all, dayOfWeek: dayOfWeekRange.all},
		{line: "4,7,9 17,23 3,5 10,12 2,6 /tmp/hoge.sh", minute: []int{4, 7, 9}, hour: []int{17, 23}, dayOfMonth: []int{3, 5}, month: []int{10, 12}, dayOfWeek: []int{2, 6}},
		{line: "3-6 12-14 3-5 10-12 2-4 /tmp/hoge.sh", minute: []int{3, 4, 5, 6}, hour: []int{12, 13, 14}, dayOfMonth: []int{3, 4, 5}, month: []int{10, 11, 12}, dayOfWeek: []int{2, 3, 4}},
		{line: "3-6/2 17-20/2 3-5/2 10-12/2 1-3/2 /tmp/hoge.sh", minute: []int{3, 5}, hour: []int{17, 19}, dayOfMonth: []int{3, 5}, month: []int{10, 12}, dayOfWeek: []int{1, 3}},
		{line: "3-6/2 17-20/2 3-5/2 10-12/2 0/3 /tmp/hoge.sh", minute: []int{3, 5}, hour: []int{17, 19}, dayOfMonth: []int{3, 5}, month: []int{10, 12}, dayOfWeek: []int{0, 3, 6}},
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
