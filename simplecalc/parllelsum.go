package simplecalc

func SumPart(numbers []int, ch chan int) {
	sum := 0
	for _, num := range numbers {
		sum += num
	}
	ch <- sum
}

func ParallelSum(numbers []int, numGoroutines int) int {
	length := len(numbers)
	ch := make(chan int, numGoroutines)

	partSize := length / numGoroutines

	for i := 0; i < numGoroutines; i++ {
		start := i * partSize
		end := start + partSize

		if i == numGoroutines-1 {
			end = length
		}

		go SumPart(numbers[start:end], ch)
	}

	totalSum := 0
	for i := 0; i < numGoroutines; i++ {
		totalSum += <-ch
	}

	return totalSum
}
