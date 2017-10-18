package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "rubusy"
	app.Usage = "cron job schedule searcher"
	app.UsageText = "rubusy [-l line] [-fr from] [-p] [--line value] [--from value] [--plain] [file]"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "job, j",
			Usage: "specific job search of the input.",
		},
		cli.StringFlag{
			Name:  "from, fr",
			Usage: "specific schedule search from time formatted like 2006-01-02-15-04",
		},
		cli.BoolFlag{
			Name:  "plain, p",
			Usage: "only result is written to standard output.",
		},
	}
	app.Action = func(c *cli.Context) error {
		// 検索範囲
		t := time.Now()
		str := c.String("from")
		if str != "" {
			const layout = "2006-01-02-15-04"
			// 時刻変換に失敗したらエラー
			var err error
			t, err = time.Parse(layout, str)
			if err != nil {
				return cli.NewExitError("error: from format is something wrong", 1)
			}
		}

		timeCondition := searchRange{
			from: t,
			to:   t.Add(time.Duration(24*365) * time.Hour),
		}

		plain := c.Bool("plain")
		if !plain {
			fmt.Println(timeCondition)
			fmt.Println("==============================================")
		}

		var jobs []CronJob
		targetJob := c.String("job")
		if targetJob != "" {
			var err error
			job, err := Parse(targetJob)
			if err != nil {
				return cli.NewExitError("error: cron format is something wrong.", 1)
			}
			jobs = []CronJob{job}

		} else {
			fileName := c.Args().Get(0)
			var err error
			jobs, err = readCrontabFile(fileName)
			if err != nil {
				if err == ErrParseJob {
					return cli.NewExitError("error: cron format is something wrong.", 1)
				}
				return cli.NewExitError("error: read crontab file failed", 1)
			}
		}

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
					fmt.Printf("%v : %s\n", job.schedule.Format("2006/01/02 15:04"), job.line)
					break
				}
				cond = cond.Add(time.Duration(1) * time.Minute)
			}
		}

		return nil
	}

	app.Run(os.Args)
}

// ErrCrontabFile はcrontabファイルが不正だった場合に発生するエラー
var ErrCrontabFile = errors.New("crontab file is something wrong")

// ErrParseJob はcronのフォーマットが不正だった場合に発生するエラー
var ErrParseJob = errors.New("cron format is something wrong")

func readCrontabFile(fileName string) ([]CronJob, error) {
	_, err := os.Stat(fileName)
	if err != nil {
		return nil, ErrCrontabFile
	}
	fp, err := os.Open(fileName)
	if err != nil {
		return nil, ErrCrontabFile
	}
	defer fp.Close()

	jobs := make([]CronJob, 0)
	scanner := bufio.NewScanner(fp)
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" || strings.HasPrefix(line, "#") {
			continue
		}
		job, err := Parse(line)
		if err != nil {
			return nil, ErrParseJob
		}
		jobs = append(jobs, job)
	}

	return jobs, nil
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
