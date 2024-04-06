package ratio

import (
	"os"
	"testing"
)

func TestIsSymbolNot(t *testing.T) {

	got := isSymbol('x')
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestIsSymbolPass(t *testing.T) {
	got := isSymbol('*')
	want := true

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestIsSymbolNotDot(t *testing.T) {
	got := isSymbol('.')
	want := false

	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestIsAdjacentToSymbolAfter(t *testing.T) {
	got := isAdjacentToSymbol([]byte(".1+"), 1)
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestIsAdjacentToSymbolBefore(t *testing.T) {
	got := isAdjacentToSymbol([]byte("...+12..a"), 4)
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestIsAdjacentToSymbolDiagonalAbove(t *testing.T) {
	got := isAdjacentToSymbol([]byte("..*..\n...12..a"), 9)
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestIsAdjacentToSymbolDiagonalBelow(t *testing.T) {
	got := isAdjacentToSymbol([]byte(".*..\n...12..a\n..*"), 8)
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestIsAdjacentToSymbolDiagonalBelow2(t *testing.T) {
	got := isAdjacentToSymbol([]byte(".*..\n...12..a\n.....*"), 9)
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestRowOne(t *testing.T) {
	got := isAdjacentToSymbol([]byte("467..114..\n...*......"), 2)
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestLaterSection(t *testing.T) {
	got := isAdjacentToSymbol([]byte(".....+.58.\n..592....."), 15)
	want := true
	if got != want {
		t.Errorf("got %t, wanted %t", got, want)
	}
}

func TestExample(t *testing.T) {
	dat, _ := os.ReadFile("./example.txt")
	got := getNums(dat)
	want := 4361
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestReal(t *testing.T) {
	dat, _ := os.ReadFile("./input.txt")
	got := getNums(dat)
	println("got: ", got)
}
