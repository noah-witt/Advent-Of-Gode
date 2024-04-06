package ratio

import "testing"

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
