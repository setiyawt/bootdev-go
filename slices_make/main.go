package main

func getMessageCosts(messages []string) []float64 {
	messagesCost := make([]float64, len(messages)) // Preallocate a slice for the message costs of the same length as the messages slice
	for i := 0; i < len(messages); i++ {           // Fill the costs slice with costs for each message.
		message := messages[i]                        // the message in the messages slice at the same index
		messagesCosts := float64(len(message)) * 0.01 // The cost of a message is the length of the message multiplied by 0.01.
		messagesCost[i] = messagesCosts               // The cost in the cost slice should correspond to the message in the messages slice at the same index

	}
	return messagesCost
}
