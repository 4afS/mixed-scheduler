package schedule

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadScheduleFile(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	return string(bytes), err
}

type Schedule struct {
	Base string `yaml:"base"`
	Plan []Plan `yaml:"plan"`
}

type Plan struct {
	StartAt string `yaml:"start"`
	Term    int    `yaml:"term,omitempty"`
	Title   string `yaml:"title"`
}

func Parse(loaded string) (Schedule, error) {
	schedule := Schedule{}

	err := yaml.UnmarshalStrict([]byte(loaded), &schedule)
	if err != nil {
		return Schedule{}, err
	}

	if schedule.Base == "" {
		return Schedule{}, fmt.Errorf("base time is empty")
	}

	return schedule, nil
}
