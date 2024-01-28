package main

import "fmt"

func main() {
	fact := factorialIterative(10)
	fmt.Println(fact)
	fact = factorialRecursive(10)
	fmt.Println(fact)
}

func factorialIterative(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result = result * i
	}

	return result
}

func factorialRecursive(number int) int {
	if number == 0 {
		return 1
	}

	return number * factorialRecursive(number-1)
}
