package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"book", "magazine", "dictionary", "comics"}
	fmt.Println(productNames)
	fmt.Println(productNames[0])
	fmt.Println(productNames[1])
	fmt.Println(productNames[2])
	fmt.Println(productNames[3])

	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)
	fmt.Println(prices[0])
	fmt.Println(prices[1])
	fmt.Println(prices[2])
	fmt.Println(prices[3])
}
