// `for` is Go's only looping construct. Here are
// some basic types of `for` loops.

package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	// The most basic type, with a single condition.
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// A classic initial/condition/after `for` loop.
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// `for` without a condition will loop repeatedly
	// until you `break` out of the loop or `return` from
	// the enclosing function.
	for {
		fmt.Println("whoop whoop")

		// get a rando num
		s1 := rand.NewSource(time.Now().UnixNano())
		randofak := rand.New(s1).Intn(100000)

		if randofak%100 == 0 {
			fmt.Println("Breaking out of randofak loop")
			break // break here
		}

		fmt.Println("flatso")
	}

	// You can also `continue` to the next iteration of
	// the loop.
	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
