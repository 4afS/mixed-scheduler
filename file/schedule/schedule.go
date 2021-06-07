package schedule

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

func LoadScheduleFile(path string) ([]byte, error) {
	return ioutil.ReadFile(path)
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

func Parse(bytes []byte) (Schedule, error) {
	schedule := Schedule{}

	err := yaml.UnmarshalStrict(bytes, &schedule)
	if err != nil {
		return Schedule{}, err
	}

	if schedule.Base == "" {
		return Schedule{}, fmt.Errorf("base time is empty")
	}

	return schedule, nil
}
