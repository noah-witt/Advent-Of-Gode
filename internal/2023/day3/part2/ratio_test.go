package ratio

import (
	"os"
	"testing"
)

func TestExample(t *testing.T) {
	dat, _ := os.ReadFile("./example.txt")
	got := process(dat)
	want := 467835
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test1(t *testing.T) {
	str := "77.22\n..*.."
	got := process([]byte(str))
	want := 1694
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test2(t *testing.T) {
	str := "111\n..*222."
	got := process([]byte(str))
	want := 24642
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test3(t *testing.T) {
	str := "111\n..*..\n333"
	got := process([]byte(str))
	want := 36963
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test4(t *testing.T) {
	str := "...111\n..*..\n333"
	got := process([]byte(str))
	want := 36963
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test5(t *testing.T) {
	str := "...111\n..*..\n...333"
	got := process([]byte(str))
	want := 36963
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test6(t *testing.T) {
	str := "....111\n..*..\n...333"
	got := process([]byte(str))
	want := 0
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test7(t *testing.T) {
	str := "...111\n..*1..\n...333"
	got := process([]byte(str))
	want := 0
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test8(t *testing.T) {
	str := "......111\n..2*1..\n......333"
	got := process([]byte(str))
	want := 2
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func Test9(t *testing.T) {
	str := "......111\n..2.*.1..\n......333\n3*2\n...\n5\n6*2\n...\n**\n77*77\n..3"
	got := process([]byte(str))
	want := 6
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestEnd(t *testing.T) {
	str := "..889................695........654..750.....*.............637........./...............................780....*726....233...*...............\n..................../.................*.....453.....642....*.........828......@...94...........152/...*....790.......*.....445......../.....\n...........................51.......681........................271..........719.......................964......399..426...............456..."
	got := process([]byte(str))
	want := 1935468
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestReal(t *testing.T) {
	dat, _ := os.ReadFile("./input.txt")
	got := process(dat)
	println("got: ", got)
}
