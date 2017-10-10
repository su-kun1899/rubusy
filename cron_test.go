package main

import "testing"
import "reflect"

func TestNewCronJob(t *testing.T) {
	job := "32 17 3 10 Tue /tmp/hoge.sh"
	expected := CronJob{
		minute:     []int{32},
		hour:       []int{17},
		dayOfMonth: []int{3},
		month:      []int{10},
		dayOfWeek:  []int{2},
		line:       job,
	}

	actual := NewCronJob(job)

	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected: %+v but actual: %+v\n", expected, actual)
	}
}
