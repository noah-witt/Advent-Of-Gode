package scratchcards

import (
	"math"
	"strconv"
	"strings"
)

func numInArray(num int, list []int) bool {
	for _, b := range list {
		if b == num {
			return true
		}
	}
	return false
}

// e.g. Input 41 48 83 86 17
// e.g. result [41, 48, 83, 86, 17]
func getNumsFromStr(section string) []int {
	parts := strings.Split(section, " ")
	result := []int{}
	for _, val := range parts {
		val = strings.Trim(val, " ")
		if len(val) == 0 {
			continue
		}
		num, _ := strconv.Atoi(val)
		result = append(result, num)
	}
	return result
}

// e.g. Input 41 48 83 86 17 | 83 86  6 31 17  9 48 53
func getWinnerCount(section string) int {
	parts := strings.Split(section, "|")
	if len(parts) != 2 {
		panic("MUST HAVE 2 parts")
	}
	winningNums := getNumsFromStr(strings.Trim(parts[0], " "))
	haveNums := getNumsFromStr(strings.Trim(parts[1], " "))
	winningNumsHave := []int{}
	for _, num := range haveNums {
		if numInArray(num, winningNums) {
			winningNumsHave = append(winningNumsHave, num)
		}
	}
	return len(winningNumsHave)
}

func calcScore(correctCount int) int {
	if correctCount == 0 {
		return 0
	}
	if correctCount == 1 {
		return 1
	}
	return int(math.Pow(float64(2), float64(correctCount-1)))
}

func scoreLine(line string) int {
	parts := strings.Split(line, ":")
	winCount := getWinnerCount(parts[1])
	return calcScore(winCount)
}

func sumCards(lines []string) int {
	accumulator := 0
	for _, line := range lines {
		accumulator += scoreLine(line)
	}
	return accumulator
}
