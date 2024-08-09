package main

import (
	"strings"
)

type sms struct {
	id      string
	content string
	tags    []string
}

func tagMessages(messages []sms, tagger func(sms) []string) []sms {
	for i, message := range messages {
		messages[i].tags = tagger(message)
	}
	return messages
}

func tagger(msg sms) []string {
	tags := []string{}
	if containsWord(msg.content, "urgent") || containsWord(msg.content, "Urgent") {
		tags = append(tags, "Urgent")
	}
	if containsWord(msg.content, "sale") || containsWord(msg.content, "Sale") {
		tags = append(tags, "Promo")
	}
	return tags
}

// Helper function to check if content contains a specific word
func containsWord(content, word string) bool {
	return strings.Contains(strings.ToLower(content), word)
}
