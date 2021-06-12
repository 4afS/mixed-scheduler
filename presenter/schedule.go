package presenter

import (
	"fmt"

	"github.com/4afs/mixed-scheduler/model"
)

func Show(p model.Plan) {
	h := p.StartAt.Hour()
	m := p.StartAt.Minute()
	fmt.Printf("%v:%v - %v\n", h, m, p.Title)
}
