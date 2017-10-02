package main

import (
	"bufio"
	"fmt"
	"os"
	"time"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "rubusy"
	app.Usage = "tell you which cron will be executed."
	app.Action = func(c *cli.Context) error {
		// fmt.Printf("Hello %s\n", c.Args().Get(0))
		fmt.Println(targetTime(time.Now()))

		return nil
	}

	app.Run(os.Args)
}

func targetTime(t time.Time) string {
	return t.Format("target time: 2006-01-02 15:04:05")
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
