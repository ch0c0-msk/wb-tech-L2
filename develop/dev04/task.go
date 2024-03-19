package dev04

import (
	"sort"
	"strings"
)

// FindAnagrams returns a set of anagrams list. Key is first occured word in list
func FindAnagrams(words []string) map[string][]string {
	res := make(map[string][]string)
	anagrams := make(map[string][]string)
	uniqueWords := make(map[string]bool)
	for _, w := range words {
		w = strings.ToLower(w)
		if !uniqueWords[w] {
			uniqueWords[w] = true
			sortedWord := sortWord(w)
			anagrams[sortedWord] = append(anagrams[sortedWord], w)
		}
	}
	for _, words := range anagrams {
		if len(words) > 1 {
			key := words[0]
			sort.Strings(words)
			res[key] = words
		}
	}
	return res
}

func sortWord(word string) string {
	res := []rune(word)
	sort.Slice(res, func(i, j int) bool { return res[i] < res[j] })
	return string(res)
}
