package main

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

func getMaxCubes() Counts {
	maxCubes := Counts{}
	maxCubes.red = 12
	maxCubes.green = 13
	maxCubes.blue = 14
	return maxCubes
}

func parseAndVerifySingularCount(str string) bool {
	counts := Counts{}
	trimmed := strings.Trim(str, "\n\t\r ")
	parts := strings.Split(trimmed, " ")
	if len(parts) != 2 {
		println(str)
		println("Must Include 2 parts")
		return true
	}
	count, _ := strconv.Atoi(parts[0])
	color := parts[1]
	if color == "red" {
		counts.red += count
	}
	if color == "blue" {
		counts.blue += count
	}
	if color == "green" {
		counts.green += count
	}
	max := getMaxCubes()
	if counts.red > max.red || counts.green > max.green || counts.blue > max.blue {
		return false
	}
	return true
}

// hande a line like below
// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
func verifyLine(str string) bool {
	strs := strings.Split(str, ":")
	scoreStrs := strs[len(strs)-1]
	splitFunc := func(c rune) bool {
		return c == ',' || c == ';'
	}
	pullStrs := strings.FieldsFunc(scoreStrs, splitFunc)
	for _, pullStr := range pullStrs {
		result := parseAndVerifySingularCount(pullStr)
		if !result {
			return false
		}
	}
	return true
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
	workCount := 0
	for index, line := range lines {
		result := verifyLine(line)
		if result {
			workCount += index + 1
		} else {
			println("Fail ", index+1)
		}
	}
	println(workCount)
}
