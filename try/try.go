package try

//try Struct
type Student struct {
	Name string
	Age  int
}

//start with Capital letter means public
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

func Insertionsort(items []int) {
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
