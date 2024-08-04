package main

import (
	"fmt"
)

func printPrimes(max int) {
	for i := 2; i <= max; i++ {
		if i == 2 {
			fmt.Printf("%d is prime\n", i)
			continue
		}
		if i%2 == 0 {
			fmt.Printf("%d is not prime\n", i)
			continue
		}

		for j := 3; j*j <= i; j += 2 {
			if i%j == 0 {
				fmt.Printf("%d is not prime\n", i)

				break
			}
		}
		fmt.Printf("%d is prime\n", i)
	}
}

// don't edit below this line

func test(max int) {
	fmt.Printf("Primes up to %v:\n", max)
	printPrimes(max)
	fmt.Println("===============================================================")
}

func main() {
	test(10)
	test(20)
	test(30)
}
