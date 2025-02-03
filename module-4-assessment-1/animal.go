package main

import (
	"fmt"
	"strings"
)

type Animal interface {
	Eat()
	Move()
	Speak()
}

type Cow struct{}
type Bird struct{}
type Snake struct{}

func (cow Cow) Eat() {
	fmt.Println("grass")
}

func (bird Bird) Eat() {
	fmt.Println("worms")
}

func (snake Snake) Eat() {
	fmt.Println("mice")
}

func (a Cow) Move() {
	fmt.Println("walk")
}

func (bird Bird) Move() {
	fmt.Println("fly")
}

func (snake Snake) Move() {
	fmt.Println("slither")
}

func (cow Cow) Speak() {
	fmt.Println("moo")
}

func (bird Bird) Speak() {
	fmt.Println("peep")
}

func (snake Snake) Speak() {
	fmt.Println("hss")
}

func Information(a Animal, info string) {
	switch info {
	case "eat":
		a.Eat()
	case "speak":
		a.Speak()
	case "move":
		a.Move()
	default:
		fmt.Println("Invalid information asked for animal. Please, try again")
	}
}

func main() {
	animalTemplate := map[string]Animal{
		"cow":   Cow{},
		"bird":  Bird{},
		"snake": Snake{},
	}

	animals := map[string]Animal{}

	var command string
	var name string
	var extra string

	for {
		fmt.Print("> ")
		fmt.Scan(&command, &name, &extra)

		if command == "" || name == "" || extra == "" {
			fmt.Println("Invalid number of arguments. Please, try again:")
			continue
		}

		switch strings.ToLower(command) {
		case "newanimal":
			animals[strings.ToLower(name)] = animalTemplate[strings.ToLower(extra)]
			fmt.Println("Created it!")
		case "query":
			animal := animals[strings.ToLower(name)]
			Information(animal, strings.ToLower(extra))
		default:
			fmt.Println("Invalid request. Please, try again:")
		}
	}
}
