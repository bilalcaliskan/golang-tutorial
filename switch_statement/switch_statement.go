package switch_statement

import "fmt"

func number() int {
	num := 15 * 5
	return num
}

func RunSwitchStatement() {
	fmt.Printf("\nBeginning of switch statement...\n")
	/*
	A switch is a conditional statement which evaluates an expression and compares it against a list of possible matches
	and executes blocks of code according to the match. It can be considered as an idiomatic way of writing multiple if
	else clauses.
	 */
	finger := 4
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	}

	fmt.Printf("\nBeginning of default case...\n")
	/*
	We have only 5 fingers in our hand. What will happen if we input a incorrect finger number. This is where the default
	case comes into picture. The default case will be executed when none of the other cases match.
	 */
	/*
	A switch can include an optional statement which is executed before the expression is evaluated. In this line switch
	finger := 8; finger finger is first declared and also used in the expression. The scope of finger in this case is
	limited to the switch block.
	 */
	switch finger := 8; finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default: //default case
		fmt.Println("incorrect finger number")
	}

	fmt.Printf("\nBeginning of multiple expressions in case...\n")
	/*
	It is possible to include multiple expressions in a case by separating them with comma.
	 */
	letter := "i"
	switch letter {
	case "a", "e", "i", "o", "u": //multiple expressions in case
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

	fmt.Printf("\nBeginning of expressionless switch...\n")
	/*
	The expression in a switch is optional and it can be omitted. If the expression is omitted, the switch is considered
	to be switch true and each of the case expression is evaluated for truth and the corresponding block of code is executed.
	 */
	num := 75
	switch { // expression is omitted
	case num >= 0 && num <= 50:
		fmt.Println("num is greater than 0 and less than 50")
	case num >= 51 && num <= 100:
		fmt.Println("num is greater than 51 and less than 100")
	case num >= 101:
		fmt.Println("num is greater than 100")
	}
	/*
	In the above program the expression is absent in switch and hence it is considered as true and each of the case is
	evaluated.
	 */

	fmt.Printf("\nBeginning of fallthrough...\n")
	/*
	In Go the control comes out of the switch statement immediately after a case is executed. A fallthrough statement is
	used to transfer control to the first statement of the case that is present immediately after the case which has been
	executed.
	 */
	switch num := number(); { //num is not a constant
	case num < 50:
		fmt.Printf("%d is lesser than 50\n", num)
		fallthrough
	case num < 100:
		fmt.Printf("%d is lesser than 100\n", num)
		fallthrough
	case num < 200:
		fmt.Printf("%d is lesser than 200", num)
		// fallthrough // Fallthrough cannot be called on the final case of a switch
	}
	/*
	Switch and case expressions need not be only constants. They can be evaluated at runtime too.
	In the program above num is initialised to the return value of the function number(). The control comes inside the
	switch and the cases are evaluated. case num < 100: is true and the program prints 75 is lesser than 100. The next
	statement is fallthrough. When fallthrough is encountered the control moves to the first statement of the next case
	and also prints 75 is lesser than 200. Fallthrough cannot be called on the final case of a switch
	 */
	/*
	fallthrough should be the last statement in a case. If it present somewhere in the middle, the compiler will throw
	error fallthrough statement out of place
	 */
}