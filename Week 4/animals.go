package main

import "fmt"

type Animal interface {
	Eat()
	Move()
	Speak()
}

func PerformTraitForAnimal(animal Animal, trait string) {
	switch trait {
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

type Cow struct{}

func (*Cow) Eat() {
	fmt.Println("grass")
}

func (*Cow) Move() {
	fmt.Println("walk")
}

func (*Cow) Speak() {
	fmt.Println("moo")
}

type Bird struct{}

func (*Bird) Eat() {
	fmt.Println("worms")
}

func (*Bird) Move() {
	fmt.Println("fly")
}

func (*Bird) Speak() {
	fmt.Println("peep")
}

type Snake struct{}

func (*Snake) Eat() {
	fmt.Println("mice")
}

func (*Snake) Move() {
	fmt.Println("slither")
}

func (*Snake) Speak() {
	fmt.Println("hsss")
}

func NewAnimal(animalType string) Animal {
	switch animalType {
	case "cow":
		return &Cow{}
	case "bird":
		return &Bird{}
	case "snake":
		return &Snake{}
	default:
		panic("Invalid Animal Type chosen")
	}
}

func main() {
	animals := make(map[string]Animal)

	fmt.Println("\nStarting Prompt for User Input for Animals")
	for {
		fmt.Print(">  ")
		var commandType, animalName, typeOrTrait string
		fmt.Scanln(&commandType, &animalName, &typeOrTrait)

		switch commandType {
		case "newanimal":
			animals[animalName] = NewAnimal(typeOrTrait)
			fmt.Println("Created it!")
			break
		case "query":
			if animal, ok := animals[animalName]; ok {
				PerformTraitForAnimal(animal, typeOrTrait)
			} else {
				panic("Animal with the specific name does not exist")
			}
			break
		default:
			panic("Invalid Command Type specified")
		}
	}
}
