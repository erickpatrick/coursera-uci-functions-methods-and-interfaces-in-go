package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(slice []int, index int) {
	if slice[index] > slice[index+1] {
		slice[index], slice[index+1] = slice[index+1], slice[index]
	}
}

func BubbleSort(slice []int) {
	sliceLength := len(slice)
	for range slice[:sliceLength] {
		for inner := range slice[:sliceLength-1] {
			Swap(slice, inner)
		}
	}
	for _, number := range slice {
		fmt.Print(number, " ")
	}
}

func GetUserInput() []int {
restart:
	fmt.Println("Please, type a series of 10 integers separated by space")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	explodedInput := strings.Split(scanner.Text(), " ")
	numbers := make([]int, 0, 1)

	for _, value := range explodedInput {
		val, err := strconv.Atoi(value)
		if err != nil {
			goto restart
		}
		numbers = append(numbers, val)
	}

	if len(numbers) < 10 {
		goto restart
	}

	return numbers
}

func main() {
	numbers := GetUserInput()
	BubbleSort(numbers)
}
