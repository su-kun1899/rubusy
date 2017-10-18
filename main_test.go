package main

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestReadCrontabFile(t *testing.T) {
	// setup
	dir, _ := os.Getwd()
	filename := dir + "/crontab_test"
	expected := []CronJob{
		parseForTest("15 9 * * * /tmp/hoge.sh"),
		parseForTest("10 20 * * * /tmp/fuga.sh"),
		parseForTest("10 10 * 12 * /tmp/fuga.sh"),
		parseForTest("10 10 3 10 * /tmp/fuga.sh"),
	}

	// execute
	actual, _ := readCrontabFile(filename)

	// assert
	if !reflect.DeepEqual(actual, expected) {
		t.Fatalf("expected: %v but actual: %v\n", expected, actual)
	}
}

func TestReadCrontabFile_file_error(t *testing.T) {
	// setup
	filename := "un_existing_file_name"

	// execute
	_, err := readCrontabFile(filename)

	// assert
	if err != ErrCrontabFile {
		t.Error("ファイルが存在しない場合、ErrCrontabFileを返すべきです")
	}
}

func TestReadCrontabFile_parse_error(t *testing.T) {
	// setup
	dir, _ := os.Getwd()
	filename := dir + "/crontab_error"

	// execute
	_, err := readCrontabFile(filename)

	// assert
	if err != ErrParseJob {
		t.Error("ファイルにParseできない行が含まれている場合、ErrParseJobを返すべきです")
	}
}

func TestSearchRange(t *testing.T) {
	target := searchRange{
		time.Date(1985, 1, 7, 7, 7, 5, 0, time.UTC),
		time.Date(2017, 10, 3, 17, 18, 9, 0, time.UTC),
	}
	expected := "from: 1985/01/07 07:07 - to: 2017/10/03 17:18"
	actual := target.String()

	if actual != expected {
		t.Fatalf("expected: %q but actual: %q\n", expected, actual)
	}
}
