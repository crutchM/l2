package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	input := []string{"тест", "листок", "пятка", "пятак", "тяпка", "листок", "пятка", "слиток", "столик"}
	fmt.Println(input)
	fmt.Println(findAnagrams(&input))
}

func findAnagrams(input *[]string) map[string][]string {
	if len(*input) < 2 {
		return nil
	}
	buffer := make(map[string][]string)
	result := make(map[string][]string)
	processedWords := findUnique(input)
	for _, value := range processedWords {
		sorted := []rune(value)
		sort.Slice(sorted, func(i, j int) bool {
			return sorted[i] < sorted[j]
		})
		word := string(sorted)
		buffer[word] = append(buffer[word], value)
	}
	for _, value := range buffer {
		if len(value) > 1 {
			sort.Strings(value)
			result[value[0]] = value
		}
	}
	return result
}

func findUnique(input *[]string) []string {
	result := make([]string, 0, len(*input))
	unique := make(map[string]bool)
	for _, value := range *input {
		if !(unique[value]) {
			result = append(result, strings.ToLower(value))
			unique[value] = true
		}
	}
	return result
}
