/*
The source file for Week 1, Exercise 1 bubblesort.go
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Swap function is exportable, per the instructions, thus the capital S.
func Swap(slice []int, index int) {
	first := slice[index]
	second := slice[index-1]
	slice[index] = second
	slice[index-1] = first
}

// BubbleSort function is exportable, per the instructions, thus the capital B.
func BubbleSort(slice []int) {
	slen := len(slice) - 1
	for i := 1; i <= slen; i++ {
		if slice[i-1] > slice[i] {
			Swap(slice, i)
		}
	}
}

func main() {
	sslice := []string{}
	sender := []int{}

	fmt.Println("please enter up to ten integers to sort.")
	// bufio Reader needed to test for example that contained spaces
	reader := bufio.NewReader(os.Stdin)
	var line string
	line, _ = reader.ReadString('\n')
	sslice = strings.Fields(line)

	for i := range sslice {
		number := sslice[i]
		num, _ := strconv.Atoi(number)
		sender = append(sender, num)
	}

	for {
		if sort.IntsAreSorted(sender) {
			fmt.Println("The output of your sorted slice is: ")
			fmt.Println(sender)
			break
		} else {
			BubbleSort(sender)
		}
	}
}
