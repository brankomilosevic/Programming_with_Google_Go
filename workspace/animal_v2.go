package main

import (
	"fmt"
)

type animal struct {
	food       string
	locomotion string
	noise      string
}

type animalInterface interface {
	Eat()
	Move()
	Speak()
}

func (a animal) Eat() {
	fmt.Println(a.food)
	return
}

func (a animal) Move() {
	fmt.Println(a.locomotion)
	return
}

func (a animal) Speak() {
	fmt.Println(a.noise)
	return
}

func main() {
	var animal_if animalInterface
	var animals_table map[string]animal
	var command, request_animal_name, request_data string

	animals_table = make(map[string]animal)
	animals_table["cow"] = animal{"grass", "walk", "moo"}
	animals_table["bird"] = animal{"worms", "fly", "peep"}
	animals_table["snake"] = animal{"mice", "slither", "hsss"}

	fmt.Println("Here is the animals table")
	fmt.Print(animals_table)
	fmt.Println("Enter the animal type and request:")

	for true {

		fmt.Print("> ")

		fmt.Scan(&command, &request_animal_name, &request_data)

		switch command {
		case "newanimal":
			animals_table[request_animal_name] = animals_table[request_data]
		case "query":
			animal_if = animals_table[request_animal_name]
			switch request_data {
			case "eat":
				animal_if.Eat()
			case "move":
				animal_if.Move()
			case "speak":
				animal_if.Speak()
			default:
				fmt.Println("Something's wrong, please check the input.")
			}

		default:
			fmt.Println("Something's wrong, please check the input.")
		}
	}
}
