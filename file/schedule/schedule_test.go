package schedule

import (
	"reflect"
	"testing"
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
    term: 30
    title: eat breakfast
  - start: 10:00
    term: 30
    title: take a shower`,
			Schedule{
				Base: "9:00",
				Plan: []Plan{
					{
						StartAt: "9:30",
						Term:    30,
						Title:   "eat breakfast",
					},
					{
						StartAt: "10:00",
						Term:    30,
						Title:   "take a shower",
					},
				},
			},
			false,
		},
		{
			"has no term",
			`
base: 9:00
plan:
  - start: 9:30
    title: eat breakfast`,
			Schedule{
				Base: "9:00",
				Plan: []Plan{
					{
						StartAt: "9:30",
						Term:    0,
						Title:   "eat breakfast",
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
			"term -1",
			`
base: 9:00
plan:
  - start: 9:00
    term: -1
    title: eat breakfast`,
			Schedule{},
			true,
		},
		{
			"term over 1440",
			`
base: 9:00
plan:
  - start: 9:00
    term: 1441
    title: eat breakfast`,
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
