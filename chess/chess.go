package chess

import (
	"errors"
	"strings"
)

func CanKnightAttack(white, black string) (bool, error) {
	if len(white) != 2 || len(black) != 2 || strings.EqualFold(white, black) {
		return false, errors.New("incorrect values")
	}
	var posX1, posY1, posX2, posY2 int

	posX1 = int(white[0]) - 'a' + 1
	posX2 = int(black[0]) - 'a' + 1
	posY1 = int(white[1]) - '0'
	posY2 = int(black[1]) - '0'

	if posX1 <= 0 || posX1 >= 9 || posX2 <= 0 || posX2 >= 9 ||
		posY1 <= 0 || posY1 >= 9 || posY2 <= 0 || posY2 >= 9 {
		return false, errors.New("incorrect positions")
	}

	// 8 position knights are possible

	if posX1+1 == posX2 && posY1+2 == posY2 {
		return true, nil
	}

	if posX1-1 == posX2 && posY1+2 == posY2 {
		return true, nil
	}

	if posX1+2 == posX2 && posY1+1 == posY2 {
		return true, nil
	}

	if posX1-2 == posX2 && posY1+1 == posY2 {
		return true, nil
	}

	if posX1+2 == posX2 && posY1-1 == posY2 {
		return true, nil
	}

	if posX1-2 == posX2 && posY1-1 == posY2 {
		return true, nil
	}

	if posX1-1 == posX2 && posY1-2 == posY2 {
		return true, nil
	}

	if posX1+1 == posX2 && posY1-2 == posY2 {
		return true, nil
	}

	return false, nil
}
