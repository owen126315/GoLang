package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello world!")
	var num = make([]int, 10)
	for i := 0; i < len(num); i++ {
		num[i] = rand.Intn(100)
	}
	// sum := try.Add_multi(num)
	insertionsort(num)
	fmt.Println(num)
}

func insertionsort(items []int) {
	var n = len(items)
	for i := 1; i < n; i++ {
		j := i
		for j > 0 {
			if items[j-1] > items[j] {
				items[j-1], items[j] = items[j], items[j-1]
			}
			j = j - 1
		}
	}
}
