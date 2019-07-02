package main

import (
	"first_Golang/try"
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Hello world!")
	num := make([]int, 10)
	for i := 0; i < len(num); i++ {
		num[i] = rand.Intn(100)
	}
	// sum := try.Add_multi(num)
	try.Insertionsort(num)
	for i := 0; i < len(num); i++ {
		fmt.Printf("num[%d] = %v\n", i, num[i])
	}

}

// 1D
// 2C
// 3C
