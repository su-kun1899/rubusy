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

		// TODO validation的なこと

		// 検索範囲
		t := time.Now()
		str := c.String("from")
		if str != "" {
			const layout = "2006-01-02-15-04"
			// TODO 時刻変換に失敗したらエラー
			t, _ = time.Parse(layout, str)
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
		// TODO ファイル名もJobも未指定の場合、エラー
		targetJob := c.String("line")
		if targetJob != "" {
			jobs = []CronJob{Parse(targetJob)}
		} else {
			fileName := c.Args().Get(0)
			var err error
			jobs, err = readCrontabFile(fileName)
			if err != nil {
				return cli.NewExitError("read crontab file failed", 1)
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
		jobs = append(jobs, Parse(line))
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
