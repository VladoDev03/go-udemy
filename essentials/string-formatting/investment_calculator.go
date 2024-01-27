package main

import (
	"fmt"
	"math"
)

func main() {
	const inflationRate = 2.5
	var investmentAmount float64
	var expectedReturnRate float64
	var years float64

	fmt.Print("Investment Amount: ")
	fmt.Scan(&investmentAmount)

	fmt.Print("Expected Return Rate: ")
	fmt.Scan(&expectedReturnRate)

	fmt.Print("Years: ")
	fmt.Scan(&years)

	futureValue := investmentAmount * math.Pow(1+expectedReturnRate/1000, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	formattedFV := fmt.Sprintf("Future Real Value: %.2f", futureRealValue)

	fmt.Println("Future Value:", futureValue)
	fmt.Println("Future Real Value:", futureRealValue)

	fmt.Printf("Future Real Value: %v", futureRealValue)
	fmt.Println()
	fmt.Printf("Future Real Value: %.2f", futureRealValue)
	fmt.Println()
	fmt.Println(formattedFV)

	fmt.Printf(`Future Real Value: %.2f
	`, futureRealValue)
}
