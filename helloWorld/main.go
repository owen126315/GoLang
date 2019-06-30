package main

import (
	"first_Golang/try"
	"fmt"
)

func main() {
	fmt.Println("Hello world!")
	// for i := 0; i < 100; i++ {
	// 	// fmt.Println(i)
	// 	switch i {
	// 	case 10:
	// 		fmt.Println(10)
	// 	case 20:
	// 		fmt.Println(try.My_str)
	// 	}
	// }
	// fmt.Println(try.Add(50, 2))
	var sum int
	num := []int{1, 6, 3}
	sum = try.Add_multi(num)
	fmt.Println(sum)
}
