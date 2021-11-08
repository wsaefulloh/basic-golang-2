package main

import (
	"fmt"
	"strconv"
	"sync"
)

var wg = sync.WaitGroup{}

type deretBilang struct {
	limit int
}

func Append(slice []int, data int) []int {
	m := len(slice)
	newSlice := make([]int, (m + 1))
	copy(newSlice, slice)
	newSlice[m] = data
	return newSlice
}

func (deret deretBilang) prima() {
	n := deret.limit
	y := 0
	array := []int{}
	for i := 0; i <= n; i++ {
		if (i%2) != 0 && (i%3) != 0 && (i%5) != 0 {
			array = Append(array, i)
			y = y + 1
		}
	}
	fmt.Println(array)
	wg.Done()
}

func (deret deretBilang) ganjil() {
	n := deret.limit
	y := 0
	array := []int{}
	for i := 0; i <= n; i++ {
		if (i % 2) != 0 {
			array = Append(array, i)
			y = y + 1
		}
	}
	fmt.Println(array)
	wg.Done()
}

func (deret deretBilang) genap() {
	n := deret.limit
	y := 0
	array := []int{}
	for i := 0; i <= n; i++ {
		if (i % 2) == 0 {
			array = Append(array, i)
			y = y + 1
		}
	}
	fmt.Println(array[1:y])
	wg.Done()
}

func (deret deretBilang) fibonacci() {
	n := deret.limit
	i := 0
	j := 1
	array := []int{1}
	for i >= 0 {
		bilangan_fibonacci := i + j
		if bilangan_fibonacci <= n {
			array = Append(array, bilangan_fibonacci)
		} else {
			break
		}
		i = j
		j = bilangan_fibonacci
	}
	fmt.Println(array)
	wg.Done()
}

func main() {
	var limits string
	fmt.Print("Tentukan limit (angka): ")
	fmt.Scanln(&limits)
	var jenis string
	fmt.Print("Tentukan jenis deret bilangan: ")
	fmt.Scanln(&jenis)
	var lim, err = strconv.Atoi(limits)
	if err == nil {
		deret := deretBilang{lim}
		if jenis == "prima" {
			wg.Add(1)
			fmt.Println("limit deret angka", deret.limit)
			fmt.Println("deret bilangan prima mulai dari 0 -", deret.limit, "adalah")
			go deret.prima()
			wg.Wait()
		} else if jenis == "ganjil" {
			wg.Add(1)
			fmt.Println("limit deret angka", deret.limit)
			fmt.Println("deret bilangan ganjil mulai dari 0 -", deret.limit, "adalah")
			go deret.ganjil()
			wg.Wait()
		} else if jenis == "genap" {
			wg.Add(1)
			fmt.Println("limit deret angka", deret.limit)
			fmt.Println("deret bilangan genap mulai dari 0 -", deret.limit, "adalah")
			go deret.genap()
			wg.Wait()
		} else if jenis == "fibonacci" {
			wg.Add(1)
			fmt.Println("limit deret angka", deret.limit)
			fmt.Println("deret bilangan fibonacci mulai dari 0 -", deret.limit, "adalah")
			go deret.fibonacci()
			wg.Wait()
		} else {
			fmt.Println("Jenis deret tidak diketahui")
		}
	} else {
		fmt.Println(err.Error())
		fmt.Println("Error : Masukkan data angka")
	}
}
