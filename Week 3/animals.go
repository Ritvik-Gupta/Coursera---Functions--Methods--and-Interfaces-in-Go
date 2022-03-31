package main

import (
	"fmt"
	"strings"
)

type Animal struct {
	food, locomotion, noise string
}

func NewAnimal(food, locomotion, noise string) *Animal {
	return &Animal{food, locomotion, noise}
}

func (animal *Animal) Eat() {
	fmt.Println(animal.food)
}

func (animal *Animal) Move() {
	fmt.Println(animal.locomotion)
}

func (animal *Animal) Speak() {
	fmt.Println(animal.noise)
}

func (animal *Animal) PerformTrait(trait string) {
	switch strings.ToLower(trait) {
	case "eat":
		animal.Eat()
		break
	case "move":
		animal.Move()
		break
	case "speak":
		animal.Speak()
		break
	default:
		panic("Invalid Animal Trait to perform chosen")
	}
}

func main() {
	animals := map[string]*Animal{
		"cow":   NewAnimal("grass", "walk", "moo"),
		"bird":  NewAnimal("worms", "fly", "peep"),
		"snake": NewAnimal("mice", "slither", "hsss"),
	}

	fmt.Println("\nStarting Prompt for User Input for Animals")
	var animalName, animalTrait string
	for {
		fmt.Print(">  ")
		fmt.Scanln(&animalName, &animalTrait)

		if animalChosen, ok := animals[strings.ToLower(animalName)]; ok {
			animalChosen.PerformTrait(animalTrait)
		} else {
			panic("Invalid Animal chosen")
		}
	}
}
