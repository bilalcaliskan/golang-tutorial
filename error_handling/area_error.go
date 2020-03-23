package error_handling

import "fmt"

type areaError struct {
	err string
	radius float64
}

func (e *areaError) Error() string {
	return fmt.Sprintf("radius %0.2f: %s", e.radius, e.err)
}