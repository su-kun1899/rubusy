package main

import (
	"strconv"
	"time"
)

// CronJob はcrontabの一行を表す
type CronJob struct {
	minute     []int
	hour       []int
	dayOfMonth []int
	month      []int
	dayOfWeek  []int
	line       string
	next       time.Time
}

func (job *CronJob) match(cond time.Time) (bool, *CronJob) {
	// TODO 曜日の対応が別途必要
	monthInt, _ := strconv.Atoi(cond.Format("1"))
	if contains(monthInt, job.month) &&
		contains(cond.Hour(), job.hour) &&
		contains(cond.Day(), job.dayOfMonth) &&
		contains(cond.Minute(), job.minute) {

		return true, job
	}

	return false, job
}

func contains(num int, slice []int) bool {
	for _, target := range slice {
		if num == target {
			return true
		}
	}
	return false
}

var dayOfWeekMap = map[string]int{
	"sun": 0,
	"mon": 1,
	"tue": 2,
	"wed": 3,
	"thu": 4,
	"fri": 5,
	"sat": 6,
}

// var dayOfWeekMap = map[int]string{
// 	0: "sun",
// 	1: "mon",
// 	2: "tue",
// 	3: "wed",
// 	4: "thu",
// 	5: "fri",
// 	6: "sat",
// 	7: "sun",
// }
