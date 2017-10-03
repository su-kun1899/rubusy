package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "rubusy"
	app.Usage = "tell you which cron will be executed."
	app.Action = func(c *cli.Context) error {
		// fmt.Printf("Hello %s\n", c.Args().Get(0))
		fmt.Println(newTargetTime(time.Now()))

		return nil
	}

	app.Run(os.Args)
}

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

func readCrontabFile(fileName string) []string {
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	lines := make([]string, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
