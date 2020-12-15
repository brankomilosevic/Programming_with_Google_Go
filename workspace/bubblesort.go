package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var numbers []int

	fmt.Println("Enter 10 integers in one line separated by space.")

	br := bufio.NewReader(os.Stdin)
	str, _, _ := br.ReadLine()
	tokens := strings.Split(string(str), " ")

	for _, s := range tokens {
		n, _ := strconv.Atoi(s)
		numbers = append(numbers, n)
	}

	BubbleSort(numbers)

	fmt.Println("Sorted array of integers:")
	fmt.Println(numbers)
}

func BubbleSort(a []int) {

	lenght := len(a)

	for i := 0; i < lenght; i++ {
		for j := 0; j < lenght-1-i; j++ {
			if a[j] > a[j+1] {
				Swap(a, j)
			}
		}
	}
}

func Swap(a []int, j int) {
	var tmp int
	tmp = a[j]
	a[j] = a[j+1]
	a[j+1] = tmp
}
