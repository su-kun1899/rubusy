package main

import (
	"strconv"
	"time"
)

// CronJob はcrontabの一行を表す
type CronJob struct {
	minute       []int
	hour         []int
	dayOfMonth   []int
	month        []int
	dayOfWeek    []int
	line         string
	schedule     time.Time
	dayOfWeekFlg bool
}

func (job *CronJob) match(cond time.Time) (bool, *CronJob) {
	// 月をintに変換
	monthInt, _ := strconv.Atoi(cond.Format("1"))

	// 曜日をintに変換
	weekday := cond.Format("Mon")
	weekdayInt := weekDayMap[weekday]
	weekDayMatch := contains(weekdayInt, job.dayOfWeek) || weekday == "Sun"
	if !weekDayMatch && weekday == "Sun" {
		// 日曜日は7の場合もある
		weekdayInt = 7
		weekDayMatch = contains(weekdayInt, job.dayOfWeek)
	}

	if contains(monthInt, job.month) &&
		contains(cond.Hour(), job.hour) &&
		// 日付と曜日が両方指定されている場合は、どちらかが満たされた場合両方でコマンドが実行される
		(contains(cond.Day(), job.dayOfMonth) || (job.dayOfWeekFlg && weekDayMatch)) &&
		contains(cond.Minute(), job.minute) {
		job.schedule = cond

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

// dayOfWeekMap はcrontabの曜日をintに変換する用
var dayOfWeekMap = map[string]int{
	"sun": 0,
	"mon": 1,
	"tue": 2,
	"wed": 3,
	"thu": 4,
	"fri": 5,
	"sat": 6,
}

// weekDayMap はtime.Dateの曜日をintに変換する用
var weekDayMap = map[string]int{
	"Sun": 0,
	"Mon": 1,
	"Tue": 2,
	"Wed": 3,
	"Thu": 4,
	"Fri": 5,
	"Sat": 6,
}
