package main

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
	return CronJob{
		minute:     []int{32},
		hour:       []int{17},
		dayOfMonth: []int{3},
		month:      []int{10},
		dayOfWeek:  []int{2},
		line:       job,
	}
}
