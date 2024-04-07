package planting

import (
	"errors"
	"math"
	"strconv"
	"strings"
)

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

func listMin(arr []int) int {
	min := math.MaxInt
	for _, value := range arr {
		if value < min {
			min = value
		}
	}
	return min
}

func stringOnlyNumberAndSpace(str string) bool {
	for _, char := range str {
		if !((char >= '0' && char <= '9') || char == ' ') {
			return false
		}
	}
	return true
}

type MapItem struct {
	source      int
	destination int
	rangeSize   int
}

func (v MapItem) inRange(input int) bool {
	if input < v.source {
		return false
	}
	if input >= v.source+v.rangeSize {
		return false
	}
	return true
}

func (v MapItem) apply(input int) int {
	diff := v.source - v.destination
	return input - diff
}

// input: 49 53 8
// format destination source range
func newMapItem(row string) MapItem {
	numbers := getNumsFromStr(row)
	if len(numbers) != 3 {
		panic("Each Observation must have 3 members")
	}
	mapItem := MapItem{}
	mapItem.destination = numbers[0]
	mapItem.source = numbers[1]
	mapItem.rangeSize = numbers[2]
	return mapItem
}

type Map struct {
	items []MapItem
}

func (v Map) apply(input int) (int, error) {
	for _, item := range v.items {
		if item.inRange(input) {
			return item.apply(input), nil
		}
	}
	return 0, errors.New("unable to apply")
}

/*
*
parse format like
light-to-temperature map:
45 77 23
81 45 19
68 64 13
*/
func newMap(sections []string) Map {
	newMap := Map{}
	for _, section := range sections {
		trimmedSection := strings.Trim(section, " ")
		if len(trimmedSection) > 0 && stringOnlyNumberAndSpace(section) {
			newMap.items = append(newMap.items, newMapItem(section))
		}
	}
	return newMap
}

type Almanac struct {
	maps []Map
}

func (v Almanac) apply(input int) int {
	at := input
	for _, item := range v.maps {
		nowAt, e := item.apply(at)
		if e != nil {
			// any not mapped is pass through
			nowAt = at
		}
		at = nowAt
	}
	return at
}

func newAlmanac(sections []string) Almanac {
	newAlmanac := Almanac{}
	blankLines := []int{}
	for index, item := range sections {
		// find indexes of blank lines
		if len(strings.Trim(item, " ")) == 0 {
			blankLines = append(blankLines, index)
		}
	}
	// imaginary blank line at end
	blankLines = append(blankLines, len(sections))
	for listIndex, sectionEnd := range blankLines {
		if listIndex == 0 {
			continue
		}
		sectionStart := blankLines[listIndex-1] + 1
		newAlmanac.maps = append(newAlmanac.maps, newMap(sections[sectionStart:sectionEnd]))
	}
	return newAlmanac
}

func processInput(lines []string) []int {
	seeds := lines[0]
	seedNumbers := getNumsFromStr(strings.Split(seeds, ":")[1])
	almanac := newAlmanac(lines[1:])
	results := []int{}
	for _, seedNumber := range seedNumbers {
		results = append(results, almanac.apply(seedNumber))
	}
	return results
}

func minSeedMap(lines []string) int {
	results := processInput(lines)
	return listMin(results)
}
