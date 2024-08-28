package main

import (
	"fmt"
	"goTraining/simplecalc"
)

// go get -u github.com/go-sql-driver/mysql

func main() {

	//ASSIGNMENT 1

	n := 5
	fmt.Println("using For loop")
	simplecalc.EvenOddNumbersFor(n)

	fmt.Println("using Switch case")
	simplecalc.EvenOddNumbersSwicth(n)

	// GO ROUTINES AND CHANNELS
	//PRIME NUMBER GENERATOR
	primes := simplecalc.GeneratePrimes(n)
	fmt.Printf("Prime numbers up to %d: %v\n", n, primes)

	//PARLLEL SUM CALCULATION
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	numGoroutines := 4
	totalSum := simplecalc.ParallelSum(numbers, numGoroutines)
	fmt.Printf("Total sumof array %v: %d\n", numbers, totalSum)

}
