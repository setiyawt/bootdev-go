package main

func indexOfFirstBadWord(msg []string, badWords []string) int {
	for i, word := range msg {
		for _, badword := range badWords {
			if word == badword {
				return i
			}
		}
	}
	return -1
}
