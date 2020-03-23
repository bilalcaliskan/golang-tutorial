package oop

type fixedBilling struct {
	projectName string
	biddedAmount int
}

func (fb fixedBilling) calculate() int {
	return fb.biddedAmount
}

func (fb fixedBilling) source() string {
	return fb.projectName
}