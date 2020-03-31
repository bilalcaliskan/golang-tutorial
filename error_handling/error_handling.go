package error_handling

import (
	"errors"
	"fmt"
	"math"
	"net"
	"os"
	"path/filepath"
)

func circleArea(radius float64) (float64, error) {
	if radius < 0 {
		return 0, errors.New("Area calculation failed, radius is less than zero")
	}
	return math.Pi * radius * radius, nil
}

func circleAreaWithErrorf(radius float64) (float64, error) {
	if radius < 0 {
		// fmt.Errorf returns error
		return 0, fmt.Errorf("Area calculation failed, radius %0.2f is less than zero", radius)
	}
	return math.Pi * radius * radius, nil
}

func circleAreaWithCustomError(radius float64) (float64, error) {
	if radius < 0 {
		return 0, &areaError{"radius is negative", radius}
	}
	return math.Pi * radius * radius, nil
}

func rectArea(length, width float64) (float64, error) {
	err := ""
	if length < 0 {
		err += "length is less than zero"
	}
	if width < 0 {
		if err == "" {
			err = "width is less than zero"
		} else {
			err += ", width is less than zero"
		}
	}

	if err != "" {
		return 0, &areaError2{err, length, width}
	}
	return length * width, nil
}

func RunErrorHandling() {
	fmt.Printf("\nBeginning of errors...\n")
	/*
	Errors in Go are plain old values. Errors are represented using the built-in error type.
	If a function or method returns an error, then by convention it has to be the last value returned from the function.
	The idiomatic way of handling error in Go is to compare the returned error to nil. A nil value indicates that no
	error has occurred and a non nil value indicates the presence of an error.
	 */
	file, err := os.Open("/test.txt")
	if err != nil {
		// "Error() string" method of error interface is just a string, so we can print it
		fmt.Println(err)
	} else {
		fmt.Println(file.Name(), "opened successfully")
	}

	fmt.Printf("\nBeginning of error type representation...\n")
	/*
	error is an interface type with the following definition:
		type error interface {
	    	Error() string
		}
	It contains a single method with signature Error() string. Any type which implements this interface can be used as
	an error. This method provides the description of the error.
	When printing the error, fmt.Println function calls the Error() string method internally to get the description of
	the error.
	 */

	fmt.Printf("\nBeginning of extracting more information from errors method 1...\n")
	/*
	Asserting the underlying struct type and getting more information from the struct fields.
	If you read the documentation of the Open function carefully, you can see that it returns an error of type *PathError.
	PathError is a struct type.
	*PathError implements the error interface by declaring the Error() string method. This method concatenates the
	operation, path and actual error and returns it.
	The Path field of PathError struct contains the path of the file which caused the error.
	*/
	f, err := os.Open("/test.txt")
	if err, ok := err.(*os.PathError); ok {
		fmt.Printf("type=%T, value=%s, ok=%v\n", err, err, ok)
		fmt.Println("File at path", err.Path, "failed to open")
	} else {
		fmt.Println(f.Name(), "opened successfully")
	}
	/*
	In the above program, we use type assertion to get the underlying value of the error interface. Then
	we print the path using err.Path.
	 */

	fmt.Printf("\nBeginning of extracting more information from errors method 2...\n")
	/*
	The second way to get more information is to assert for the underlying type and get more information by calling
	methods on the struct type.
	In the program below, we are trying to get the ip address of an invalid domain name golangbot123.com. We get the
	underlying value of the error by asserting it to type *net.DNSError. Then we check whether the error is due to
	timeout or temporary respectively.
	 */
	addr, err := net.LookupHost("golangbot123.com")
	if err, ok := err.(*net.DNSError); ok {
		fmt.Printf("type=%T, value=%s, ok=%v\n", err, err, ok)
		if err.Timeout() {
			fmt.Println(err.Name)
			fmt.Println("operation timed out")
		} else if err.Temporary() {
			fmt.Println(err.Name)
			fmt.Println("temporary error")
		} else {
			fmt.Println(err.Name)
			fmt.Println("generic error: ", err)
		}
	} else if err == nil {
		fmt.Println(addr)
	}

	fmt.Printf("\nBeginning of extracting more information from errors method 3...\n")
	/*
	The third way to get more details about an error is the direct comparison with a variable of type error.
	The Glob function of the filepath package is used to return the names of all files that matches a pattern. This
	function returns an error ErrBadPattern when the pattern is malformed.
	ErrBadPattern is defined in the filepath package as follows:
		var ErrBadPattern = errors.New("syntax error in pattern")
	errors.New() is used to create a new error. We will discuss about this in detail in the next tutorial.
	ErrBadPattern is returned by the Glob function when the pattern is malformed.
	 */
	files, err := filepath.Glob("[")
	if err != nil && err == filepath.ErrBadPattern {
		fmt.Println(err)
		fmt.Printf("type=%T, value=%s\n", err, err)
	} else if err == nil {
		fmt.Println("matched files", files)
	}

	fmt.Printf("\nBeginning of ignoring errors...\n")
	/*
	Never ever ignore an error. Ignoring errors is inviting for trouble.
	 */
	files2, _ := filepath.Glob("[")
	fmt.Println("matched files", files2)

	fmt.Printf("\nBeginning of creating custom errors using the New function...\n")
	/*
	The simplest way to create a custom error is to use the New function of the errors package.
	 */
	radius := -20.0
	area, err := circleArea(radius)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Area of circle %0.2f\n", area)
	}

	fmt.Printf("\nBeginning of adding more information to the errors using Errorf...\n")
	/*
	The above program works well but wouldn't it be nice if we print the actual radius which caused the error. This is
	where the Errorf function of the fmt package comes in handy. This function formats the error according to a format
	specifier and returns a string as value that satisfies error.
	 */
	radius = -20.0
	area, err = circleAreaWithErrorf(radius)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Area of circle %0.2f\n", area)
	}

	fmt.Printf("\nBeginning of providing more information about the error using struct type and fields...\n")
	/*
	It is also possible to use struct types which implement the error interface as errors. This gives us more
	flexibility with error handling.
	We will create a struct type that implements the error interface and use its fields to provide more information
	about the error.
	 */
	radius = -20.0
	area, err = circleAreaWithCustomError(radius)
	if err != nil {
		if err, ok := err.(*areaError); ok {
			fmt.Printf("Radius %0.2f is less than zero\n\n", err.radius)
		}
	}

	fmt.Printf("\nBeginning of providing more information about the error using methods on struct types...\n")
	/*
	We have used methods on struct error types to provide more information about the error(check the struct on area_error2.go file).
	Now that we have the error type, lets implement the error interface and add a couple of methods on the error type to
	provide more information about the error.
	we return the description of the error from the Error() string method. The lengthNegative() bool method returns true
	when the length is less than zero and widthNegative() bool method returns true when the width is less than zero.
	These two methods provide more information about the error, in this case they say whether the area calculation failed
	because of the length being negative or width being negative. Thus we have used methods on struct error types to
	provide more information about the error.
	 */
	length, width := -5.0, -9.0
	area, err = rectArea(length, width)
	if err != nil {
		if err, ok := err.(*areaError2); ok {
			if err.lengthNegative() {
				fmt.Printf("error: length %0.2f is less than zero\n", err.length)
			}
			if err.widthNegative() {
				fmt.Printf("error: width %0.2f is less than zero\n", err.width)
			}
			return
		}
	}
	fmt.Println("area of rect", area)
}