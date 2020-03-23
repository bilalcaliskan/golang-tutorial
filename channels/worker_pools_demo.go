package channels

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var jobs = make(chan Job, 10)
var results = make(chan Result, 10)

type Job struct {
	id int
	randomNo int
}

type Result struct {
	job Job
	sumOfDigits int
}

func digits(number int) int {
	sum := 0
	no := number
	for no != 0 {
		digit := no % 10
		sum += digit
		no /= 10
	}
	time.Sleep(2 * time.Second)
	return sum
}

func worker(wg *sync.WaitGroup) {
	for job := range jobs {
		output := Result{job, digits(job.randomNo)}
		results <- output
	}
	wg.Done()
}

func createWorkerPool(noOfWorkers int) {
	var wg sync.WaitGroup
	for i := 0; i < noOfWorkers; i++ {
		wg.Add(1)
		go worker(&wg)
	}
	wg.Wait()
	fmt.Println("Closing results channel...")
	close(results)
}

func allocate(noOfJobs int) {
	for i := 0; i < noOfJobs; i++ {
		randomNo := rand.Intn(999)
		job := Job{
			id:       i,
			randomNo: randomNo,
		}
		jobs <- job
	}
	fmt.Println("Closing jobs channel...")
	close(jobs)
}

func result(done chan bool) {
	for result := range results {
		fmt.Printf("Len results %d, Job id %d, input random no %d , sum of digits %d\n", len(results),
			result.job.id, result.job.randomNo, result.sumOfDigits)
	}
	done <- true
}

func RunWorkerPoolsDemo() {
	startTime := time.Now()
	noOfJobs := 100
	go allocate(noOfJobs)
	done := make(chan bool)
	go result(done)
	noOfWorkers := 10
	createWorkerPool(noOfWorkers)
	<- done
	endTime := time.Now()
	diff := endTime.Sub(startTime)
	fmt.Println("total time taken ", diff.Seconds(), "seconds")
}