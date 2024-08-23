package main

import (
	"fmt"
	"goTraining/simplecalc"
)

func main() {
	fmt.Println("hello world!")
	n := 10
	fmt.Println("using For loop")
	simplecalc.EvenOddNumbersFor(n)

	fmt.Println("using Switch case")
	simplecalc.EvenOddNumbersSwicth(n)

}
