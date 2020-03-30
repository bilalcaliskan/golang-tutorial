package concurrency

import "fmt"

func RunConcurrency() {
	fmt.Printf("\nBeginning of introduction...\n")
	/*
	Go is a concurrent language and not a parallel one.
	Concurrency is the capability to deal with lots of things at once.
	Let's consider a person jogging. During his morning jog, lets say his shoe laces become untied. Now the person
	stops running, ties his shoe laces and then starts running again. This is a classic example of concurrency. The
	person is capable of handling both running and tying shoe laces, that is the person is able to deal with lots of
	things at once :)
	*/

	fmt.Printf("\nBeginning of parallelism vs concurrency...\n")
	/*
	Parallelism is doing lots of things at the same time.
	Lets understand it better with the same jogging example. In this case lets assume that the person is jogging and
	also listening to music in his iPod. In this case the person is jogging and listening to music at the same time,
	that is he is doing lots of things at the same time. This is called parallelism.

	Lets say we are programming a web browser. The web browser has various components. Two of them are the web page
	rendering area and the downloader for downloading files from the internet. Lets assume that we have structured our
	browser's code in such a way that each of these components can be executed independently (This is done using threads
	in languages such as Java and in Go we can achieve this using Goroutines, more on this later). When this browser is
	run in a single core processor, the processor will context switch between the two components of the browser. It
	might be downloading a file for some time and then it might switch to render the html of a user requested web page.
	This is know as concurrency. Concurrent processes start at different points of time and their execution cycles
	overlap. In this case the downloading and the rendering start at different points in time and their executions
	overlap.

	Lets say the same browser is running on a multi core processor. In this case the file downloading component and the
	HTML rendering component might run simultaneously in different cores. This is known as parallelism.

	Parallelism will not always result in faster execution times. This is because the components running in parallel
	have might have to communicate with each other. For example, in the case of our browser, when the file downloading
	is complete, this should be communicated to the user, say using a popup. This communication happens between the
	component responsible for downloading and the component responsible for rendering the user interface. This
	communication overhead is low in concurrent systems. In the case when components run in parallel in multiple cores,
	this communication overhead is high. Hence parallel programs do not always result in faster execution times!
	 */

	fmt.Printf("\nBeginning of support for concurrency in Go...\n")
	/*
	Concurrency is an natural part of the Go programming language. Concurrency is handled in Go using Goroutines
	and channels.
	 */
}