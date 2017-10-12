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

var minutesRange = newCronRange(0, 59)
var hourRange = newCronRange(0, 23)
var dayOfMonthRange = newCronRange(1, 31)
var monthRange = newCronRange(1, 12)
var dayOfWeekRange = newCronRange(0, 7)

// Parse creates CronJob from crontab line
func Parse(job string) CronJob {
	splited := strings.Split(job, " ")
	minuteBlock := splited[0]
	hourBlock := splited[1]
	dayOfMonthBlock := splited[2]
	monthBlock := splited[3]
	dayOfWeekBlock := splited[4]

	// TODO 曜日は文字列表現の場合がある mon,tue,etc..

	return CronJob{
		minute:     parseBlock(minuteBlock, minutesRange),
		hour:       parseBlock(hourBlock, hourRange),
		dayOfMonth: parseBlock(dayOfMonthBlock, dayOfMonthRange),
		month:      parseBlock(monthBlock, monthRange),
		dayOfWeek:  parseBlock(dayOfWeekBlock, dayOfWeekRange),
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

func parseBlock(block string, maxRange CronRange) []int {
	var cycle int
	if strings.Contains(block, "/") {
		splited := strings.Split(block, "/")
		block = splited[0]
		cycle, _ = strconv.Atoi(splited[1])
	}

	var blockRange []int
	if block == "*" {
		blockRange = maxRange.all

	} else if strings.Contains(block, ",") {
		splited := strings.Split(block, ",")
		blockRange = make([]int, 0, len(splited))
		for _, s := range splited {
			item, _ := strconv.Atoi(s)
			blockRange = append(blockRange, item)
		}

	} else if strings.Contains(block, "-") {
		splited := strings.Split(block, "-")
		start, _ := strconv.Atoi(splited[0])
		end, _ := strconv.Atoi(splited[1])
		blockRange = newCronRange(start, end).all

	} else {
		item, _ := strconv.Atoi(block)
		blockRange = []int{item}
	}

	if cycle == 0 {
		return blockRange
	}

	// `/`を含んでいる場合は、skipする要素を除外して詰め直し
	if len(blockRange) == 1 {
		// `/` の前に指定されてるのが1要素の場合、始点としてあつかう
		blockRange = newCronRange(blockRange[0], maxRange.to).all
	}
	cycles := make([]int, 0)
	for index, item := range blockRange {
		if index%cycle == 0 {
			cycles = append(cycles, item)
		}
	}

	return cycles
}
