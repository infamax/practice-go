package reverseparentheses

func reverseSlicesRunes(runes []rune) {
	for i := 0; i < len(runes)/2; i++ {
		runes[i], runes[len(runes)-i-1] = runes[len(runes)-i-1], runes[i]
	}
}

func Reverse(s string) string {
	res := make([]rune, 0, len(s))
	runes := []rune(s)
	cnt := 0
	tmp := make([]rune, 0)
	for _, ch := range runes {
		if ch == '(' {
			cnt++
		}

		if ch == ')' {
			cnt--
		}

		if cnt%2 == 0 {
			if cnt == 0 && len(tmp) != 0 {
				reverseSlicesRunes(tmp)
				res = append(res, tmp...)
				tmp = nil
			}
			if ch != '(' && ch != ')' {
				res = append(res, ch)
			}
		} else if ch != '(' && ch != ')' {
			tmp = append(tmp, ch)
		}
	}

	if cnt == 0 && len(tmp) != 0 {
		reverseSlicesRunes(tmp)
		res = append(res, tmp...)
		tmp = nil
	}

	return string(res)
}
