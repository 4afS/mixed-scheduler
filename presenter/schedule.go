package presenter

import (
	"fmt"

	"github.com/4afs/mixed-scheduler/model"
)

func PrintPlans(p model.Plan) {
	h := p.StartAt.Hour()
	m := p.StartAt.Minute()
	fmt.Printf("%02v:%02v - %v\n", h, m, p.Title)
}
