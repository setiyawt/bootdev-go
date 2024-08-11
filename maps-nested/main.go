package main

func getNameCounts(names []string) map[rune]map[string]int {
	counts := make(map[rune]map[string]int)
	for _, name := range names {
		if name == "" {
			continue
		}
		initial := []rune(name)[0]
		if _, exists := counts[initial]; !exists {
			counts[initial] = make(map[string]int)
		}
		counts[initial][name]++
	}
	return counts
}
