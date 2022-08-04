package buildworld

import (
	"math"
	"sort"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func find(a []string, str string) bool {
	l := -1
	r := len(a)
	for r-l > 1 {
		m := (l + r) / 2
		if a[m] < str {
			l = m
		} else {
			r = m
		}
	}

	if r == len(a) {
		return false
	}

	if a[r] != str {
		return false
	}

	return true
}

func substr(str []rune, start, finish int) []rune {
	ans := make([]rune, finish-start+1, finish-start+1)
	for i := start; i <= finish; i++ {
		ans[i-start] = str[i]
	}
	return ans
}

const inf = math.MaxInt - 1

func BuildWord(word string, fragments []string) int {
	n := len(word)
	dp := make([]int, n+1)
	//dp[n] - какое количество слов нужно из словаря, чтобы построить слово длины n
	dp[0] = 0 // База динамики
	for i := 1; i < len(dp); i++ {
		dp[i] = inf
	}

	sort.Slice(fragments, func(i, j int) bool {
		return fragments[i] < fragments[j]
	})

	str := make([]rune, 0, n)

	// Cложность алгоритма O(n^2 log(n))
	for i := 1; i < len(dp); i++ {
		str = append(str, rune(word[i-1]))
		// Если данное слово длины n уже есть в словаре, очевидно что dp[n] = 1
		if find(fragments, string(str)) {
			dp[i] = 1
		} else {
			// Иначе мы должны проверить все предыдущие dp[i] != inf и узнать можно ли из букв получить слово длины n
			for j := 1; j < i; j++ {
				if dp[j] != inf {
					if find(fragments, string(substr(str, j, i-1))) {
						dp[i] = min(dp[i], dp[j]+1)
					}
				}
			}
		}
	}

	if dp[n] == inf {
		return 0
	}

	return dp[n]
}
