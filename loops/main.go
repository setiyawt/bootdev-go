package main

func bulkSend(numMessages int) float64 {
	var total float64 = 0.0
	for i := 0; i < numMessages; i++ {
		total += 1 + float64(i)/100

	}

	return total
}
