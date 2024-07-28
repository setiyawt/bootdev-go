package main

import "fmt"

type sender struct {
	rateLimit int
	user
}

type user struct {
	name   string
	number int
}

func main() {
	sender := sender{
		rateLimit: 10000,
		user: user{
			name:   "Deborah",
			number: 18055558790,
		},
	}
	fmt.Println("Name: ", sender.name)
	fmt.Println("Number: ", sender.number)
	fmt.Println("RateLimit: ", sender.rateLimit)
}
