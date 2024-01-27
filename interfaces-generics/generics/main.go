package main

import "fmt"

func main() {
	result := add(2, 3)
	fmt.Println(result)
}

// func add[T any](a, b T) T {
// func add[T interface{}] (a, b T) T {
func add[T int | float64 | string](a, b T) T {
	return a + b
}
