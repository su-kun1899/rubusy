package main

import (
	"testing"
	"time"
)

func TestTargetTime_1985_01_07_07_07_00(t *testing.T) {
	// setup
	tm := time.Date(1985, 1, 7, 7, 7, 0, 0, time.UTC)
	expected := "target time: 1985-01-07 07:07:00"

	// execute
	actual := targetTime(tm)

	// assert
	if actual != expected {
		t.Fatalf("actual is differet from expected. expected: %q actual: %q\n", expected, actual)
	}
}
