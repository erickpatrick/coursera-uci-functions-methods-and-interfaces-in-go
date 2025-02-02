package main

import (
	"fmt"
	"strconv"
)

func main() {
	var inputAcceleration string
	var inputInitialVelocity string
	var inputInitialDisplacement string
	var inputTime string

	fmt.Println("What's the acceleration?")
	fmt.Scan(&inputAcceleration)
	accelaration, err := strconv.ParseFloat(inputAcceleration, 64)
	if err != nil {
		fmt.Println("Wrong value for accelaration. Please enter a valid number")
	}

	fmt.Println("What's the initial velocity?")
	fmt.Scan(&inputInitialVelocity)
	initialVelocity, err := strconv.ParseFloat(inputInitialVelocity, 64)
	if err != nil {
		fmt.Println("Wrong value for initial velocity. Please enter a valid number")
	}

	fmt.Println("What's the initial displacement?")
	fmt.Scan(&inputInitialDisplacement)
	initialDisplacement, err := strconv.ParseFloat(inputInitialDisplacement, 64)
	if err != nil {
		fmt.Println("Wrong value for initial displacement. Please enter a valid number")
	}

	fmt.Println("What's duration - time -  in seconds?")
	fmt.Scan(&inputTime)
	time, err := strconv.ParseFloat(inputTime, 64)
	if err != nil {
		fmt.Println("Wrong value for duration/time. Please enter a valid number")
	}

	fn := GenDisplaceFn(accelaration, initialVelocity, initialDisplacement)
	fmt.Println(fn(time))
}

func GenDisplaceFn(acceleration float64, initialVelocity float64, initialDisplacement float64) func(time float64) float64 {
	return func(time float64) float64 {
		return (0.5 * acceleration * (time * time)) + (initialVelocity * time) + initialDisplacement
	}
}
