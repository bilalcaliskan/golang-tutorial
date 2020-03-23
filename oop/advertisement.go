package oop

type advertisement struct {
	adName string
	cpc int
	noOfClicks int
}

func (a advertisement) calculate() int {
	return a.cpc * a.noOfClicks
}

func (a advertisement) source() string {
	return a.adName
}