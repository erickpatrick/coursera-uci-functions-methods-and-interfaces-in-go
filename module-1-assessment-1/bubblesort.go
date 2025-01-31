package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Swap(slice int, index int) {

}

func BubbleSort(slice []int) {
	fmt.Println(slice)
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
			fmt.Println(err)
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
