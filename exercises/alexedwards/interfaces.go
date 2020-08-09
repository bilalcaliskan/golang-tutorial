package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

/*
One example of an interface type from the standard library is the fmt.Stringer interface, which looks like this.
We say that something satisfies this interface (or implements this interface) if it has a method with the exact
signature String() string.
 */
type Stringer interface {
	String() string
}

/*
For example, the following Book type satisfies the fmt.Stringer interface because it has a String() string method.
It's not really important what this Book type is or does. The only thing that matters is that is has a method called
String() which returns a string value.
 */
type Book struct {
	Title  string
	Author string
}

func (b Book) String() string {
	return fmt.Sprintf("Book: %s - %s", b.Title, b.Author)
}

/*
Or, as another example, the following Count type also satisfies the fmt.Stringer interface — again because it has a
method with the exact signature String() string.
 */
type Count int

func (c Count) String() string {
	return strconv.Itoa(int(c))
}

// Declare a WriteLog() function which takes any object that satisfies
// the fmt.Stringer interface as a parameter.
func WriteLog(s fmt.Stringer) {
	log.Println(s.String())
}

// Create a Customer type
type Customer struct {
	Name string
	Age  int
}

// Implement a WriteJSON method that takes an io.Writer as the parameter.
// It marshals the customer struct to JSON, and if the marshal worked
// successfully, then calls the relevant io.Writer's Write() method.
func (c *Customer) WriteJSON(w io.Writer) error {
	js, err := json.Marshal(c)
	if err != nil {
		return err
	}

	_, err = w.Write(js)
	return err
}

func runInterfaces() {
	/*
	The important thing to grasp is that we have two different types, Book and Count, which do different things. But
	the thing they have in common is that they both satisfy the fmt.Stringer interface.
	You can think of this the other way around too. If you know that an object satisfies the fmt.Stringer interface, you
	can rely on it having a method with the exact signature String() string that you can call.

	Now for the important part:
	Wherever you see declaration in Go (such as a variable, function parameter or struct field) which has an interface
	type, you can use an object of any type so long as it satisfies the interface.
	For example, let's say that you have the following function:
		func WriteLog(s fmt.Stringer) {
			log.Println(s.String())
		}
	Because this WriteLog() function uses the fmt.Stringer interface type in its parameter declaration, we can pass in
	any object that satisfies the fmt.Stringer interface. For example, we could pass either of the Book and Count types
	that we made earlier to the WriteLog() method, and the code would work OK.

	Additionally, because the object being passed in satisfies the fmt.Stringer interface, we know that it has a String()
	string method that the WriteLog() function can safely call.
	 */

	/*
	Let's put this together in an example, which gives us a peek into the power of interfaces.
	 */
	book := Book{"Alice in Wonderland", "Lewis Carrol"}
	WriteLog(book)

	// Initialize a Count object and pass it to WriteLog().
	count := Count(3)
	WriteLog(count)
	/*
	the key thing to take away is that by using a interface type in our WriteLog() function declaration, we have made the
	function agnostic (or flexible) about the exact type of object it receives. All that matters is what methods it has.
	*/

	/*
	Q: Why interfaces are useful?
	A: There are all sorts of reasons that you might end up using a interface in Go, but in my experience the three
	most common are:
		To help reduce duplication or boilerplate code.
	    To make it easier to use mocks instead of real objects in unit tests.
	    As an architectural tool, to help enforce decoupling between parts of your codebase.
	*/

	/*
	Reducing boilerplate code:
	OK, imagine that we have a Customer struct containing some data about a customer. In one part of our codebase we want
	to write the customer information to a bytes.Buffer, and in another part of our codebase we want to write the customer
	information to an os.File on disk. But in both cases, we want to serialize the customer struct to JSON first.
	This is a scenario where we can use Go's interfaces to help reduce boilerplate code.
	The first thing you need to know is that Go has an io.Writer interface type which looks like this:
		type Writer interface {
	        Write(p []byte) (n int, err error)
		}
	And we can leverage the fact that both bytes.Buffer and the os.File type satisfy this interface, due to them having
	the bytes.Buffer.Write() and os.File.Write() methods respectively.
	Let's take a look at a simple implementation:
	*/
	// Initialize a customer struct.
	c := &Customer{Name: "Alice", Age: 21}

	// We can then call the WriteJSON method using a buffer...
	var buf bytes.Buffer
	err := c.WriteJSON(&buf)
	if err != nil {
		log.Fatal(err)
	}

	// Or using a file.
	f, err := os.Create("/tmp/customer")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()


	err = c.WriteJSON(f)
	if err != nil {
		log.Fatal(err)
	}
	/*
	Of course, this is just a toy example (and there are other ways we could structure the code to achieve the same end
	result). But it nicely illustrates the benefit of using an interface — we can create the Customer.WriteJSON() method
	once, and we can call that method any time that we want to write to something that satisfies the io.Writer interface.

	Q: But if you're new to Go, this still begs a couple of questions: How do you know that the io.Writer interface even
	exists? And how do you know in advance that bytes.Buffer and os.File both satisfy it?
	A: There's no easy shortcut here I'm afraid — you simply need to build up experience and familiarity with the interfaces
	and different types in the standard library. Spending time thoroughly reading the standard library documentation, and
	looking at other people's code will help here.
	*/

	/*
	Application architecture:
	Let's say that you are building a web application which interacts with a database. If you create an interface that
	describes the exact methods for interacting with the database, you can refer to the interface throughout your HTTP
	handlers instead of a concrete type. Because the HTTP handlers only refer to an interface, this helps to decouple
	the HTTP layer and database-interaction layer. It makes it easier to work on the layers independently, and to swap
	out one layer in the future without affecting the other.
	Check more on https://www.alexedwards.net/blog/organising-database-access
	 */

	/*
	Q: What is the empty interface?
	A: An interface type in Go is kind of like a definition. It defines and describes the exact methods that some other
	type must have. The empty interface type essentially describes no methods. It has no rules. And because of that, it
	follows that any and every object satisfies the empty interface. Or to put it in a more plain-English way, the empty
	interface type interface{} is kind of like a wildcard. Wherever you see it in a declaration (such as a variable,
	function parameter or struct field) you can use an object of any type.
	Take a look at the following code:
	 */
	person := make(map[string]interface{}, 0)
	person["name"] = "Alice"
	person["age"] = 21
	person["height"] = 167.64

	fmt.Printf("%+v", person)
	/*
	In this code snippet we initialize a person map, which uses the string type for keys and the empty interface type
	interface{} for values. We've assigned three different types as the map values (a string, int and float32) — and
	that's OK. Because objects of any and every type satisfy the empty interface, the code will work just fine.
	But there's an important thing to point out when it comes to retrieving and using a value from this map.
	If you try to get the "age" key and increment it by 1, compiler will complain about that "mismatched types interface {} and int".
	This happens because the value stored in the map takes on the type interface{}, and ceases to have it's original,
	underlying, type of int. Because it's no longer an int type we cannot add 1 to it.
	To get around this this, you need to type assert the value back to an int before using it. Like so:
	 */
	age, ok := person["age"].(int)
	if !ok {
		log.Fatal("could not assert value to int")
		return
	}
	person["age"] = age + 1
	log.Printf("%+v", person)

	/*
	Q: So when should you use the empty interface type in your own code?
	A: The answer is probably not that often. If you find yourself reaching for it, pause and consider whether using
	interface{} is really the right option. As a general rule it's clearer, safer and more performant to use concrete
	types — or non-empty interface types — instead. In the code snippet above, it would have been more appropriate to
	define a Person struct with relevant typed fields similar to this:
		type Person struct {
			Name   string
			Age    int
			Height float32
		}
	But that said, the empty interface is useful in situations where you need to accept and work with unpredictable or
	user-defined types.
	*/

	/*
	Q: What are the common and useful types?
	A: Lastly, here's a short list of some of the most common and useful interfaces in the standard library. If you're
	not familiar with them already, then I recommend taking out a bit of time to look at the relevant documentation for them.
	    builtin.Error
	    fmt.Stringer
	    io.Reader
	    io.Writer
	    io.ReadWriteCloser
	    http.ResponseWriter
	    http.Handler
	There is also a longer and more comprehensive listing of standard libraries available in https://gist.github.com/asukakenji/ac8a05644a2e98f1d5ea8c299541fce9
	*/
}