package main

import "os"

func main() {
	argsWithoutProg := os.Args[1:]
	for _, arg := range argsWithoutProg {
		println(arg)
	}
	cmd := argsWithoutProg[0]
	println("------")
	if cmd == "3/1" {
		println("TODO run")
		// ratio.day3part1()
		// ratio.day3part1()
	}
}
