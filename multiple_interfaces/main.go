package main

import "fmt"

func (e email) cost() int {
	if e.isSubscribed == false {
		cost := 5 * len(e.body)
		return cost
	}
	cost := 2 * len(e.body)
	return cost

}

func (e email) format() string {
	if e.isSubscribed == false {
		return fmt.Sprintf("'%s' | Not Subscribed", e.body)
	}
	return fmt.Sprintf("'%s' | Subscribed", e.body)

}

type expense interface {
	cost() int
}

type formatter interface {
	format() string
}

type email struct {
	isSubscribed bool
	body         string
}
