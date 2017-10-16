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
	//app.UsageText = "hogefugapiyo"
	app.Version = "0.0.1"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "line, l",
			Usage: "crontab one line",
		},
		cli.StringFlag{
			Name:  "from, fr",
			Usage: "crontab schedule search. format should be 2006-01-02-15-04",
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
		fmt.Println(timeCondition)
		fmt.Println("==============================================")

		var jobs []CronJob
		// TODO ファイル名もJobも未指定の場合、エラー
		targetJob := c.String("line")
		if targetJob != "" {
			jobs = []CronJob{Parse(targetJob)}
		} else {
			fileName := c.Args().Get(0)
			jobs = readCrontabFile(fileName)
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
