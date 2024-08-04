package main

func maxMessages(thresh int) int {
	var total int = 0
	var numMessages int = 0
	for {
		messagesCost := 100 + numMessages
		if total+messagesCost > thresh {
			break
		}
		total += messagesCost
		numMessages++
	}

	return numMessages
}
