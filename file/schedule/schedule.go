package schedule

import (
	"io/ioutil"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/4afs/mixed-scheduler/model"
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
	Title   string `yaml:"title" validate:"required"`
}

func (s *Schedule) validate() error {
	validate := validator.New()
	if err := validate.RegisterValidation("time", validateTime); err != nil {
		return err
	}

	if err := validate.Struct(s); err != nil {
		return err
	}

	return nil
}

func validateTime(fl validator.FieldLevel) bool {
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

func (schedule Schedule) ToModel(now time.Time) (model.Base, []model.Plan) {
	return schedule.toBaseModel(now), schedule.toPlanModels(now)
}

func (schedule Schedule) toBaseModel(now time.Time) model.Base {
	h, m := getTime(schedule.Base)

	base := model.Base{
		Time: todayWithTime(now, h, m),
	}

	return base
}

func (schedule Schedule) toPlanModels(now time.Time) []model.Plan {
	hasChanged := false

	var plans []model.Plan

	for _, p := range schedule.Plan {
		h, m := getTime(p.StartAt)

		date := todayWithTime(now, h, m)
		if hasChanged {
			date = date.Add(24 * time.Hour)
		}

		plans = append(plans, model.Plan{
			StartAt: date,
			Title:   p.Title,
		})
	}

	return plans
}

func getTime(time string) (int, int) {
	splited := strings.Split(time, ":")

	h, _ := strconv.Atoi(splited[0])
	m, _ := strconv.Atoi(splited[1])
	return h, m
}

func todayWithTime(now time.Time, h int, m int) time.Time {
	return time.Date(now.Year(), now.Month(), now.Day(), h, m, 0, 0, time.Now().Location())
}
