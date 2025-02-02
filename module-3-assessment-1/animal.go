package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (animal *Animal) Eat() string {
	return animal.food
}

func (animal *Animal) Move() string {
	return animal.locomotion
}

func (animal *Animal) Speak() string {
	return animal.noise
}

func main() {
	cow := Animal{food: "grass", locomotion: "walk", noise: "moo"}
	bird := Animal{food: "worms", locomotion: "fly", noise: "peep"}
	snake := Animal{food: "mice", locomotion: "slither", noise: "hss"}

	for {
		fmt.Println("Choose either cow, bird, or snake. Ask for eat, move, speak. Ex: 'cow move'")
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		choices := strings.Split(scanner.Text(), " ")
		var animal Animal

		switch choices[0] {
		case "cow":
			animal = cow
		case "bird":
			animal = bird
		case "snake":
			animal = snake
		default:
			continue
		}

		switch choices[1] {
		case "eat":
			fmt.Println(animal.Eat())
		case "move":
			fmt.Println(animal.Move())
		case "speak":
			fmt.Println(animal.Speak())
		default:
			continue
		}
	}
}
