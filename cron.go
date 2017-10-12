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

func monthInt(t time.Time) (int, error) {
	monthInt, err := strconv.Atoi(t.Format("1"))
	return monthInt, err
}

func (job *CronJob) match(cond time.Time) (bool, *CronJob) {
	// TODO 曜日の対応が別途必要
	monthInt, _ := monthInt(cond)
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
