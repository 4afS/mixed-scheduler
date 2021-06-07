package schedule

import (
	"io/ioutil"
	"regexp"

	"github.com/go-playground/validator"
	"gopkg.in/yaml.v2"
)

func LoadScheduleFile(path string) (string, error) {
	bytes, err := ioutil.ReadFile(path)
	return string(bytes), err
}

type Schedule struct {
	Base string `yaml:"base" validate:"required,time"`
	Plan []Plan `yaml:"plan" validate:"dive"`
}

type Plan struct {
	StartAt string `yaml:"start" validate:"required,time"`
	Term    int    `yaml:"term" validate:"min=0,max=1440,numeric"`
	Title   string `yaml:"title" validate:"required"`
}

func (s *Schedule) validate() error {
	validate := validator.New()
	if err := validate.RegisterValidation("time", ValidateTime); err != nil {
		return err
	}

	if err := validate.Struct(s); err != nil {
		return err
	}

	return nil
}

func ValidateTime(fl validator.FieldLevel) bool {
	timeRegex := regexp.MustCompile("^([01]?[0-9]|2[0-3]):([0-5]?[0-9])$")
	return timeRegex.MatchString(fl.Field().String())
}

func Parse(loaded string) (Schedule, error) {
	schedule := Schedule{}

	err := yaml.UnmarshalStrict([]byte(loaded), &schedule)
	if err != nil {
		return schedule, err
	}

	if err = schedule.validate(); err != nil {
		return schedule, err
	}

	return schedule, nil
}
