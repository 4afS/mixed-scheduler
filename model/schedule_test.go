package model

import (
	"reflect"
	"testing"
	"time"
)

func TestPlan_AddDiffBetweenBaseAndGiven(t *testing.T) {
	type fields struct {
		StartAt time.Time
		Title   string
	}
	type args struct {
		base Base
		now  time.Time
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   Plan
	}{
		{
			name: "default",
			fields: fields{
				time.Date(2000, 1, 1, 10, 0, 0, 0, time.Now().Location()),
				"",
			},
			args: args{
				base: Base{
					time.Date(2000, 1, 1, 9, 0, 0, 0, time.Now().Location()),
				},
				now: time.Date(2000, 1, 1, 12, 0, 0, 0, time.Now().Location()),
			},
			want: Plan{time.Date(2000, 1, 1, 13, 0, 0, 0, time.Now().Location()), ""},
		},
		{
			name: "extended next day",
			fields: fields{
				time.Date(2000, 1, 1, 10, 0, 0, 0, time.Now().Location()),
				"",
			},
			args: args{
				base: Base{
					time.Date(2000, 1, 1, 9, 0, 0, 0, time.Now().Location()),
				},
				now: time.Date(2000, 1, 1, 23, 0, 0, 0, time.Now().Location()),
			},
			want: Plan{time.Date(2000, 1, 2, 0, 0, 0, 0, time.Now().Location()), ""},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			plan := Plan{
				StartAt: tt.fields.StartAt,
				Title:   tt.fields.Title,
			}
			if got := plan.AddDiffBetweenBaseAndGiven(tt.args.base, tt.args.now); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Plan.AddDiffBetweenBaseAndGiven() = %v, want %v", got, tt.want)
			}
		})
	}
}
