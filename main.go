package main

import (
	"database/sql"
	"fmt"
	"goTraining/api"
	"log"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// go get -u github.com/go-sql-driver/mysql

func main() {

	//ASSIGNMENT 1

	// n := 5
	// fmt.Println("using For loop")
	// simplecalc.EvenOddNumbersFor(n)

	// fmt.Println("using Switch case")
	// simplecalc.EvenOddNumbersSwicth(n)

	// // GO ROUTINES AND CHANNELS
	// //PRIME NUMBER GENERATOR
	// primes := simplecalc.GeneratePrimes(n)
	// fmt.Printf("Prime numbers up to %d: %v\n", n, primes)

	// //PARLLEL SUM CALCULATION
	// numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// numGoroutines := 4
	// totalSum := simplecalc.ParallelSum(numbers, numGoroutines)
	// fmt.Printf("Total sumof array %v: %d\n", numbers, totalSum)

	dsn := "root:password@tcp(127.0.0.1:3306)/productsDB"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	api.RegisterRoutes(db)
	log.Println("server starting on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))

	fmt.Println("Successfully connected to the database on localhost!")

}
