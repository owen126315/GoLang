package try

var My_str = "hi"

func Add(a int, b int) int {
	return a + b
}

func Add_multi(x []int) int {
	var sum int
	for _, n := range x {
		sum += n
	}
	return sum
}

func insersort(num *[]int) {

}
