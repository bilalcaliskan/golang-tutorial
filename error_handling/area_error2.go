package error_handling

type areaError2 struct {
	err string
	length float64
	width float64
}

func (e *areaError2) Error() string {
	return e.err
}

func (e *areaError2) lengthNegative() bool {
	return e.length < 0
}

func (e *areaError2) widthNegative() bool {
	return e.width < 0
}