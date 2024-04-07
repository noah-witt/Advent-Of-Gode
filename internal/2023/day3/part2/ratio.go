package ratio

import "strconv"

func isNumber(c byte) bool {
	return c >= '0' && c <= '9'
}

func isNewline(c byte) bool {
	return c == '\n'
}

func nextNewline(characters []byte, pointer int) int {
	for pointer < len(characters) {
		if isNewline(characters[pointer]) {
			return pointer
		}
		pointer++
	}
	return -1
}

func prevNewline(characters []byte, pointer int) int {
	for pointer >= 0 {
		if isNewline(characters[pointer]) {
			return pointer
		}
		pointer--
	}
	return -1
}

func findEndPtr(characters []byte, pointer int) int {
	endPtr := pointer
	for endPtr < len(characters) && isNumber(characters[endPtr]) {
		endPtr++
	}
	return endPtr
}

func getNumberStrFromPtrs(characters []byte, pointer int, endPointer int) string {
	return string(characters[pointer : endPointer+1])
}

func nextStar(characters []byte, pointer int) int {
	for pointer < len(characters) {
		if characters[pointer] == '*' {
			return pointer
		}
		pointer++
	}
	return -1
}

func safeAppendIfDigit(characters []byte, tryPointer int, result *[]int) {
	if tryPointer < 0 {
		return
	}
	if tryPointer >= len(characters) {
		return
	}
	if isNumber(characters[tryPointer]) {
		*result = append(*result, tryPointer)
	}
}

// get the index of adjacent digits
func starAdjacentDigitPointers(characters []byte, starPointer int) []int {
	result := []int{}
	// before
	safeAppendIfDigit(characters, starPointer-1, &result)
	// after
	safeAppendIfDigit(characters, starPointer+1, &result)
	// now find how far from the start of the line we are
	prevNewLinePointer := prevNewline(characters, starPointer)
	charsFromStartOfLine := starPointer - prevNewLinePointer
	if prevNewLinePointer > 0 {
		// pointer is not pointing to first line so we can go back home more newline
		startOfPrevLine := prevNewline(characters, prevNewLinePointer-1)
		pointerAbove := startOfPrevLine + charsFromStartOfLine
		// order is important because we assume ascending in condenseAdjacentPointers
		safeAppendIfDigit(characters, pointerAbove-1, &result)
		safeAppendIfDigit(characters, pointerAbove, &result)
		safeAppendIfDigit(characters, pointerAbove+1, &result)
	}
	startNextLine := nextNewline(characters, starPointer)
	pointerBelow := startNextLine + charsFromStartOfLine
	// line up with line below
	if startNextLine > 0 {
		// check below
		// order is important because we assume ascending in condenseAdjacentPointers
		safeAppendIfDigit(characters, pointerBelow-1, &result)
		safeAppendIfDigit(characters, pointerBelow, &result)
		safeAppendIfDigit(characters, pointerBelow+1, &result)
	}
	return result
}

// condense the adjacent pointers that are sequential
func condenseAdjacentPointers(pointers []int) []int {
	result := []int{}
	skipped := -2 // start at neg 2 so skipped+1 can never equal value
	for _, value := range pointers {
		if len(result) == 0 {
			result = append(result, value)
			continue
		}
		if skipped+1 == value {
			// skip again
			skipped = value
			continue
		}
		if result[len(result)-1]+1 == value {
			skipped = value
			continue
		}
		result = append(result, value)
	}
	return result
}

func rangeAdjacentToDigit(characters []byte, pointer int) (int, int) {
	first := pointer
	for first > 0 && isNumber(characters[first-1]) {
		first--
	}
	last := pointer
	for last+1 < len(characters) && isNumber(characters[last+1]) {
		last++
	}
	return first, last
}

func process(characters []byte) int {
	accumulator := 0
	pointer := nextStar(characters, 0)
	for pointer >= 0 {
		// pointing at star find adjacent digits
		adjacentDigitPointers := starAdjacentDigitPointers(characters, pointer)
		adjacentDigitPointers = condenseAdjacentPointers(adjacentDigitPointers)
		if len(adjacentDigitPointers) == 2 {
			println("Star at ", pointer, "is adjacent to 2 nums")
			// now step backward and forward on both to construct ranges
			aStart, aEnd := rangeAdjacentToDigit(characters, adjacentDigitPointers[0])
			aNumStr := getNumberStrFromPtrs(characters, aStart, aEnd)
			aNum, _ := strconv.Atoi(aNumStr)
			bStart, bEnd := rangeAdjacentToDigit(characters, adjacentDigitPointers[1])
			bNumStr := getNumberStrFromPtrs(characters, bStart, bEnd)
			bNum, _ := strconv.Atoi(bNumStr)
			println("Gar Values ", aNum, " ", bNum)
			accumulator += aNum * bNum
		} else {
			println("Star at ", pointer, "is adjacent to ", len(adjacentDigitPointers), " not 2.")
		}
		// iterate
		pointer = nextStar(characters, pointer+1)
	}
	return accumulator
}
