package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

func main() {
	fmt.Println("Main Routine Start")
	ch := make(chan int)
	wg.Add(2)

	go cetakPrima(ch)
	go seleksiFibonacci(ch, 150)

	wg.Wait()
	fmt.Println("\nMain Routine End")
}

//Receiver channel only
func cetakPrima(ch <-chan int) {
	// i := <-ch
	for i := range ch {
		if i%2 == 0 {
			fmt.Println("bilangan genap dengan nilai ", i)
		} else {
			fmt.Println("bilangan ganjil dengan nilai ", i)
		}
	}
	wg.Done()
}

//Sender channel only
func seleksiFibonacci(ch chan<- int, sum int) {
	defer func() {
		close(ch)
		wg.Done()
	}()

	i := 0
	j := 1
	bilangan_fibonacci := 1
	for i >= 0 {
		if bilangan_fibonacci <= sum {
			ch <- bilangan_fibonacci
		} else {
			break
		}
		bilangan_fibonacci = i + j
		i = j
		j = bilangan_fibonacci
	}
}
