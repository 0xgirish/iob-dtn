// Package bicycle provides simulation of bicycles in the enviornments
package bicycle

import (
	"flag"
	"math"
	"math/rand"
	"time"

	"github.com/zkmrgirish/iob-dtn/env/comdevice"
	"github.com/zkmrgirish/iob-dtn/env/sensor"
)

var speed uint
var sensorRange float64

func init() {
	flag.UintVar(&speed, "speed", 1, "speed of the bicycle, e.g. unit movement in per second")
	flag.Float64Var(&sensorRange, "range", 3.0, "bike sensor range")
}

type Position struct {
	x, y int
}

func (p Position) Distance(d Position) float64 {
	dx := float64(p.x - d.x)
	dy := float64(p.y - d.y)
	return math.Sqrt(dx*dx + dy*dy)
}

type bicycle struct {
	pos      Position
	dest     Position
	receiver chan comdevice.Message
	comdevice.Comdevice
	sensor sensor.Sensor

	moving bool
}

// New bicycle with current position and destionation as pos
func New(pos Position, sensor sensor.Sensor) bicycle {
	receiver := make(chan comdevice.Message)
	return bicycle{
		pos:       pos,
		dest:      pos,
		receiver:  receiver,
		Comdevice: comdevice.New(receiver),
		sensor:    sensor,
	}
}

// SetDestination of the bicycle if the bicycle has reached the station
func (b bicycle) SetDestination(dest Position) {
	if b.Reached() {
		b.dest = dest
	}
}

// Move moves the bicycle randomly towards the destination
func (b bicycle) Move() {

	if b.Reached() {
		time.Sleep(100 * time.Millisecond)
		b.moving = false
		// TODO: request destination from environment
		b.moving = true
	}

	vertical := rand.Int()%2 == 0
	if vertical {
		if b.dest.y-b.pos.y > 0 {
			b.pos.y++
			return
		} else if b.dest.y-b.pos.y < 0 {
			b.pos.y--
			return
		}
	}

	if b.dest.x-b.pos.x > 0 {
		b.pos.x++
	} else if b.dest.x-b.pos.x < 0 {
		b.pos.x--
	}

}

// Reached checks if the bicycle has reached the destination
func (b bicycle) Reached() bool {
	return (b.dest.x == b.pos.x) && (b.dest.y == b.pos.y)
}
