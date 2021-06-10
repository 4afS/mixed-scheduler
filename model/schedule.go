package model

import "time"

type Base struct {
	Time time.Time
}

type Plan struct {
	StartAt time.Time
	Title   string
}

func (plan Plan) AddDiffBetweenBaseAndGiven(base Base, now time.Time) Plan {
	diff := now.Sub(base.Time)

	return Plan{
		StartAt: plan.StartAt.Add(diff),
		Title:   plan.Title,
	}
}
