package ratio

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func isSymbol(c byte) bool {
	return c != '.' && (unicode.IsSymbol(rune(c)) || unicode.IsPunct(rune(c)))
}

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

func getNumberFromPtrs(characters []byte, pointer int, endPointer int) string {
	result := ""
	for pointer < endPointer {
		result += string(characters[pointer])
		pointer++
	}
	return result
}

func isAdjacentToSymbol(characters []byte, pointer int) bool {
	println("checking if ", string(characters[pointer]), " is adjacent to symbol. at pointer:", pointer)
	// if a pointer is proceeded or followed by a symbol it touches a symbol
	if pointer-1 >= 0 && isSymbol(characters[pointer-1]) {
		return true
	}
	if pointer+1 < len(characters) && isSymbol(characters[pointer+1]) {
		return true
	}
	// now find how far from the start of the line we are
	prevNewLinePointer := prevNewline(characters, pointer)
	charsFromStartOfLine := pointer - prevNewLinePointer
	if prevNewLinePointer > 0 {
		// pointer is not pointing to first line so we can go back home more newline
		startOfPrevLine := prevNewline(characters, prevNewLinePointer-1)
		pointerAbove := startOfPrevLine + charsFromStartOfLine
		// check line above 1 row
		// println("Checking", string(characters[pointerAbove]), " ", string(characters[pointerAbove+1]), " and ", string(characters[pointerAbove-1]))
		if isSymbol(characters[pointerAbove]) {
			return true
		}
		// check one over for diagonals
		if isSymbol(characters[pointerAbove+1]) {
			return true
		}
		// check one over for diagonals
		if isSymbol(characters[pointerAbove-1]) {
			return true
		}
	}
	startNextLine := nextNewline(characters, pointer)
	pointerBelow := startNextLine + charsFromStartOfLine
	// line up with line below
	if startNextLine > 0 {
		// check below
		if pointerBelow < len(characters) && isSymbol(characters[pointerBelow]) {
			return true
		}
		// check below -1 for diagonals
		if pointerBelow-1 < len(characters) && isSymbol(characters[pointerBelow-1]) {
			return true
		}
		// check below +1 for diagonals
		if pointerBelow+1 < len(characters) && isSymbol(characters[pointerBelow+1]) {
			return true
		}
	}
	return false
}

func rangeIsAdjacentToSymbol(characters []byte, pointer int, endPointer int) bool {
	println("checking range pointers: ", pointer, " ", endPointer, " string: ", string(characters[pointer:endPointer]))
	startPointer := pointer
	for pointer < endPointer {
		result := isAdjacentToSymbol(characters, pointer)
		if result {
			return true
		}
		pointer++
	}
	println(string(characters[startPointer:endPointer]), "IS NOT ADJACENT TO A CHARACTER")
	return false
}

// get the full num from a pointer
// figure out if it is adjacent to a symbol
func processNum(characters []byte, pointer int) (int, int) {
	endPtr := findEndPtr(characters, pointer)
	// number is ptr...EndPtr inclusive
	numStr := getNumberFromPtrs(characters, pointer, endPtr)
	num, _ := strconv.Atoi(numStr)
	result := rangeIsAdjacentToSymbol(characters, pointer, endPtr)
	if result {
		println("Number ", num, " is adjacent to a symbol. from str: ", numStr)
		return num, endPtr
	}
	return 0, endPtr
}

func getNums(characters []byte) int {
	pointer := 0
	accumulator := 0
	for pointer < len(characters) {
		// Iterate until hit number
		if !isNumber(characters[pointer]) {
			pointer++
			continue
		}
		// pointer is a num
		result, endPtr := processNum(characters, pointer)
		pointer = endPtr
		accumulator += result
		pointer++
	}
	return accumulator
}

func readUntilEmpty() []string {
	strs := []string{}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	for len(text) > 0 {
		trimmed := strings.Trim(text, "\n\t\r ")
		if len(trimmed) == 0 {
			break
		}
		strs = append(strs, trimmed)
		text, _ = reader.ReadString('\n')
	}
	return strs
}

func convertStr(strs []string) string {
	result := ""
	for _, str := range strs {
		result += str + "\n"
	}
	return result
}

func main() {
	strs := readUntilEmpty()
	str := convertStr(strs)
	bytes := []byte(str)
	println(getNums(bytes))
}
