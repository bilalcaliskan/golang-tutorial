package mutexes

import (
	"fmt"
	"sync"
)

var x int = 0

func increment(wg *sync.WaitGroup) {
	x = x + 1
	wg.Done()
}

func incrementWithMutex(wg *sync.WaitGroup, m *sync.Mutex) {
	/*
	Mutex is a struct type and we create a zero valued variable m of type Mutex in line no. 15. In the above program
	we have changed the increment function so that the code which increments x x = x + 1 is between m.Lock() and
	m.Unlock(). Now this code is void of any race conditions since only one Goroutine is allowed to execute this
	piece of code at any point of time.
	 */
	m.Lock()
	x = x + 1
	m.Unlock()
	wg.Done()
}

func incrementWithChannel(wg *sync.WaitGroup, ch chan bool) {
	/*
	we have created a buffered channel of capacity 1 and this is passed to the increment Goroutine in line no. 18.
	This buffered channel is used to ensure that only one Goroutine access the critical section of code which
	increments x. This is done by passing true to the buffered channel in line no. 8 just before x is incremented.
	Since the buffered channel has a capacity of 1, all other Goroutines trying to write to this channel are blocked
	until the value is read from this channel after incrementing x in line no. 10. Effectively this allows only one
	Goroutine to access the critical section.
	 */
	ch <- true
	x = x + 1
	<- ch
	wg.Done()
}

func RunMutexes() {
	fmt.Printf("\nBeginning of critical selection...\n")
	/*
	Race condition is the term that output of the program depends on the sequence of execution of Goroutines.
	When a program runs concurrently, the parts of code which modify shared resources should not be accessed by
	multiple Goroutines at the same time. This section of code which modifies shared resources is called critical section.
	The race condition could have been avoided if only one Goroutine was allowed to access the critical section of
	the code at any point of time. This is made possible by using Mutex.
	*/

	fmt.Printf("\nBeginning of mutex...\n")
	/*
	A Mutex is used to provide a locking mechanism to ensure that only one Goroutine is running the critical section
	of code at any point of time to prevent race condition from happening.
	Mutex is available in the sync package. There are two methods defined on Mutex namely Lock and Unlock. Any code
	that is present between a call to Lock and Unlock will be executed by only one Goroutine, thus avoiding race condition.
	 */
	// In the below code, x = x + 1 will be executed by only one Goroutine at any point of time thus preventing race condition.
	/*
	mutex.Lock()
	x = x + 1
	mutex.Unlock()
	*/
	/*
	If one Goroutine already holds the lock and if a new Goroutine is trying to acquire a lock, the new Goroutine will
	be blocked until the mutex is unlocked.
	 */

	fmt.Printf("\nBeginning of a simple program with race condition...\n")
	var w sync.WaitGroup
	for i := 0; i < 1000; i++ {
		w.Add(1)
		go increment(&w)
	}
	w.Wait()
	fmt.Println("final value of x", x)

	fmt.Printf("\nBeginning of solution of above problem with mutex...\n")
	x = 0
	var w1 sync.WaitGroup
	var m1 sync.Mutex
	for i := 0; i < 1000; i++ {
		w1.Add(1)
		go incrementWithMutex(&w1, &m1)
	}
	w1.Wait()
	fmt.Println("final value of x after solved with mutex", x)

	fmt.Printf("\nBeginning of solution of above problem with channel...\n")
	x = 0
	var w2 sync.WaitGroup
	ch := make(chan bool, 1)
	for i := 0; i < 1000; i++ {
		w2.Add(1)
		go incrementWithChannel(&w2, ch)
	}
	w2.Wait()
	fmt.Println("final value of x after solved with channel", x)

	fmt.Printf("\nBeginning of mutex vs channel comparison...\n")
	/*
	In general use channels when Goroutines need to communicate with each other and mutexes when only one Goroutine
	should access the critical section of code.
	 */
}