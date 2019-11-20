// Packege util provides basic utility
package util

import "math"

type Position struct {
	x, y int
}

func (p Position) Distance(d Position) float64 {
	dx := float64(p.x - d.x)
	dy := float64(p.y - d.y)
	return math.Sqrt(dx*dx + dy*dy)
}
