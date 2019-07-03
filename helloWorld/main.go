package main

import (
	"first_Golang/try"
	"fmt"
	"math/rand"
)

//declare student1
var student1 = try.Student{
	Name: "Linda",
	Age:  22,
}

var m map[string]try.Student

func main() {
	//defer will run after return
	defer fmt.Println("bye")

	fmt.Println("----Program Start----")
	fmt.Printf("%s\t%d\n", student1.Name, student1.Age)

	my_map := make(map[string]try.Student)
	my_map["EE"] = try.Student{
		Name: "Owen",
		Age:  23,
	}

	fmt.Println(my_map["EE"])
	//make(type,len,size) to declare slice
	num := make([]int, 10)

	// int array with len 10
	// var num [10]int
	for i := 0; i < len(num); i++ {
		num[i] = rand.Intn(100)
	}

	// if input array, need to define the range
	// sum := try.Add_multi(num[:])
	// try.Insertionsort(num[:])

	// if input slice, nothing else need
	// sum := try.Add_multi(num)
	// fmt.Println(sum)
	try.Insertionsort(num)

	// for i := 0; i < len(num); i++ {
	// 	fmt.Printf("num[%d] = %v\n", i, num[i])
	// }

	// v will copy the value from the slice 'num'
	for i, v := range num {
		fmt.Printf("num[%d] = %d\n", i, v)
	}
}
