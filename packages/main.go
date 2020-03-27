package packages

import (
	"fmt"
	"golang-tutorial/packages/rectangle"
	"log"
)

/*
1. package variables
*/
var pkgRectLen, pkgRectWidth float64 = 6, 7


/*
*2. init function to check if length and width are greater than zero
 */
func init() {
	println("main package initialized")
	if pkgRectLen < 0 {
		log.Fatal("length is less than zero")
	}
	if pkgRectWidth < 0 {
		log.Fatal("width is less than zero")
	}
}

func RunPackages() {
	fmt.Printf("\nBeginning of what are packages and why are they used?\n")
	/*
	So far we have seen go programs which have only one file which has a main function with a couple of other functions.
	In real world scenarios this approach to writing all source code in a single file will not work. It becomes
	impossible to reuse and maintain code written this way. This is where packages save the day.
	 */

	fmt.Printf("\nBeginning of main function and main package\n")
	/*
	Every executable go application must contain a main function. This function is the entry point for execution. The
	main function should reside in the main package.
	The line of code to specify that a particular source file belongs to a package is package packagename. This should
	be first line of every go source file.
	*/

	fmt.Printf("\nBeginning of what happens when running go install geometry\n")
	/*
	This command searches for a file with a main function inside the geometry folder. In this case it finds geometry.go.
	It then compiles it and generates a binary named geometry(geometry.exe in the case of windows) inside the bin folder
	of the workspace.
	 */

	fmt.Printf("\nBeginning of exported names\n")
	/*
	We capitalised the functions Area and Diagonal in the rectangle package. This has a special meaning in Go. Any
	variable or function which starts with a capital letter are exported names in go. Only exported functions and
	variables can be accessed from other packages. In this case we need to access Area and Diagonal functions from the
	main package. Hence they are capitalised.
	Hence if you want to access a function outside of a package, it should be capitalised.
	 */
	var rectLen, rectWidth float64 = 6, 7
	fmt.Println("Geometrical shape properties")
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("diagonal of the rectangle %.2f\n ",rectangle.Diagonal(rectLen, rectWidth))

	fmt.Printf("\nBeginning of init functions\n")
	/*
	Every package can contain a init function. The init function should not have any return type and should not have any
	parameters. The init function cannot be called explicitly in our source code. The init function looks like below:
	func init() {
	}
	The init function can be used to perform initialisation tasks and can also be used to verify the correctness of the
	program before the execution starts.
	The order of initialisation of a package is as follows:
		- Package level variables are initialised first.
		- init function is called next. A package can have multiple init functions (either in a single file or distributed
		across multiple files) and they are called in the order in which they are presented to the compiler.
	If a package imports other packages, the imported packages are initialised first.
	A package will be initialised only once even if it is imported from multiple packages.
	*/
	fmt.Println("Geometrical shape properties")
	fmt.Printf("area of rectangle %.2f\n", rectangle.Area(rectLen, rectWidth))
	fmt.Printf("diagonal of the rectangle %.2f ",rectangle.Diagonal(rectLen, rectWidth))
	/*
	In the below example; The order of initialisation of the main package is:
		- The imported packages are first initialised. Hence rectangle package is initialised first.
		- Package level variables rectLen and rectWidth are initialised next.
		- init function is called.
		- RunPackages() function is called at last
	*/

	fmt.Printf("\nBeginning of use of blank identifier\n")
	/*
	It is illegal in Go to import a package and not to use it anywhere in the code. The compiler will complain if you do
	so. The reason for this is to avoid bloating of unused packages which will significantly increase the compilation
	time.
	But it is quite common to import packages when the application is under active development and use them somewhere in
	the code later if not now. The _ blank identifier saves us in those situations. The error in the above program can
	be silenced by the following code,
	*/
	var _ = rectangle.Area //error silencer
	/*
	We should keep track of these kind of error silencers and remove them including the imported package at the end of
	application development if the package is not used.
	 */

	/*
	Sometimes we need to import a package just to make sure the initialisation takes place even though we do not need to
	use any function or variable from the package. For example, we might need to ensure that the init function of the
	rectangle package is called even though we do not use that package anywhere in our code. The _ blank identifier can
	be used in this case too as show below.
	 */
}