package main

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestTargetTime_1985_01_07_07_07_00(t *testing.T) {
	// setup
	tm := time.Date(1985, 1, 7, 7, 7, 0, 0, time.UTC)
	expected := "target time from: \"1985-01-07 07:07:00\" to: \"1985-01-07 08:07:00\""

	// execute
	actual := targetTime(tm)

	// assert
	if actual != expected {
		t.Fatalf("expected: %q actual: %q\n", expected, actual)
	}
}

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
