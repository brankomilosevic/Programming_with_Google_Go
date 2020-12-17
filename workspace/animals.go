package main

import (
	"fmt"
)

type Animal struct {
	food, locomotion, sound string
}

func (v Animal) Eat() {
	fmt.Println(v.food)
}

func (v Animal) Move() {
	fmt.Println(v.locomotion)
}

func (v Animal) Speak() {
	fmt.Println(v.sound)
}

func main() {
	var animal_name string
	var animal_request string

	animals_table := map[string]Animal{
		"cow":   Animal{"grass", "walk", "moo"},
		"bird":  Animal{"worms", "fly", "peep"},
		"snake": Animal{"mice", "slither", "hsss"},
	}

	fmt.Println("Here is the animals table")
	fmt.Print(animals_table)
	fmt.Println("Enter the animal type and request:")

	for true {
		fmt.Print("> ")

		fmt.Scan(&animal_name, &animal_request)

		switch animal_request {
		case "eat":
			animals_table[animal_name].Eat()
		case "move":
			animals_table[animal_name].Move()
		case "speak":
			animals_table[animal_name].Speak()
		default:
			fmt.Println("Something's wrong, please check the input.")
		}
	}
}