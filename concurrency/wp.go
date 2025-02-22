package main

import "fmt"

func Worker(in int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Printf("Worker with id %d started fib with %d\n", in, j)
		fib := Fibonacci(j)
		fmt.Printf("Worker with id %d, job %d and fib %d\n", in, j, fib)
		results <- fib
	}
}

func Fibonacci(n int) int {
	if n <= 1 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	tasks := []int{10, 20, 30, 40, 50, 60, 70, 80, 90, 100}
	nWorkers := 3
	jobs := make(chan int, len(tasks))
	results := make(chan int, len(tasks))

	for i := 0; i < nWorkers; i++ {
		go Worker(i, jobs, results)
	}

	for _, task := range tasks {
		jobs <- task
	}
	close(jobs)

	for r := 0; r < len(tasks); r++ {
		<-results
	}
}
