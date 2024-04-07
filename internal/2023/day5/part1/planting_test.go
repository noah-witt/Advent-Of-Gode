package planting

import (
	"bufio"
	"os"
	"testing"
)

func TestStringOnlyNumberAndSpaceFalse(t *testing.T) {
	got := stringOnlyNumberAndSpace("23423:   ")
	want := false
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestStringOnlyNumberAndSpaceFalse2(t *testing.T) {
	got := stringOnlyNumberAndSpace("23423*   ")
	want := false
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestStringOnlyNumberAndSpaceTrue(t *testing.T) {
	got := stringOnlyNumberAndSpace("23423 23 2 4 9 0 2 7 ")
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestExample(t *testing.T) {
	file, _ := os.OpenFile("./example.txt", os.O_RDONLY, 0644)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	lines := []string{}
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	got := minSeedMap(lines)
	want := 35
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestReal(t *testing.T) {
	file, _ := os.OpenFile("./input.txt", os.O_RDONLY, 0644)
	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)
	lines := []string{}
	for fileScanner.Scan() {
		lines = append(lines, fileScanner.Text())
	}
	got := minSeedMap(lines)
	want := 3374647
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
