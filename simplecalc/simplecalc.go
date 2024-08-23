package simplecalc

import "fmt"

func EvenOddNumbersFor(n int) {
	for i := 1; i < n; i++ {
		if i%2 == 0 {
			fmt.Println("even:", i)
		} else {
			fmt.Println("odd:", i)
		}
	}

}

func EvenOddNumbersSwicth(n int) {
	for i := 1; i < n; i++ {
		switch i % 2 {
		case 0:
			fmt.Println("even:", i)
		default:
			fmt.Println("odd:", i)
		}
	}
}
