package main

import (
	"strconv"
	"strings"
)

type CronJob struct {
	minute     []int
	hour       []int
	dayOfMonth []int
	month      []int
	dayOfWeek  []int
	line       string
}

// NewCronJob creates CronJob from crontab line
func NewCronJob(job string) CronJob {
	splited := strings.Split(job, " ")
	minuteBlock := splited[0]
	// hourBlock := splited[1]
	// dayOfMonthBlock := splited[2]
	// monthBlock := splited[3]
	// dayOfWeekBlock := splited[4]

	var minuteRange []int
	if strings.Contains(minuteBlock, "*") {
		minuteRange = MinutesRange.all
	} else {
		minute, _ := strconv.Atoi(minuteBlock)
		minuteRange = []int{minute}
	}

	return CronJob{
		minute:     minuteRange,
		hour:       []int{17},
		dayOfMonth: []int{3},
		month:      []int{10},
		dayOfWeek:  []int{2},
		line:       job,
	}
}

type CronRange struct {
	from int
	to   int
	all  []int
}

func newCronRange(from int, to int) CronRange {
	r := CronRange{from: from, to: to}
	all := make([]int, 0, r.to-r.from+1)
	for index := r.from; index <= r.to; index++ {
		all = append(all, index)
	}
	r.all = all
	return r
}

var MinutesRange = newCronRange(0, 59)
