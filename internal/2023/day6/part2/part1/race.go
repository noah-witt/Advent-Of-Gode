package race

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

/**
https://adventofcode.com/2023/day/6

As part of signing up, you get a sheet of paper (your puzzle input) that lists the time allowed for each race and also the best distance ever recorded in that race. To guarantee you win the grand prize, you need to make sure you go farther in each race than the current record holder.

The organizer brings you over to the area where the boat races are held. The boats are much smaller than you expected - they're actually toy boats, each with a big button on top. Holding down the button charges the boat, and releasing the button allows the boat to move. Boats move faster if their button was held longer, but time spent holding the button counts against the total race time. You can only hold the button at the start of the race, and boats don't move until the button is released.

For example:

Time:      7  15   30
Distance:  9  40  200
This document describes three races:

The first race lasts 7 milliseconds. The record distance in this race is 9 millimeters.
The second race lasts 15 milliseconds. The record distance in this race is 40 millimeters.
The third race lasts 30 milliseconds. The record distance in this race is 200 millimeters.
Your toy boat has a starting speed of zero millimeters per millisecond. For each whole millisecond you spend at the beginning of the race holding down the button, the boat's speed increases by one millimeter per millisecond.

So, because the first race lasts 7 milliseconds, you only have a few options:

Don't hold the button at all (that is, hold it for 0 milliseconds) at the start of the race. The boat won't move; it will have traveled 0 millimeters by the end of the race.
Hold the button for 1 millisecond at the start of the race. Then, the boat will travel at a speed of 1 millimeter per millisecond for 6 milliseconds, reaching a total distance traveled of 6 millimeters.
Hold the button for 2 milliseconds, giving the boat a speed of 2 millimeters per millisecond. It will then get 5 milliseconds to move, reaching a total distance of 10 millimeters.
Hold the button for 3 milliseconds. After its remaining 4 milliseconds of travel time, the boat will have gone 12 millimeters.
Hold the button for 4 milliseconds. After its remaining 3 milliseconds of travel time, the boat will have gone 12 millimeters.
Hold the button for 5 milliseconds, causing the boat to travel a total of 10 millimeters.
Hold the button for 6 milliseconds, causing the boat to travel a total of 6 millimeters.
Hold the button for 7 milliseconds. That's the entire duration of the race. You never let go of the button. The boat can't move until you let go of the button. Please make sure you let go of the button so the boat gets to move. 0 millimeters.
Since the current record for this race is 9 millimeters, there are actually 4 different ways you could win: you could hold the button for 2, 3, 4, or 5 milliseconds at the start of the race.

In the second race, you could hold the button for at least 4 milliseconds and at most 11 milliseconds and beat the record, a total of 8 different ways to win.

In the third race, you could hold the button for at least 11 milliseconds and no more than 19 milliseconds and still beat the record, a total of 9 ways you could win.

To see how much margin of error you have, determine the number of ways you can beat the record in each race; in this example, if you multiply these values together, you get 288 (4 * 8 * 9).

Determine the number of ways you could beat the record in each race. What do you get if you multiply these numbers together?
**/

// race performance is represented by the quadratic equation
// D=(TM)^2+TM*TT+0
// y=x^2+x*tt+0
// A is always -1 in this example
// B is the total time
// C is always 0 in this example
// Represents the coefficients of a quadratic equation y=(ax^2 + bx + c)
type Parabola struct {
	A float64
	B float64
	C float64
}

func (p *Parabola) y(x float64) float64 {
	return p.A*math.Pow(x, 2) + p.B*x + p.C
}

// Finds the x-intercepts (roots) of the parabola where it crosses the y-axis (y = 0)
func (p *Parabola) roots() (float64, float64, error) {
	discriminant := p.B*p.B - 4*p.A*p.C
	if discriminant < 0 {
		return 0, 0, fmt.Errorf("parabola has no real roots (discriminant is negative)")
	}
	root1 := (-p.B + math.Sqrt(discriminant)) / (2 * p.A)
	root2 := (-p.B - math.Sqrt(discriminant)) / (2 * p.A)
	return root1, root2, nil
}

// Finds the x-values where the parabola intersects a specific y-value
func (p *Parabola) findXForY(y float64) (float64, float64, error) {
	equation := Parabola{
		A: p.A,
		B: p.B,
		C: p.C - y,
	}
	root1, root2, e := equation.roots()
	if e != nil {
		return 0, 0, e
	}
	return root1, root2, nil
}

func makeParabolaForRace(raceTimeMS float64) Parabola {
	return Parabola{
		A: -1,
		B: raceTimeMS,
		C: 0,
	}
}

// Find the intercepts
func findMinMaxHoldForRace(raceTimeMS int64, goalMM int64) (int64, int64, error) {
	// parabola represents all possible distances covered by boat
	parabola := makeParabolaForRace(float64(raceTimeMS))
	// find the intercepts where the boat goes goalMM
	startMS, endMS, e := parabola.findXForY(float64(goalMM))
	if e != nil {
		return 0, 0, e
	}
	return int64(math.Ceil(startMS)), int64(math.Floor(endMS)), nil
}

type Race struct {
	raceTimeMS int64
	recordMM   int64
}

func (r *Race) quantityOfWaysToWin() (int64, error) {
	a, b, e := findMinMaxHoldForRace(r.raceTimeMS, r.recordMM)
	if e != nil {
		return 0, e
	}
	// inclusive so add 1
	return (b - a) + 1, nil
}

// compute product of all race wins
func ComputeProductOfRaceWins(races []Race) int64 {
	product := int64(1)
	for _, race := range races {
		wins, e := race.quantityOfWaysToWin()
		if e != nil {
			panic(e)
		}
		product *= wins
	}
	return product
}

func getNumberFromStr(section string) int {
	numStr := strings.ReplaceAll(section, " ", "")
	num, _ := strconv.Atoi(numStr)
	return num
}

func ParseStringInput(input string) int {
	lines := strings.Split(input, "\n")
	timeLine := lines[0]
	distanceLine := lines[1]
	timeLineTimes := (strings.Split(timeLine, ":"))[1]
	distanceLineDistances := (strings.Split(distanceLine, ":"))[1]
	time := getNumberFromStr(timeLineTimes)
	distance := getNumberFromStr(distanceLineDistances)
	race := Race{recordMM: int64(distance + 1), raceTimeMS: int64(time)}
	result := ComputeProductOfRaceWins([]Race{race})
	print(result)
	return int(result)
}
