package race

import (
	"math"
	"os"
	"testing"
)

func approxEqual(a float64, b float64) bool {
	return math.Abs(a-b) <= 0.001
}

// the following parabola tests were cross checked with wolfram alpha
// https://www.wolframalpha.com/input?i=y%3D-2*%28x%29%5E2%2B%28100x%29%2B2
func TestParabolaRoots1(t *testing.T) {
	parabola := Parabola{
		A: -2,
		B: 100,
		C: 2,
	}
	a, b, e := parabola.roots()
	if e != nil {
		t.Errorf("got error %s", e)
	}
	want := float64(-0.019996)
	if !approxEqual(a, want) {
		t.Errorf("got %f, wanted %f", a, want)
	}
	want = float64(50.020)
	if !approxEqual(b, want) {
		t.Errorf("got %f, wanted %f", b, want)
	}
}

// https://www.wolframalpha.com/input?i=y%3D-5*%28x%29%5E2%2B%28-3x%29%2B0
func TestParabolaRoots2(t *testing.T) {
	parabola := Parabola{
		A: -5,
		B: -3,
		C: 0,
	}
	a, b, e := parabola.roots()
	if e != nil {
		t.Errorf("got error %s", e)
	}
	want := float64(-0.6)
	if !approxEqual(a, want) {
		t.Errorf("got %f, wanted %f", a, want)
	}
	want = float64(0)
	if !approxEqual(b, want) {
		t.Errorf("got %f, wanted %f", b, want)
	}
}

// https://www.wolframalpha.com/input?i=y%3D-2*%28x%29%5E2%2B%28100x%29%2B2%3B+y%3D75
func TestParabolaYIntercept1(t *testing.T) {
	parabola := Parabola{
		A: -2,
		B: 100,
		C: 2,
	}
	a, b, e := parabola.findXForY(75)
	if e != nil {
		t.Errorf("got error %s", e)
	}
	want := float64(0.740981)
	if !approxEqual(a, want) {
		t.Errorf("got %f, wanted %f", a, want)
	}
	want = float64(49.259)
	if !approxEqual(b, want) {
		t.Errorf("got %f, wanted %f", b, want)
	}
}

// https://www.wolframalpha.com/input?i=y%3D-5*%28x%29%5E2%2B%28-3x%29%2B0%2C+y%3D-12
func TestParabolaIntercept2(t *testing.T) {
	parabola := Parabola{
		A: -5,
		B: -3,
		C: 0,
	}
	a, b, e := parabola.findXForY(-12)
	if e != nil {
		t.Errorf("got error %s", e)
	}
	want := float64(-1.87797)
	if !approxEqual(a, want) {
		t.Errorf("got %f, wanted %f", a, want)
	}
	want = float64(1.27797)
	if !approxEqual(b, want) {
		t.Errorf("got %f, wanted %f", b, want)
	}
}

func TestFindMinMaxHoldForRace(t *testing.T) {
	a, b, e := findMinMaxHoldForRace(7, 9)
	if e != nil {
		t.Errorf("got error %s", e)
	}
	want := int64(2)
	if a != want {
		t.Errorf("got %d, wanted %d", a, want)
	}
	want = int64(5)
	if b != want {
		t.Errorf("got %d, wanted %d", b, want)
	}
}

func TestQuantityOfWaysToWin(t *testing.T) {
	race := Race{
		raceTimeMS: 7,
		recordMM:   9,
	}
	got, e := race.quantityOfWaysToWin()
	if e != nil {
		t.Errorf("got error %s", e)
	}
	want := int64(4)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestComputeProductOfRaceWins(t *testing.T) {
	race1 := Race{
		raceTimeMS: 7,
		recordMM:   9,
	}
	race2 := Race{
		raceTimeMS: 15,
		recordMM:   40,
	}
	races := []Race{race1, race2}
	got := ComputeProductOfRaceWins(races)
	want := int64(32)
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestComputeExample(t *testing.T) {
	bytes, _ := os.ReadFile("./example.txt")
	str := string(bytes)
	got := ParseStringInput(str)
	want := 288
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestComputeReal(t *testing.T) {
	bytes, _ := os.ReadFile("./input.txt")
	str := string(bytes)
	got := ParseStringInput(str)
	want := 840336
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}
