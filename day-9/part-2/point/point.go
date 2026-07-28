package point

import (
	"fmt"
	"strings"
)

func abs(a int) int {
	if a < 0 {
		a = a * -1
	}
	return a
}

type CornerType uint8

func (c CornerType) IsTop() bool                    { return c/2 == 1 }
func (c CornerType) IsBottom() bool                 { return c/2 == 0 }
func (c CornerType) IsLeft() bool                   { return (c+1)/2 == 0 }
func (c CornerType) IsRight() bool                  { return (c+1)/2 == 0 }
func (c CornerType) IsMirror(other CornerType) bool { return c == ((other + 2) % 4) }

const (
	BottomLeft CornerType = iota
	BottomRight
	TopRight
	TopLeft
)

type Point struct {
	X              int
	Y              int
	Next           *Point
	Prev           *Point
	CornerType     CornerType
	InvertedCorner bool
}

func (p *Point) CrossesY(y int) bool {
	switch {
	case p.Y > y:
		if p.Next.Y < y {
			return true
		}
	case p.Y < y:
		if p.Next.Y > y {
			return true
		}
	case p.Y == y:
		// find the closest points on either side of p, that deviate from the line y
		prev := p.Prev
		next := p.Next
		for prev.Y == y {
			prev = prev.Prev
		}
		for next.Y == y {
			next = next.Next
		}

		// if those points are on either side of the line y, then we crossed
		if (prev.Y < y && next.Y > y) || (prev.Y > y && next.Y < y) {
			return true
		}
	}

	return false
}

// JUST FIX THE CORNERT TYPE PROBLEM AND THEN CHECKING ONLY POSSIBLE MATCHES, JUST CHECK FOR ANY POINTS WITHIN THE RECTANGLE. IF THERE IS A POINT IN THE RECTANGLE, EXCLUDE IT

// I think that I'm thinking about the crossing for the case when p.X == x wrong
//
// If it is on the line and it is within the rectangle (it must be within the Y bounds because we used withinY() to get the point in the first place), then ANY dip into the rectangle disqualifies immidiately. I would need to know whether the line is top or bottom
// Honestly, if i need to add that info in, then i should make a new method because it isn't exactly counting the crossings. I just need to know whether anything falls within the rectangle.
// // This means that if left < p.X < right then it fails immediately.
func (p *Point) CheckXBounds(start, end *Point, left, right int) bool {
	if p.X > left && p.X < right {
		return false
	}
	prev := p.Prev
	next := p.Next
	for prev.X == p.X {
		if prev == start || prev == end {
			break
		}
		prev = prev.Prev
	}
	if prev != start && prev != end && prev.X > left && prev.X < right {
		return false
	}
	for next.X == p.X {
		if next == start || next == end {
			break
		}
		next = next.Next
	}
	if next != start && next != end && next.X > left && next.X < right {
		return false
	}

	return true
}

func (p *Point) CheckYBounds(start, end *Point, top, bottom int) bool {
	if p.X > top && p.X < bottom {
		return false
	}
	prev := p.Prev
	next := p.Next
	for prev.X == p.X {
		if prev == start || prev == end {
			break
		}
		prev = prev.Prev
	}
	if prev != start && prev != end && prev.X > top && prev.X < bottom {
		return false
	}
	for next.X == p.X {
		if next == start || next == end {
			break
		}
		next = next.Next
	}
	if next != start && next != end && next.X > top && next.X < bottom {
		return false
	}

	return true
}

func (p *Point) CrossesX(x int) bool {
	switch {
	case p.X > x:
		if p.Next.X < x {
			return true
		}
	case p.X < x:
		if p.Next.X > x {
			return true
		}
	case p.X == x:
		// find the closest points on either side of p, that deviate from the line x
		prev := p.Prev
		next := p.Next
		for prev.X == x {
			prev = prev.Prev
		}
		for next.X == x {
			next = next.Next
		}

		// if those points are on either side of the line x, then we crossed
		if (prev.X < x && next.X > x) || (prev.X > x && next.X < x) {
			return true
		}
	}

	return false
}

func (p *Point) ComputeCornerType(sortedPoints SortedPoints) {
	switch {
	// top corner
	case
		p.X == p.Prev.X && p.Y < p.Prev.Y,
		p.X == p.Next.X && p.Y < p.Next.Y:

		switch {
		// left corner
		case
			p.Y == p.Prev.Y && p.X < p.Prev.X,
			p.Y == p.Next.Y && p.X < p.Next.X:

			p.CornerType = TopLeft

		// right corner
		case
			p.Y == p.Prev.Y && p.X > p.Prev.X,
			p.Y == p.Next.Y && p.X > p.Next.X:

			p.CornerType = TopRight

		default:
			panic(fmt.Sprintf("unknown corner type for %s -> %s -> %s", p.Prev.String(), p.String(), p.Next.String()))
		}

	// bottom corner
	case
		p.X == p.Prev.X && p.Y > p.Prev.Y,
		p.X == p.Next.X && p.Y > p.Next.Y:

		switch {
		// left corner
		case
			p.Y == p.Prev.Y && p.X < p.Prev.X,
			p.Y == p.Next.Y && p.X < p.Next.X:

			p.CornerType = BottomLeft

		// right corner
		case
			p.Y == p.Prev.Y && p.X > p.Prev.X,
			p.Y == p.Next.Y && p.X > p.Next.X:

			p.CornerType = BottomRight

		default:
			panic(fmt.Sprintf("unknown corner type for %s -> %s -> %s", p.Prev.String(), p.String(), p.Next.String()))
		}

	default:
		panic(fmt.Sprintf("unknown corner type for %s -> %s -> %s", p.Prev.String(), p.String(), p.Next.String()))
	}

	var potentialCrosses []Point
	var numCrosses int
	switch p.CornerType {
	case BottomLeft:
		potentialCrosses = sortedPoints.PointsWithinX(0, p.X+1)
		for _, point := range potentialCrosses {
			if point.CrossesY(p.Y - 1) {
				numCrosses++
			}
		}
	case BottomRight:
		potentialCrosses = sortedPoints.PointsWithinX(0, p.X-1)
		for _, point := range potentialCrosses {
			if point.CrossesY(p.Y - 1) {
				numCrosses++
			}
		}
	case TopRight:
		potentialCrosses = sortedPoints.PointsWithinX(0, p.X-1)
		for _, point := range potentialCrosses {
			if point.CrossesY(p.Y + 1) {
				numCrosses++
			}
		}
	case TopLeft:
		potentialCrosses = sortedPoints.PointsWithinX(0, p.X+1)
		for _, point := range potentialCrosses {
			if point.CrossesY(p.Y + 1) {
				numCrosses++
			}
		}
	default:
		panic(fmt.Sprintf("unknown corner type %d for %s", p.CornerType, p.String()))
	}

	if numCrosses%2 == 0 {
		p.InvertedCorner = true
	}
}

func (p Point) String() string {
	return fmt.Sprintf("(%d, %d)%s -> [%d, %d]%s -> (%d, %d)%s", p.Prev.X, p.Prev.Y, p.Prev.cornerString(), p.X, p.Y, p.cornerString(), p.Next.X, p.Next.Y, p.Next.cornerString())
}

func (p Point) cornerString() string {
	var builder strings.Builder
	if p.InvertedCorner {
		builder.WriteByte('^')
	}
	switch p.CornerType {
	case TopRight:
		builder.WriteString("TR")
	case TopLeft:
		builder.WriteString("TL")
	case BottomLeft:
		builder.WriteString("BL")
	case BottomRight:
		builder.WriteString("BR")
	}
	return builder.String()
}

func (p *Point) RectArea(sorted SortedPoints, other Point) int {
	if !p.InvertedCorner && !other.InvertedCorner && !p.CornerType.IsMirror(other.CornerType) {
		// fmt.Printf("comparing %s to %s\t\tcorners not compatable\n", p.String(), other.String())
		return 0
	}

	//// grab the left, right, top, and bottom for the rectangle
	left := min(p.X, other.X)
	right := max(p.X, other.X)
	top := min(p.Y, other.Y)
	bottom := max(p.Y, other.Y)

	// grab all points between the left and right
	points := sorted.PointsWithinX(left, right)
	// fmt.Printf("comparing %s to %s\t\tpoints within L/R [%d, %d]\n", p.String(), other.String(), left, right)
	// for _, p := range points {
		// fmt.Println(p)
	// }
	// fmt.Println()
	for _, point := range points {
		if point.Y > top && point.Y < bottom {
			// fmt.Printf("comparing %s to %s\t\tT/B: [%d, %d], L/R: [%d, %d]\npoint %s contained within\n\n", p.String(), other.String(), top, bottom, left, right, point.String())
			return 0
		}
		if point.Y <= top && (point.Prev.Y > top || point.Next.Y > top) {
			// fmt.Printf("comparing %s to %s\t\tT/B: [%d, %d], L/R: [%d, %d]\npoint %s crosses top\n\n", p.String(), other.String(), top, bottom, left, right, point.String())
			return 0
		}
		if point.Y >= bottom && (point.Prev.Y < bottom || point.Next.Y < bottom) {
			// fmt.Printf("comparing %s to %s\t\tT/B: [%d, %d], L/R: [%d, %d]\npoint %s crosses bottom\n\n", p.String(), other.String(), top, bottom, left, right, point.String())
			return 0
		}
	}
	points = sorted.PointsWithinY(top, bottom)
	// fmt.Printf("comparing %s to %s\t\tpoints within T/B [%d, %d]\n", p.String(), other.String(), top, bottom)
	// for _, p := range points {
		// fmt.Println(p)
	// }
	// fmt.Println()
	for _, point := range points {
		if point.X <= left && (point.Prev.X > left || point.Next.X > left) {
			// fmt.Printf("comparing %s to %s\t\tT/B: [%d, %d], L/R: [%d, %d]\npoint %s crosses left\n\n", p.String(), other.String(), top, bottom, left, right, point.String())
			return 0
		}
		if point.X >= right && (point.Prev.X < right || point.Next.X < right) {
			// fmt.Printf("comparing %s to %s\t\tT/B: [%d, %d], L/R: [%d, %d]\npoint %s crosses right\n\n", p.String(), other.String(), top, bottom, left, right, point.String())
			return 0
		}
	}

	dx := abs(p.X-other.X) + 1
	dy := abs(p.Y-other.Y) + 1
	// fmt.Printf("comparing %s to %s\t\tT/B: [%d, %d], L/R: [%d, %d]\narea:\t%d\n\n", p.String(), other.String(), top, bottom, left, right, dx*dy)
	return dx * dy
}
