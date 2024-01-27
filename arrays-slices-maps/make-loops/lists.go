package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	userNames := make([]string, 2, 5)
	userNames[0] = "vlad111"
	userNames[1] = "vladeto17ti"
	userNames = append(userNames, "vladsto")
	userNames = append(userNames, "vladodev03")
	fmt.Println(userNames)

	courseRatings1 := make(map[string]float64, 3)
	courseRatings1["flutter"] = 4.7
	courseRatings1["react"] = 4.8
	courseRatings1["angular"] = 4.6
	courseRatings1["go"] = 4.7
	fmt.Println(courseRatings1)

	courseRatings2 := make(floatMap, 3)
	courseRatings2["flutter"] = 4.7
	courseRatings2["react"] = 4.8
	courseRatings2["angular"] = 4.6
	courseRatings2["go"] = 4.7
	courseRatings2.output()

	for range userNames {
		fmt.Println("username placeholder")
	}

	for index, value := range userNames {
		fmt.Println(index+1, value)
	}

	for key, value := range courseRatings1 {
		fmt.Println(key, value)
	}
}
