package main

// CronJob はcrontabの一行を表す
type CronJob struct {
	minute     []int
	hour       []int
	dayOfMonth []int
	month      []int
	dayOfWeek  []int
	line       string
}

func (job *CronJob) match(t targetTime) (bool, *CronJob) {
	return true, job
}
