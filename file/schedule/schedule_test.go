package schedule

import (
	"reflect"
	"testing"
	"time"

	"github.com/4afs/mixed-scheduler/model"
)

func TestParse(t *testing.T) {
	tests := []struct {
		name    string
		arg     string
		want    Schedule
		wantErr bool
	}{
		{
			"default",
			`
base: 9:00
plan:
  - start: 9:30
    title: eat breakfast
  - start: 10:00
    title: take a shower`,
			Schedule{
				Base: "9:00",
				Plan: []Plan{
					{
						StartAt: "9:30",
						Title:   "eat breakfast",
					},
					{
						StartAt: "10:00",
						Title:   "take a shower",
					},
				},
			},
			false,
		},
		{
			"has no title",
			`
base: 9:00
plan:
  - start: 9:30`,
			Schedule{},
			true,
		},
		{
			"only base time",
			"base: 9:00",
			Schedule{
				"9:00",
				nil,
			},
			false,
		},
		{
			"invalid base time format",
			"base: 24:00",
			Schedule{},
			true,
		},
		{
			"invalid start time",
			`
base: 9:00
plan:
  - start: 24:00
    title: eat breakfast`,
			Schedule{},
			true,
		},
		{
			"empty title",
			`
base: 9:00
plan:
  - start: 9:00
    title: `,
			Schedule{},
			true,
		},
		{
			"empty string",
			"",
			Schedule{},
			true,
		},
		{
			"invalid string",
			"b",
			Schedule{},
			true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Parse(tt.arg)

			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, Parse() = %v, wantErr %v", err, got, tt.wantErr)
				return
			}

			if !tt.wantErr && !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parse() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestTodayWithTime(t *testing.T) {
	now := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Now().Location())
	type args struct {
		h int
		m int
	}
	tests := []struct {
		name string
		args args
		want time.Time
	}{
		{
			"default",
			args{
				10,
				11,
			},
			time.Date(2000, 1, 1, 10, 11, 0, 0, time.Now().Location()),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := todayWithTime(now, tt.args.h, tt.args.m); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("todayWithTime() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetTime(t *testing.T) {
	tests := []struct {
		name  string
		arg   string
		want  int
		want1 int
	}{
		{
			"basic 9 o'clock",
			"9:00",
			9,
			0,
		},
		{
			"basic 12 o'clock",
			"12:00",
			12,
			0,
		},
		{
			"zero-fill 9 o'clock",
			"09:00",
			9,
			0,
		},
		{
			"basic 9:01",
			"9:01",
			9,
			1,
		},
		{
			"basic 9:11",
			"9:11",
			9,
			11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getTime(tt.arg)
			if got != tt.want {
				t.Errorf("getTime() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getTime() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestScheduleToBaseModel(t *testing.T) {
	now := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Now().Location())
	type fields struct {
		Base string
		Plan []Plan
	}
	tests := []struct {
		name   string
		fields fields
		want   model.Base
	}{
		{
			"",
			fields{
				Base: "9:10",
			},
			model.Base{
				Time: time.Date(2000, 1, 1, 9, 10, 0, 0, time.Now().Location()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schedule := Schedule{
				Base: tt.fields.Base,
				Plan: tt.fields.Plan,
			}
			got := schedule.toBaseModel(now)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schedule.toBaseModel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScheduleToPlanModels(t *testing.T) {
	now := time.Date(2000, 1, 1, 0, 0, 0, 0, time.Now().Location())
	type fields struct {
		Base string
		Plan []Plan
	}
	tests := []struct {
		name   string
		fields fields
		want   []model.Plan
	}{
		{
			"one plan",
			fields{
				Plan: []Plan{
					{
						StartAt: "10:00",
						Title:   "eat breakfast",
					},
				},
			},
			[]model.Plan{
				{
					StartAt: time.Date(2000, 1, 1, 10, 0, 0, 0, time.Now().Location()),
					Title:   "eat breakfast",
				},
			},
		},
		{
			"plans",
			fields{
				Plan: []Plan{
					{
						StartAt: "10:00",
						Title:   "eat breakfast",
					},
					{
						StartAt: "11:30",
						Title:   "leave home",
					},
				},
			},
			[]model.Plan{
				{
					StartAt: time.Date(2000, 1, 1, 10, 0, 0, 0, time.Now().Location()),
					Title:   "eat breakfast",
				},
				{
					StartAt: time.Date(2000, 1, 1, 11, 30, 0, 0, time.Now().Location()),
					Title:   "leave home",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			schedule := Schedule{
				Base: tt.fields.Base,
				Plan: tt.fields.Plan,
			}
			got := schedule.toPlanModels(now)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Schedule.toPlanModels() = %v, want %v", got, tt.want)
			}
		})
	}
}
