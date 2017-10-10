package main

import "testing"
import "reflect"

func TestNewCronJob(t *testing.T) {
	jobs := []CronJob{
		{line: "32 17 3 10 Tue /tmp/hoge.sh", minute: []int{32}, hour: []int{17}, dayOfMonth: []int{3}, month: []int{10}, dayOfWeek: []int{2}},
	}

	for _, expected := range jobs {
		actual := NewCronJob(expected.line)

		if !reflect.DeepEqual(actual, expected) {
			t.Fatalf("expected: %+v but actual: %+v\n", expected, actual)
		}
	}
}
