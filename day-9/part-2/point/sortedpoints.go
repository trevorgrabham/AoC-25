package point

import (
	"cmp"
	"fmt"
	"slices"
	"strings"
)

type SortedPoints struct {
	PointsX []*Point
	PointsY []*Point
}

// not including start and end
func (s *SortedPoints) PointsWithinX(start, end int) []Point {
	var (
		points []Point
		i      int
	)

	for i = 0; i < len(s.PointsX); i++ {
		if s.PointsX[i].X > start {
			break
		}
	}

	for ; i < len(s.PointsX); i++ {
		if s.PointsX[i].X >= end {
			break
		}

		points = append(points, *s.PointsX[i])
	}

	return points
}

// not including start and end
func (s *SortedPoints) PointsWithinY(start, end int) []Point {
	var (
		points []Point
		i      int
	)

	for i = 0; i < len(s.PointsY); i++ {
		if s.PointsY[i].Y > start {
			break
		}
	}

	for ; i < len(s.PointsY); i++ {
		if s.PointsY[i].Y >= end {
			break
		}

		points = append(points, *s.PointsY[i])
	}

	return points
}

func (s *SortedPoints) Sort() {
	slices.SortFunc(s.PointsX, func(a, b *Point) int { return cmp.Compare(a.X, b.X) })
	slices.SortFunc(s.PointsY, func(a, b *Point) int { return cmp.Compare(a.Y, b.Y) })
}

func (s *SortedPoints) AddPoint(p *Point) {
	s.PointsX = append(s.PointsX, p)
	s.PointsY = append(s.PointsY, p)
}

func (s SortedPoints) String() string {
	var output strings.Builder
	for i := range len(s.PointsX) - 1 {
		output.WriteString(fmt.Sprintf("%s\t|\t%s\n", s.PointsX[i], s.PointsY[i]))
	}
	output.WriteString(fmt.Sprintf("%s\t|\t%s", s.PointsX[len(s.PointsX)-1], s.PointsY[len(s.PointsX)-1]))
	return output.String()
}
