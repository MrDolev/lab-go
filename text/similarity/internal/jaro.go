package internal

import (
	"similarity/utils"
)

func Similarity(firstStr, secondStr string) float64 {
	if firstStr == secondStr {
		return 1.0
	}
	res := apply(utils.ToRune(firstStr), utils.ToRune(secondStr))
	return res
}

func Distance(firstStr, secondStr string) float64 {
	if firstStr == "" && secondStr == "" {
		return 0.0
	}
	if firstStr == "" || secondStr == "" {
		return 1.0
	}
	res := apply(utils.ToRune(firstStr), utils.ToRune(secondStr))
	return 1 - res
}

func apply(firstRune []rune, secondRune []rune) float64 {
	triple := getTriple(firstRune, secondRune)
	if triple.matches == 0 {
		return 0.0
	}
	res := calculate(triple, len(firstRune), len(secondRune))
	if res < 0.7 {
		return float64(res)
	}
	return approximation(float64(res), triple)
}

func calculate(triple Triple, firstLen, secondLen int) float64 {
	firstRatio := float64(triple.matches) / float64(firstLen)
	secondRatio := float64(triple.matches) / float64(secondLen)
	transposeRatio := float64(float64(triple.matches)-(float64(triple.traspose/2.0))) / float64(triple.matches)
	return float64(firstRatio+secondRatio+transposeRatio) / 3.0
}

func approximation(partialRes float64, triple Triple) float64 {
	return partialRes + (0.1 * float64(triple.prefix) * float64(1.0-partialRes))
}

type Triple struct {
	matches  int
	traspose int
	prefix   int
}

func getTriple(firstRune []rune, secondRune []rune) Triple {
	var minRune, maxRune []rune
	minRune, maxRune = getMinMaxRune(firstRune, secondRune)
	minLen, maxLen := len(minRune), len(maxRune)
	rangeLen := utils.MaxInt((maxLen/2)-1, 0)

	matchedByMin := make([]bool, minLen)
	matchedByMax := make([]bool, maxLen)

	matched := 0

	for index := 0; index < minLen; index++ {
		currentChar := minRune[index]
		found := false
		lower := utils.MaxInt(index-rangeLen, 0)
		upper := utils.MinInt(index+rangeLen+1, maxLen)
		for lower < upper && !found {
			if !matchedByMax[lower] && currentChar == maxRune[lower] {
				matchedByMin[index] = true
				matchedByMax[lower] = true
				matched++
				found = true
			}
			lower++
		}
	}

	firstMatched := fillMatched(minRune, matchedByMin, make([]rune, matched))
	secondMatched := fillMatched(maxRune, matchedByMax, make([]rune, matched))

	transposition := getTransposition(firstMatched, secondMatched)
	prefix := getPrefix(minRune, firstRune, secondRune)

	return Triple{matched, transposition, prefix}
}

func getMinMaxRune(firstRune []rune, secondRune []rune) ([]rune, []rune) {
	if len(firstRune) > len(secondRune) {
		return secondRune, firstRune
	} else {
		return firstRune, secondRune
	}
}

func getPrefix(min []rune, firstRune []rune, secondRune []rune) int {
	prefix := 0
	found := false

	for i := 0; i < utils.MinInt(4, len(min)) && !found; i++ {
		if firstRune[i] != secondRune[i] {
			found = true
		}
		prefix++
	}
	return prefix
}

func getTransposition(firstMatched []rune, secondMatched []rune) int {
	transposition := 0
	for i := 0; i < len(firstMatched); i++ {
		if firstMatched[i] != secondMatched[i] {
			transposition++
		}
	}
	return transposition
}

func fillMatched(src []rune, matchedFlags []bool, dst []rune) []rune {
	for i, j := 0, 0; i < len(src); i++ {
		if matchedFlags[i] {
			dst[j] = src[i]
			j++
		}
	}
	return dst
}
