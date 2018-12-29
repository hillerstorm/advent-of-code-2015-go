package day01

func traverseInput(input string) (int, int) {
	floor := 0
	basement := -1
	for i, x := range input {
		switch x {
		case '(':
			floor++
		case ')':
			floor--
		}
		if basement == -1 && floor == -1 {
			basement = i + 1
		}
	}
	return floor, basement
}
