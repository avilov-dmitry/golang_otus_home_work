package hw03frequencyanalysis

import (
	"fmt"
	"sort"
	"strings"
)

type CountWordsCalbackType func(string) map[string]int

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

func Top10(enter string) []string {
	if len(enter) == 0 {
		result := make([]string, 0)
		return result
	}
	countWords := getCounter()
	enterArr := strings.Fields(enter)

	for _, word := range enterArr {
		countWords(word)
	}

	m := countWords("")

	type kv struct {
		Key   string
		Value int
	}

	var ss = make([]kv, 0)
	var result []string

	for k, v := range m {
		ss = append(ss, kv{k, v})
	}

	sort.Slice(ss, func(i, j int) bool {
		if ss[i].Value == ss[j].Value {
			return ss[i].Key < ss[j].Key
		}
		return ss[i].Value > ss[j].Value
	})

	maxLength := len(ss)

	if maxLength > 10 {
		maxLength = 10
	}

	for i := 0; i < maxLength; i++ {
		kv := ss[i]
		fmt.Printf("%s %d\n", kv.Key, kv.Value)
		result = append(result, kv.Key)
	}

	return result
}
