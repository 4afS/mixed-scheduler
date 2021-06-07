package schedule

import (
	"reflect"
	"testing"
)

func TestParseSchedule(t *testing.T) {
	tcs := []struct {
		arg        string
		expected   Schedule
		shouldFail bool
	}{
		{
			`
base: 9:00
plan:
  - start: 9:30
    term: 30
    title: eat breakfast
  - start: 10:00
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
						Term:    0,
						Title:   "take a shower",
					},
				},
			},
			false,
		},
		{
			"base: 9:00",
			Schedule{
				"9:00",
				nil,
			},
			false,
		},
		{
			"",
			Schedule{},
			true,
		},
		{
			"b",
			Schedule{},
			true,
		},
	}

	for _, tc := range tcs {
		result, err := Parse([]byte(tc.arg))
		if !tc.shouldFail && err != nil {
			t.Errorf("unexpected error occured: %v\n", err)
			continue
		}

		if tc.shouldFail && err == nil {
			t.Errorf("expected error not occured: %v\n", result)
			continue
		}

		if !reflect.DeepEqual(result, tc.expected) {
			t.Errorf("expected: %v\nbut got: %v\n", tc.expected, result)
			continue
		}
	}
}
