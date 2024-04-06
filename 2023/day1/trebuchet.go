package main

import (
	"bufio"
	"os"
)

// Get array of ints from string
func getNumsFromStr(str string) []int {
	nums := []int{}
	for _, char := range str {
		if char >= '0' && char <= '9' {
			nums = append(nums, int(char-'0'))
		}
	}
	return nums
}

// a*10+b=return
// Concatenate int A and int B to return the result.
// E.G. a=5, b=2, result=52
func concatInts(a int, b int) int {
	return (a * 10) + b
}

// concatenate the first and last ints from a string
// E.G. str="z1bb4bb6", result=16
func concatFirstLastIntFromStr(str string) int {
	result := getNumsFromStr(str)
	if len(result) < 1 {
		panic("Each Row Must Contain at least 1 int")
	}
	if len(result) < 2 {
		a := result[0]
		return concatInts(a, a)
	}
	a := result[0]
	b := result[len(result)-1]
	return concatInts(a, b)
}

func readUntilEmpty() []string {
	strs := []string{}
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	for len(text) > 0 {
		strs = append(strs, text)
		text, _ = reader.ReadString('\n')
	}
	return strs
}

func sumLines(strs []string) int {
	sum := 0
	for _, str := range strs {
		sum += concatFirstLastIntFromStr(str)
	}
	return sum
}

func main() {
	strs := readUntilEmpty()
	num := sumLines(strs)
	print(num, "\n")
}
