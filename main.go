package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "rubusy"
	app.Usage = "tell you which cron will be executed."
	app.Action = func(c *cli.Context) error {
		fmt.Printf("Hello %s\n", c.Args().Get(0))
		fmt.Println(newTargetTime(time.Now()))

		return nil
	}

	app.Run(os.Args)
}

// 実行されるcronの検索範囲を保持する構造体
type targetTime struct {
	from time.Time
	to   time.Time
}

func (target targetTime) String() string {
	const format = "2006-01-02 15:04:05"
	return fmt.Sprintf("from: %s, to: %s", target.from.Format(format), target.to.Format(format))
}

func newTargetTime(t time.Time) targetTime {
	const format = "2006-01-02 15:04:05"
	from := t
	to := t.Add(time.Duration(1) * time.Hour)
	// return fmt.Sprintf("target time from: %q to: %q", from.Format(format), to.Format(format))
	return targetTime{from, to}
}

// crontab一行分の情報を保持する構造体
type cronTask struct {
	minute     string
	hour       string
	dayOfMonth string
	month      string
	dayOfWeek  string
	line       string
}

func newCronTask(line string) cronTask {
	splited := strings.Split(line, " ")
	return cronTask{
		minute:     splited[0],
		hour:       splited[1],
		dayOfMonth: splited[2],
		month:      splited[3],
		dayOfWeek:  splited[4],
		line:       line,
	}
}

func searchCronTasks(fileName string, t *targetTime) []cronTask {
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	tasks := make([]cronTask, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		task := newCronTask(scanner.Text())
		if ok, _ := filterCronTask(&task, t); ok {
			tasks = append(tasks, task)
		}
	}

	return tasks
}

func filterCronTask(task *cronTask, target *targetTime) (bool, *cronTask) {
	// month
	if task.month != "*" {
		intMonth, _ := strconv.Atoi(task.month)
		numFrom, _ := strconv.Atoi(target.from.Format("1"))
		numTo, _ := strconv.Atoi(target.to.Format("1"))
		if !(numFrom <= intMonth && intMonth <= numTo) {
			return false, nil
		}
	}

	// day of month
	if task.dayOfMonth != "*" {
		intDayOfMonth, _ := strconv.Atoi(task.dayOfMonth)
		//numFrom, _ := strconv.Atoi(target.from.Format("1"))
		//numTo, _ := strconv.Atoi(target.to.Format("1"))
		if !(target.from.Day() <= intDayOfMonth && intDayOfMonth <= target.to.Day()) {
			return false, nil
		}
	}

	return true, task
}
