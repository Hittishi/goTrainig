package simplecalc

func SendNumbers(ch chan int, n int) {
	for i := 2; i <= n; i++ {
		ch <- i
	}
	close(ch)
}

func FilterNumbers(in chan int, out chan int, prime int) {
	for i := range in {
		if i%prime != 0 {
			out <- i
		}
	}
	close(out)
}

func GeneratePrimes(n int) []int {
	primes := []int{}
	ch := make(chan int)
	go SendNumbers(ch, n)

	for {
		prime, ok := <-ch
		if !ok {
			break
		}
		primes = append(primes, prime)
		newCh := make(chan int)
		go FilterNumbers(ch, newCh, prime)
		ch = newCh
	}

	return primes
}
