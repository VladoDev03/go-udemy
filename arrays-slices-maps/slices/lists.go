package main

import "fmt"

func main() {
	var productNames [4]string = [4]string{"book", "magazine", "dictionary", "comics"}
	fmt.Println(productNames)

	prices := [4]float64{10.99, 9.99, 45.99, 20.0}
	fmt.Println(prices)

	featuredPrices := prices[:]
	fmt.Println(featuredPrices)

	featuredPrices = prices[1:3]
	fmt.Println(featuredPrices)

	featuredPrices = prices[1:]
	fmt.Println(featuredPrices)

	featuredPrices = prices[:3]
	fmt.Println(featuredPrices)

	featuredPrices = prices[1:]
	fmt.Println(featuredPrices)
	highlighted := prices[:1]
	fmt.Println(featuredPrices)

	fmt.Println(len(productNames), cap(productNames))
	fmt.Println(len(highlighted), cap(highlighted))

	numbers := []float64{1, 2, 3, 4, 5}

	fmt.Print("Add 6: ")
	updatedNumbers := append(numbers, 6)
	fmt.Println(numbers, updatedNumbers)

	fmt.Print("Remove 6: ")
	updatedNumbers = numbers[0:]
	fmt.Println(numbers, updatedNumbers)

	fmt.Print("Overwrite: ")
	fmt.Print(numbers)
	fmt.Print(" ")
	numbers = append(numbers, 6)
	fmt.Println(numbers)

	discountPrices := []float64{101.99, 80.99, 20.59}
	fmt.Println(updatedNumbers)
	fmt.Println(discountPrices)
	updatedNumbers = append(updatedNumbers, discountPrices...)
	fmt.Println(updatedNumbers)
}
