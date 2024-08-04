package main

import "fmt"

func fizzbuzz() []string {
	var result []string
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			result = append(result, "FizzBuzz")
		} else if i%3 == 0 {
			result = append(result, "Fizz")
		} else if i%5 == 0 {
			result = append(result, "Buzz")
		} else {
			result = append(result, fmt.Sprintf("%d", i))
		}
	}
	return result
}

// don't touch below this line

func main() {
	fizzbuzz()
}
