package day2part2

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

type Counts struct {
	red   int
	green int
	blue  int
}

func parseCount(str string, counts *Counts) {
	trimmed := strings.Trim(str, "\n\t\r ")
	parts := strings.Split(trimmed, " ")
	if len(parts) != 2 {
		println(str)
		panic("Must Include 2 parts")
	}
	count, _ := strconv.Atoi(parts[0])
	color := parts[1]
	if color == "red" && count > counts.red {
		counts.red = count
	}
	if color == "blue" && count > counts.blue {
		counts.blue = count
	}
	if color == "green" && count > counts.green {
		counts.green = count
	}
}

// hande a line like below
// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func multLines(str string) int {
	strs := strings.Split(str, ":")
	scoreStrs := strs[len(strs)-1]
	splitFunc := func(c rune) bool {
		return c == ',' || c == ';'
	}
	count := Counts{}
	pullStrs := strings.FieldsFunc(scoreStrs, splitFunc)
	for _, pullStr := range pullStrs {
		parseCount(pullStr, &count)
	}
	return count.red * count.green * count.blue
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

func main() {
	lines := readUntilEmpty()
	sum := 0
	for _, line := range lines {
		sum += multLines(line)
	}
	println(sum)
}
