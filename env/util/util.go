// Packege util provides basic utility
package util

import (
	"math"
)

type Position struct {
	X, Y int
}

func (p Position) Distance(d Position) float64 {
	dx := float64(p.X - d.X)
	dy := float64(p.Y - d.Y)
	return math.Sqrt(dx*dx + dy*dy)
}
