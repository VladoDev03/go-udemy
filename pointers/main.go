package main

import "fmt"

func main() {
	age := 21

	// var agePointer *int
	agePointer := &age

	fmt.Println("Age:", age)
	fmt.Println("Age Dereferenced:", *agePointer)
	fmt.Println("Age Pointer:", agePointer)

	// adultYears := getAdultYears(age)
	adultYears := getAdultYearsPointer(agePointer)
	editAgeToAdultYears(agePointer)

	fmt.Println(adultYears)
	fmt.Println(age)
}

func getAdultYears(age int) int {
	return age - 18
}

func getAdultYearsPointer(age *int) int {
	return *age - 18
}

func editAgeToAdultYears(age *int) {
	*age = *age - 18
}
