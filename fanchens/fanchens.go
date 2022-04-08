package main

import "fmt"

func plus(a int, b int) int {
	return a + b
}

func zero(a, b, v, c, s, z int) int {
	return 0
}

func main() {
	res := plus(1, 2)
	fmt.Println(res)

	res2 := res + zero(1, 2, 3, 4, 5, res)
	fmt.Println(res2)
}
