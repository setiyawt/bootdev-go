package main

import (
	"fmt"
)

func sendSMSToCouple(msgToCustomer, msgToSpouse string) (int, error) {
	sentCostomer, err := sendSMS(msgToCustomer)
	if err != nil {
		return 0, err
	} else {
		sentSpouse, err := sendSMS(msgToSpouse)
		if err != nil {
			return 0, err
		}
		return sentCostomer + sentSpouse, nil
	}
}

// don't edit below this line

func sendSMS(message string) (int, error) {
	const maxTextLen = 25
	const costPerChar = 2
	if len(message) > maxTextLen {
		return 0, fmt.Errorf("can't send texts over %v characters", maxTextLen)
	}
	return costPerChar * len(message), nil
}
