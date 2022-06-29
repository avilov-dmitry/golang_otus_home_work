package hw03frequencyanalysis

import (
	"fmt"
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

func sortCountedMap(countedMap map[string]int) (ss []kv) {
	ss = make([]kv, 0, len(countedMap))

	for k, v := range countedMap {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value == ss[j].Value {
			return ss[i].Key < ss[j].Key
		}
		return ss[i].Value > ss[j].Value
	})

	return ss
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
	var result []string

	maxLength := len(sorted)

	if maxLength > 10 {
		maxLength = 10
	}

	for i := 0; i < maxLength; i++ {
		kv := sorted[i]
		fmt.Printf("%s %d\n", kv.Key, kv.Value)
		result = append(result, kv.Key)
	}

	return result
}
