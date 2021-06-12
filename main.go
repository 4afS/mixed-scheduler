package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/4afs/mixed-scheduler/file/schedule"
	"github.com/4afs/mixed-scheduler/presenter"
)

var (
	nowUsage  = "display schedule based on the current time"
	onUsage   = "display schedule based on given time"
	onExample = "14:20"
)

func usage() {
	fmt.Printf(`
Usage:
	-now
		%v
	-on
		%v
		ex) %v
`, nowUsage, onUsage, onExample)
	os.Exit(0)
}

func scheduleFilePath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		presenter.PrintErr(fmt.Errorf("user home directory not found"))
	}

	return fmt.Sprintf("%v/.mxs/schedule.yaml", home)
}

func main() {
	var (
		nowF = flag.Bool("now", false, nowUsage)
		onF  = flag.String("on", "", onUsage)
	)

	configPath := scheduleFilePath()

	flag.Usage = usage
	flag.Parse()

	if *nowF && len(*onF) > 0 {
		presenter.PrintErr(fmt.Errorf("select either `now` or `on`"))
	}

	loaded, err := schedule.LoadScheduleFile(configPath)
	if err != nil {
		presenter.PrintErr(fmt.Errorf(
			`schedule file not found
put the schedule file to ~/.mxs/schedule.yaml`))
	}

	s, err := schedule.Parse(loaded)
	if err != nil {
		presenter.PrintErr(fmt.Errorf("invalid schedule file format"))
	}

	now := time.Now()

	var (
		date time.Time
		base string
	)

	switch {
	case *nowF:
		base = fmt.Sprintf("%02v:%02v", now.Hour(), now.Minute())
		date = now

	case len(*onF) > 0:
		timeRegex := regexp.MustCompile("^([01]?[0-9]|2[0-3]):([0-5]?[0-9])$")
		if !timeRegex.MatchString(*onF) {
			presenter.PrintErr(fmt.Errorf("invalid time format\nex) 14:20"))
		}

		splited := strings.Split(*onF, ":")
		h, _ := strconv.Atoi(splited[0])
		m, _ := strconv.Atoi(splited[1])
		date = time.Date(now.Year(), now.Month(), now.Day(), h, m, 0, 0, time.Now().Location())
		base = fmt.Sprintf("%02v:%02v", date.Hour(), date.Minute())
	default:
		flag.Usage()
	}

	fmt.Printf("base: %v\n", base)

	b, plans := s.ToModel(now)
	for _, plan := range plans {
		added := plan.AddDiffBetweenBaseAndGiven(b, date)
		presenter.PrintPlans(added)
	}
}
