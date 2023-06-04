package utils

func MaxInt(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func MinInt(x, y int) int {
	if x > y {
		return y
	}
	return x
}

func MinFloat(x, y float64) float64 {
	if x > y {
		return y
	}
	return x
}

func ToRune(str string) []rune {
	return []rune(str)
}
