package main

import (
	"fmt"
)

type teams interface {
	getTeamPlan() string
}

//*******
// to struct /s that each one contains a different types.
type teamOne struct {
	someProp float64
	name     string
}

type teamTwo struct {
	someProp int
	name     string
} //*******

func main() {
	t1 := teamOne{}
	t2 := teamTwo{}

	getPlan(t1)
	getPlan(t2)
}

func getPlan(t teams) {
	// no need to DRY, since that the same.
	fmt.Println(t.getTeamPlan())
}

// two receivers functions that is contains a very *custom* logic:
func (t teamOne) getTeamPlan() string {
	// custom logic here
	return "4,2,3,1"
}

// you can get rid of 't' since you aren't use it to extract data from it.
// like so: func (teamTwo)...
func (t teamTwo) getTeamPlan() string {
	// custom logic here
	return "4,4,2,2"
}
