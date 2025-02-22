package main

import "fmt"

func Generator(c chan<- int) {
	for i := 0; i < 10; i++ {
		c <- i
	}
	close(c)
}

func Double(in <-chan int, out chan int) {
	for i := range in {
		out <- i * 2
	}
	close(out)
}

func Print(c <-chan int) {
	for i := range c {
		fmt.Println(i)
	}
}

func main() {
	generator := make(chan int)
	doubler := make(chan int)
	go Generator(generator)
	go Double(generator, doubler)
	Print(doubler)
}
