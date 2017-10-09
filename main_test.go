package main

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestReadCrontabFile_crontab_example(t *testing.T) {
	// setup
	dir, _ := os.Getwd()
	filename := dir + "/crontab_example"
	expected := []cronTask{
		newCronTask("15 9 * * * /tmp/hoge.sh"),
		newCronTask("10 10 * * * /tmp/fuga.sh"),
	}

	targetTime := targetTime{
		time.Date(2017, 10, 4, 18, 0, 0, 0, time.UTC),
		time.Date(2017, 10, 4, 19, 0, 0, 0, time.UTC),
	}

	// execute
	actual := searchCronTasks(filename, &targetTime)

	// assert
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected: %q but actual: %q\n", expected, actual)
	}
}

func TestNewTargetTime_1985_01_07_07_07_00(t *testing.T) {
	// setup
	param := time.Date(1985, 1, 7, 7, 7, 0, 0, time.UTC)
	expectedFrom := param
	expectedTo := time.Date(1985, 1, 7, 8, 7, 0, 0, time.UTC)
	expected := targetTime{expectedFrom, expectedTo}

	// execute
	actual := newTargetTime(param)

	// asert
	if actual != expected {
		t.Fatalf("expected: %q but actual: %q\n", expected, actual)
	}
}

func TestStringTargetTime(t *testing.T) {
	target := targetTime{
		time.Date(1985, 1, 7, 7, 7, 5, 0, time.UTC),
		time.Date(2017, 10, 3, 17, 18, 9, 0, time.UTC),
	}
	expected := "from: 1985-01-07 07:07:05, to: 2017-10-03 17:18:09"
	actual := target.String()

	if actual != expected {
		t.Fatalf("expected: %q but actual: %q\n", expected, actual)
	}
}
func TestNewCrontab(t *testing.T) {
	crontabLine := "32 17 3 10 Tue /tmp/hoge.sh"
	expected := cronTask{
		minute:     "32",
		hour:       "17",
		dayOfMonth: "3",
		month:      "10",
		dayOfWeek:  "Tue",
		line:       crontabLine,
	}

	actual := newCronTask(crontabLine)

	if actual != expected {
		t.Fatalf("expected: %q but actual: %q\n", expected, actual)
	}
}

func TestSearchCronTasks_unmatch(t *testing.T) {
	// given
	dir, _ := os.Getwd()
	testFile := dir + "/crontab_unmatch"
	targetTime := targetTime{
		time.Date(2017, 10, 4, 18, 0, 0, 0, time.UTC),
		time.Date(2017, 10, 4, 19, 0, 0, 0, time.UTC),
	}

	// when
	actual := searchCronTasks(testFile, &targetTime)

	if len(actual) != 0 {
		t.Fatal("actual should return empty slice \n")
	}
}

func TestFilterCronTask_match(t *testing.T) {
	targetTime := targetTime{
		from: time.Date(2017, 10, 4, 18, 0, 0, 0, time.UTC),
		to:   time.Date(2017, 10, 4, 19, 0, 0, 0, time.UTC),
	}
	cronTasks := []cronTask{
		newCronTask("30 * * * * /tmp/hoge.sh"),
		newCronTask("30 * * 10 * /tmp/hoge.sh"),
		newCronTask("30 * * 8,9,10 * /tmp/hoge.sh"),
		newCronTask("30 * * 9,10,11 * /tmp/hoge.sh"),
		newCronTask("30 * * 10,11,12 * /tmp/hoge.sh"),
		newCronTask("30 * * 8-10 * /tmp/hoge.sh"),
		newCronTask("30 * * 9-11 * /tmp/hoge.sh"),
		newCronTask("30 * * 10-12 * /tmp/hoge.sh"),
		newCronTask("30 * * 10/2 * /tmp/hoge.sh"),
		newCronTask("30 * * 4/2 * /tmp/hoge.sh"),
		newCronTask("30 * * 6/2 * /tmp/hoge.sh"),
		newCronTask("30 * * 1/3 * /tmp/hoge.sh"),
		newCronTask("30 * * */3 * /tmp/hoge.sh"),
	}

	for _, task := range cronTasks {
		ok, actual := filterCronTask(&task, &targetTime)
		if !ok {
			t.Fatalf("crontTask should match filter condition: %q \n", task)
		}
		if actual != &task {
			t.Fatalf("actual should match cronTask: %q \n", task)
		}
	}
}

func TestFilterCronTask_unmatch(t *testing.T) {
	targetTime := targetTime{
		from: time.Date(2017, 10, 4, 18, 0, 0, 0, time.UTC),
		to:   time.Date(2017, 10, 4, 19, 0, 0, 0, time.UTC),
	}
	cronTasks := []cronTask{
		newCronTask("30 * * 11 * /tmp/hoge.sh"),
		newCronTask("30 * * 8,9,12 * /tmp/hoge.sh"),
		newCronTask("30 * * 1-9 * /tmp/hoge.sh"),
		newCronTask("30 * * 9/2 * /tmp/hoge.sh"),
		newCronTask("30 * * 3/2 * /tmp/hoge.sh"),
		newCronTask("30 * * 2/3 * /tmp/hoge.sh"),
		newCronTask("30 * * */2 * /tmp/hoge.sh"),
		newCronTask("30 * * 1-9/3 * /tmp/hoge.sh"),
	}

	for _, task := range cronTasks {
		ok, actual := filterCronTask(&task, &targetTime)
		if ok {
			t.Fatalf("crontTask should not match filter condition: %q \n", task)
		}
		if actual != nil {
			t.Fatal("actual should return nil \n")
		}
	}
}
