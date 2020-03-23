package oop

type timeAndMaterial struct {
	projectName string
	noOfHours int
	hourlyRate int
}

func (tm timeAndMaterial) calculate() int {
	return tm.noOfHours * tm.hourlyRate
}

func (tm timeAndMaterial) source() string {
	return tm.projectName
}