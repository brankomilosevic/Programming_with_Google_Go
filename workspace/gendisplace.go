package main

import "fmt"

func main() {
	var acc float64
	fmt.Print("Enter acceleration.........: ")
	fmt.Scanln(&acc)

	var vel float64
	fmt.Print("Enter the initial velocity.: ")
	fmt.Scanln(&vel)

	var spc float64
	fmt.Print("Enter initial displacement.: ")
	fmt.Scanln(&spc)

	var tim float64
	fmt.Print("Enter time.................: ")
	fmt.Scanln(&tim)

	fn := GenDisplaceFn(acc, vel, spc)

	fmt.Println(fn(3))
	fmt.Println(fn(5))
}

func GenDisplaceFn(acc, vel, spc float64) func(float64) float64 {
	return func(t float64) float64 {
		return (1/2)*acc*t*t + vel*t + spc
	}
}
