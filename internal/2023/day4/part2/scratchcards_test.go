package scratchcards

import (
	"bufio"
	"os"
	"testing"
)

func TestCalcScore1(t *testing.T) {
	got := calcScore(1)
	want := 1
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCalcScore2(t *testing.T) {
	got := calcScore(2)
	want := 2
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestCalcScore3(t *testing.T) {
	got := calcScore(3)
	want := 4
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestRowOne(t *testing.T) {
	lines := []string{"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"}
	got := sumCards(lines)
	want := 1
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestRow2(t *testing.T) {
	lines := []string{"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"}
	got := sumCards(lines)
	want := 1
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestRow3(t *testing.T) {
	lines := []string{"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83"}
	got := sumCards(lines)
	want := 1
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestParseRow(t *testing.T) {
	line := "82 41 56 54 18 62 29 55 34 20"
	got := getNumsFromStr(line)
	want := []int{82, 41, 56, 54, 18, 62, 29, 55, 34, 20}
	if len(got) != len(want) {
		t.Errorf("length does not match")
	}
	for index, val := range got {
		if val != want[index] {
			t.Errorf("got %d but wanted %d at index %d", val, want[index], index)
		}
	}
}

func TestParseRow2(t *testing.T) {
	line := "37 14 10 80 58 11 65 96 90  8 59 32 53 21 98 83 17  9 87 25 71 77 70 73 24"
	got := getNumsFromStr(line)
	want := []int{37, 14, 10, 80, 58, 11, 65, 96, 90, 8, 59, 32, 53, 21, 98, 83, 17, 9, 87, 25, 71, 77, 70, 73, 24}
	if len(got) != len(want) {
		t.Errorf("length does not match")
	}
	for index, val := range got {
		if val != want[index] {
			t.Errorf("got %d but wanted %d at index %d", val, want[index], index)
		}
	}
}

func TestRealRow1(t *testing.T) {
	lines := []string{"Card   1: 82 41 56 54 18 62 29 55 34 20 | 37 14 10 80 58 11 65 96 90  8 59 32 53 21 98 83 17  9 87 25 71 77 70 73 24"}
	got := sumCards(lines)
	print(got)
	// want := 1
	// if got != want {
	// 	t.Errorf("got %d, wanted %d", got, want)
	// }
}

func TestCalcScore4(t *testing.T) {
	got := calcScore(4)
	want := 8
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
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
	got := sumCards(lines)
	want := 30
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
	got := sumCards(lines)
	print(got)
}
