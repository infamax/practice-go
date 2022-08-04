package secretmessage

import "sort"

// Делаем массив встречаемости символов
func createMapSymbols(str string) ([]rune, map[rune]int) {
	runes := []rune(str)
	res := make(map[rune]int)
	for _, v := range runes {
		res[v]++
	}
	return runes, res
}

type PairCharAndFrequently struct {
	char       rune
	frequently int
}

// Decode func
func Decode(encoded string) string {
	runes, m := createMapSymbols(encoded)
	frequentlyChars := make([]PairCharAndFrequently, 0, len(m))
	for key, value := range m {
		frequentlyChars = append(frequentlyChars, PairCharAndFrequently{key, value})
	}

	sort.Slice(frequentlyChars, func(i, j int) bool {
		if frequentlyChars[i].frequently == frequentlyChars[j].frequently {
			return frequentlyChars[i].char < frequentlyChars[j].char
		}
		return frequentlyChars[i].frequently > frequentlyChars[j].frequently
	})

	runes = nil
	runes = make([]rune, len(frequentlyChars), len(frequentlyChars))
	for i, _ := range frequentlyChars {
		runes[i] = frequentlyChars[i].char
	}
	i := 0
	for string(runes[i]) != "_" {
		i++
	}

	return string(runes[:i])
}
