package main

import (
	"fmt"
)

func main() {
	var number1 = 5
	var number2 = 5
	sum := Sum(number1, number2)
	fmt.Println(number1, "+", number2, "=", sum)
}

func Sum(number1, number2 int) int {
	return number1 + number2
}
