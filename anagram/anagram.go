package anagram

import (
	"sort"
	"strings"
)

func makeSortedString(str string) string {
	res := str
	res = strings.ToLower(res)
	res = strings.ReplaceAll(res, " ", "")
	res = strings.Trim(res, "\r")
	runes := []rune(res)
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})
	return string(runes)
}

func FindAnagrams(dictionary []string, word string) (result []string) {
	if len(dictionary) == 0 {
		return result
	}

	if len(word) == 0 {
		return result
	}

	searchWord := makeSortedString(word)
	word = strings.Trim(word, "\r")
	for _, str := range dictionary {
		tmp := makeSortedString(str)
		str = strings.Trim(str, "\r")
		if strings.EqualFold(tmp, searchWord) && !strings.EqualFold(str, word) {
			result = append(result, str)
		}
	}
	return result
}
