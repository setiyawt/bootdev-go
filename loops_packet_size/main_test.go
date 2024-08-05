package main

import (
	"fmt"
	"testing"
)

func TestGetPacketSize(t *testing.T) {
	type testCase struct {
		message            string
		expectedPacketSize int
	}
	testCases := []testCase{
		{"No, this is Boots the Magic Bear", 8},
		{"I've been trying to reach you about your car's extended warranty", 16},
		{"The mystery of salmon isn't a problem to solve, but a reality to experience...", 13},
	}
	if withSubmit {
		testCases = append(testCases,
			[]testCase{
				{"", 0},
				{"Hello there", 0},
				{"I've got a bad conditional statement about this.", 12},
				{"It's over, Ballan, I have the high ground. Don't try it!", 14},
			}...,
		)
	}

	passCount := 0
	failCount := 0

	for _, tc := range testCases {
		result := getPacketSize(tc.message)
		if result != tc.expectedPacketSize {
			failCount++
			t.Errorf(`---------------------------------
Message:   "%s"
Expecting: %v
Actual:    %v
Fail`, tc.message, tc.expectedPacketSize, result)
		} else {
			passCount++
			fmt.Printf(`---------------------------------
Message:   "%s"
Expecting: %v
Actual:    %v
Pass
`, tc.message, tc.expectedPacketSize, result)
		}
	}

	fmt.Println("---------------------------------")
	fmt.Printf("%d passed, %d failed\n", passCount, failCount)
}

var withSubmit = true
