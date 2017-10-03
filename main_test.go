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
	expected := []string{"15 9 * * * /tmp/hoge.sh", "10 10 * * * /tmp/fuga.sh"}

	// execute
	actual := readCrontabFile(filename)

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

// 10を含んでいる場合通る
// * の場合通る
// カンマ区切りの場合通る
func TestFilterCronTask_Month(t *testing.T) {
	// cronTask := cronTask{
	// 	minute:     "32",
	// 	hour:       "17",
	// 	dayOfMonth: "3",
	// 	month:      "10",
	// 	dayOfWeek:  "Tue",
	// 	line:       "32 17 3 10 Tue /tmp/hoge.sh",
	// }
}
