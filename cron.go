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

type Condition struct {
	from time.Time
	to   time.Time
}

func monthInt(t time.Time) (int, error) {
	monthInt, err := strconv.Atoi(t.Format("1"))
	return monthInt, err
}

func (job *CronJob) match(cond Condition) (bool, *CronJob) {
	t := cond.from
	// TODO 曜日の対応が別途必要
	for t.Before(cond.to) {
		monthInt, _ := monthInt(t)
		if contains(monthInt, job.month) &&
			contains(t.Hour(), job.hour) &&
			contains(t.Day(), job.dayOfMonth) &&
			contains(t.Minute(), job.minute) {

			job.next = t
			return true, job
		}
		t = t.Add(time.Duration(1) * time.Minute)
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
