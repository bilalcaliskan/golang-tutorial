package loops

import "fmt"

func RunLoops() {
	fmt.Printf("\nBeginning of loops...\n")
	/*
	for is the only loop available in Go. Go doesn't have while or do while loops which are present in other languages
	like C, Java etc.
	The syntax for the for loop is provided below:
		- for initialisation; condition; post {
		}
	The initialisation statement will be executed only once. After the loop is initialised, the condition will be checked.
	If the condition evaluates to true, the body of loop inside the { } will be executed followed by the post statement.
	The post statement will be executed after each successful iteration of the loop. After the post statement is executed,
	the condition will be rechecked. If it's true, the loop will continue executing, else the for loop terminates.
	All the three components namely initialisation, condition and post are optional in Go.
	 */
	for i := 1; i <= 10; i++ {
		fmt.Printf(" %d",i)
	}
	fmt.Println()

	fmt.Printf("\nBeginning of break statement...\n")
	/*
	The break statement is used to terminate the for loop abruptly before it finishes its normal execution and move the
	control to the line of code just after the for loop.
	 */
	for i := 1; i <= 10; i++ {
		if i > 5 {
			break //loop is terminated if i > 5
		}
		fmt.Printf("%d ", i)
	}
	fmt.Printf("\nline after for loop\n")

	fmt.Printf("\nBeginning of continue statement...\n")
	/*
	The continue statement is used to skip the current iteration of the for loop. All code present in a for loop after the
	continue statement will not be executed for the current iteration. The loop will move on to the next iteration.
	 */
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Printf("%d ", i)
	}
	fmt.Println()

	fmt.Printf("\nBeginning of nested for loops...\n")
	/*
	A for loop which has another for loop inside it is called a nested for loop.
	 */
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	fmt.Printf("\nBeginning of labels...\n")
	/*
	Labels can be used to break the outer loop from inside the inner for loop.
	 */
	/*
	Adding a break in the inner for loop when i and j are equal will only break from the inner for loop. This will break
	only from the inner for loop and the outer loop will continue.
	 */
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break
			}
		}
	}
	fmt.Println()
	/*
	This is not the intended output. We need to stop printing when both i and j are equal i.e when they are equal to 1.
	This is where labels come to our rescue. A label can be used to break from an outer loop. Let's rewrite the program
	above using labels.
	 */
	outer:
		for i := 0; i < 3; i++ {
			for j := 1; j < 4; j++ {
				fmt.Printf("i = %d , j = %d\n", i, j)
				if i == j {
					break outer
				}
			}
		}
	fmt.Println()

	fmt.Printf("\nBeginning of more examples...\n")
	/*
	Example 1
	The semicolons in the for loop of the below program can also be omitted.
	 */
	i := 0
	for ;i <= 10; { // initialisation and post are omitted. It can be written as for i <= 10 { }
		fmt.Printf("%d ", i)
		i += 2
	}
	fmt.Println()
	/*
	Example 2
	In the below program no and i are declared and initialised to 10 and 1 respectively. They are incremented by 1 at
	the end of each iteration. The boolean operator && is used in the condition to ensure that i is less than or equal
	to 10 and also no is less than or equal to 19.
	 */
	for no, i := 10, 1; i <= 10 && no <= 19; i, no = i + 1, no + 1 { //multiple initialisation and increment
		fmt.Printf("%d * %d = %d\n", no, i, no*i)
	}
	fmt.Println()

	fmt.Printf("\nBeginning of infinite loop...\n")
	/*
	The syntax for creating an infinite loop is:
		for {
		}
	The following program will keep printing Hello World continuously without terminating.
	*/
	for {
		fmt.Println("Hello World")
	}
}