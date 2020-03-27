package conditionals

import "fmt"

func RunConditionals() {
	fmt.Printf("\nBeginning of advanced conditionals...\n")
	/*
	There is one more variant of if which includes a optional statement component which is executed before the condition
	is evaluated. Its syntax is:
		if statement; condition {
		}
	The scope of statement is limited to the if else blocks. If we try to access statement from outside of if else, the
	compiler will complain.
	 */
	if num := 10; num % 2 == 0 {
		fmt.Println(num, "is even")
	} else {
		fmt.Println(num, "is odd")
	}
	/*
	The scope of num is limited to the if else blocks. If we try to access num from outside the if or else, the compiler
	will complain.
	 */

	fmt.Printf("\nBeginning of advanced gotcha...\n")
	/*
	The else statement should start in the same line after the closing curly brace } of the if statement. If not the
	compiler will complain.
	Below if/else block is not allowed:
		if num % 2 == 0 { //checks if number is even
	        fmt.Println("the number is even")
	    }
	    else {
	        fmt.Println("the number is odd")
	    }
	In the program above, the else statement does not start in the same line after the closing } of the if statement.
	Instead it starts in the next line. This is not allowed in Go.
	The reason is because of the way Go inserts semicolons automatically. You can read about the semicolon insertion rule
	here https://golang.org/ref/spec#Semicolons.
	In the rules, it's specified that a semicolon will be inserted after }, if that is the final token of the line. So
	a semicolon is automatically inserted after the if statement's }.
	So our program actually becomes:
		if num%2 == 0 {
	      fmt.Println("the number is even")
		};  //semicolon inserted by Go
		else {
			  fmt.Println("the number is odd")
		}
	Since if{...} else {...} is one single statement, a semicolon should not be present in the middle of it. Hence there
	is a requirement to place the else in the same line after the closing }.
	 */
	num := 10
	if num % 2 == 0 { //checks if number is even
		fmt.Println("the number is even")
	} else {
		fmt.Println("the number is odd")
	}
}