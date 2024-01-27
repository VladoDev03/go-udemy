package main

import "fmt"

type customString string

func (text customString) log() {
	fmt.Println(text)
}

func main() {
	var text customString = "Some Text To Be Printed"
	text.log()
}
