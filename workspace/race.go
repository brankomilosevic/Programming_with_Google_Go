package main

import (
	"fmt"
)

/*
"Race Condition" occurs when more than one thread have access to same piece of data with un-deterministic manner.
Change of data by one thread potentially leads to problems in another thread.

REGULAR case:   In case [y] is assigned to value [5] before [x] is assigned to value [100] and before division function takes place, the division [x/y] is regular.
IRREGULAR case: In case division [x/y] is attempted before [y] is assigned a value different of [0], then we get “division by zero” error.

Logical Processor performs the scheduling mechanism of go routines and, eventually, we will end up with attempt to divide [x] by [0] which will lead to runtime error.
*/

func main() {
	var i int
	var res int

	for true {
		var x, y int // initialized by zero

		go func(v *int) {
			*v = 100
		}(&x)

		go func(v *int) {
			*v = 5
		}(&y)

		go func(v1 int, v2 int) {
			if v2 == 0 {
				fmt.Println("... aaaa, devide by zero !!!")
			}
			res = v1 / v2
		}(x, y)

		i += 1

		fmt.Printf("attempt number: %d\n", i)
	}
}
