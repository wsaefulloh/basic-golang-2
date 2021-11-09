package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg = sync.WaitGroup{}

func Append(slice []int, data int) []int {
	m := len(slice)
	newSlice := make([]int, (m + 1))
	copy(newSlice, slice)
	newSlice[m] = data
	return newSlice
}

func main() {
	var angka string
	fmt.Print("Tentukan bilangan yang ingin dicari faktor primanya (angka 0-100): ")
	fmt.Scanln(&angka)
	var nilai, err = strconv.Atoi(angka)
	if err == nil {
		fmt.Println("Main Routine Start")
		ch1 := make(chan int, 50)
		wg.Add(2)
		go makeArray(ch1, nilai)
		go seleksiPrima(ch1, nilai)

		wg.Wait()
		fmt.Println("\nMain Routine End")
	} else {
		fmt.Println(err.Error())
		fmt.Println("Error : Masukkan data angka")
	}

}

//Receiver channel and Sender channel
func makeArray(ch1 <-chan int, nilai int) {
	// i := <-ch
	y := 0
	array1 := []int{}
	for i := range ch1 {
		// fmt.Println("value from channel 1 ", i)
		array1 = Append(array1, i)
		y = y + 1
	}
	fmt.Println("Faktor prima bilangan", nilai, "adalah :")
	fmt.Println(array1)
	wg.Done()
}

//Sender channel only
func seleksiPrima(ch1 chan<- int, nilai1 int) {
	defer func() {
		close(ch1)
		wg.Done()
	}()
	i := nilai1
	for i > 0 {
		if (i % 2) == 0 {
			// fmt.Println("send data")
			ch1 <- 2
			i = i / 2
		} else if (i % 3) == 0 {
			// fmt.Println("send data")
			ch1 <- 3
			i = i / 3
		} else if (i % 5) == 0 {
			// fmt.Println("send data")
			ch1 <- 5
			i = i / 5
		} else if (i % 7) == 0 {
			// fmt.Println("send data")
			ch1 <- 7
			i = i / 7
		} else if (i % 11) == 0 {
			// fmt.Println("send data")
			ch1 <- 11
			i = i / 11
		} else if i == 1 {
			i = 0
		} else {
			// fmt.Println("send data")
			ch1 <- i
			i = 0
		}
	}
}
