package hw03frequencyanalysis

import (
	"sort"
	"strings"
)

type CountWordsCalbackType func(string) map[string]int

type kv struct {
	Key   string
	Value int
}

func getCounter() CountWordsCalbackType {
	result := map[string]int{}

	return func(word string) map[string]int {
		if word != "" {
			if result[word] > 0 {
				result[word]++
			} else {
				result[word] = 1
			}
		}

		return result
	}
}

func sortCountedMap(countedMap map[string]int) (sorted []kv) {
	sorted = make([]kv, 0, len(countedMap))

	for k, v := range countedMap {
		sorted = append(sorted, kv{k, v})
	}

	sort.Slice(sorted, func(i, j int) bool {
		if sorted[i].Value == sorted[j].Value {
			return sorted[i].Key < sorted[j].Key
		}
		return sorted[i].Value > sorted[j].Value
	})

	return sorted
}

func Top10(enter string) []string {
	if len(enter) == 0 {
		return make([]string, 0)
	}

	count := getCounter()
	var countedMap map[string]int

	for _, word := range strings.Fields(enter) {
		countedMap = count(word)
	}

	sorted := sortCountedMap(countedMap)
	maxLength := len(sorted)

	if maxLength > 10 {
		maxLength = 10
	}

	var result []string

	for i := 0; i < maxLength; i++ {
		result = append(result, sorted[i].Key)
	}

	return result
}
