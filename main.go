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
		// TODO validation的なこと
		fileName := c.Args().Get(0)

		// 検索範囲
		t := time.Now()
		timeCondition := searchRange{
			from: t,
			to:   t.Add(time.Duration(24*365) * time.Hour),
		}
		fmt.Println(timeCondition)
		fmt.Println("==============================================")

		jobs := readCrontabFile(fileName)
		if len(jobs) == 0 {
			fmt.Println("Probably nothing to do. no cron tasks.")
			return nil
		}

		for _, job := range jobs {
			from := timeCondition.from
			to := timeCondition.to
			cond := from
			for cond.Before(to) {
				ok, job := job.match(cond)
				if ok {
					job.next = cond
					fmt.Printf("%v : %s\n", job.next.Format("2006/1/2 15:04"), job.line)
					break
				}
				cond = cond.Add(time.Duration(1) * time.Minute)
			}
		}

		return nil
	}

	app.Run(os.Args)
}

func readCrontabFile(fileName string) []CronJob {
	// TODO ファイルの存在チェック
	fp, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer fp.Close()

	jobs := make([]CronJob, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		jobs = append(jobs, Parse(line))
	}

	return jobs
}

// 実行されるcronの検索範囲を保持する構造体
type searchRange struct {
	from time.Time
	to   time.Time
}

func (s searchRange) String() string {
	const format = "2006/01/02 15:04"
	return fmt.Sprintf("from: %s - to: %s", s.from.Format(format), s.to.Format(format))
}
