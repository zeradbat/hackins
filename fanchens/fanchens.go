package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func zero(a, b, v, c, s, z int) int {
	return 0
}

func vals() (int, int) {
	return 3, 7
}

func sum(numbs ...int) {
	fmt.Println(numbs)
	total := 0
	for _, numb := range numbs {
		total += numb
	}
	fmt.Println(total)
}

func main() {
	res := plus(1, 2)
	fmt.Println(res)

	res2 := res + zero(1, 2, 3, 4, 5, res)
	fmt.Println(res2)

	res3, res4 := vals()
	fmt.Println(res3, res4)

	fmt.Println("summins")
	sum(11, 1, 1, 1, 1, 1, 1, 1, 1, 11, 1, 1, 1, 1, 1, 1, 1, 1, 11)
}
