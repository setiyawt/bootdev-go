package main

import "fmt"

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	messages := [3]string{primary, secondary, tertiary}
	costs := [3]int{getMessageSize(primary), getMessageSize(secondary), getMessageSize(tertiary)}

	costs[0] = getMessageSize(messages[0])
	costs[1] = getMessageSize(messages[1]) + costs[0]
	costs[2] = getMessageSize(messages[2]) + costs[1]
	for i := 0; i < 3; i++ {
		if messages[i] == "" {
			messages[i] = retryMessage(messages[i])
			if i == 0 {
				costs[i] = getMessageSize(messages[i])
			} else {
				costs[i] = getMessageSize(messages[i]) + costs[i-1]
			}
		}
	}

	return messages, costs
}

func retryMessage(message string) string {
	return message
}

func getMessageSize(message string) int {
	return len(message)
}

func main() {
	primary := "Hello sir/madam can I interest you in a yacht?"
	secondary := "Please I'll even give you an Amazon gift card?"
	tertiary := "You're missing out big time"

	messages, costs := getMessageWithRetries(primary, secondary, tertiary)

	fmt.Println("Messages:")
	for i, message := range messages {
		fmt.Printf("%d: %s\n", i+1, message)
	}

	fmt.Println("\nCosts:")
	for i, cost := range costs {
		fmt.Printf("%d: %d\n", i+1, cost)
	}
}
