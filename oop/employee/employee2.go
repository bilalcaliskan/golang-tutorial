package employee

import (
	"fmt"
)

type employee2 struct {
	firstName   string
	lastName    string
	totalLeaves int
	leavesTaken int
}

func NewEmployee2(firstName string, lastName string, totalLeave int, leavesTaken int) employee2 {
	e := employee2 {firstName, lastName, totalLeave, leavesTaken}
	return e
}

func (e employee2) LeavesRemaining() {
	fmt.Printf("%s %s has %d leaves remaining", e.firstName, e.lastName, (e.totalLeaves - e.leavesTaken))
}
