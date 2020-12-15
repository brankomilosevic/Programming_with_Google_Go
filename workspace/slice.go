package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	numbers := make([]int, 0, 3)
	fmt.Println("Empty slice:", numbers)

	var input_str string

	fmt.Println("Enter integer numbers slice. [X] for end.")

	for true {
		fmt.Scanln(&input_str)
		if input_str == "X" {
			break
		}

		i, err := strconv.Atoi(input_str)

		if err != nil {
			fmt.Println("Incorrect input.")
			fmt.Println("Enter integer. [X] for end.")
			continue
		}

		numbers = append(numbers, i)

		sort.Ints(numbers[:])
		fmt.Println(numbers)
	}

	fmt.Println("Final sorted slice:", numbers)
}
