package main

import (
	"first_Golang/try"
	"fmt"
)

//declare student1
var student1 = try.Student{
	Name: "Linda",
	Age:  22,
}

func main() {
	//defer will run after return
	defer fmt.Println("bye")

	fmt.Println("----Program Start----")

	fmt.Printf("%s\t%d\n", student1.Name, student1.Age)
	//create a pointer and set student1's address to it
	var student2 *try.Student
	student2 = &student1
	student2.Name = "Owen"
	fmt.Printf("%s\t%d\n", student2.Name, student2.Age)
	// //make(type,len,size) to declare slice
	// //slice is pass by reference, similar to pointer
	// num := make([]int, 10)
	// for i := 0; i < len(num); i++ {
	// 	num[i] = rand.Intn(100)
	// }
	// // sum := try.Add_multi(num)
	// try.Insertionsort(num)
	// for i := 0; i < len(num); i++ {
	// 	fmt.Printf("num[%d] = %v\n", i, num[i])
	// }
}
