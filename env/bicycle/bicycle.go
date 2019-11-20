// Package bicycle provides simulation of bicycles in the enviornments
package bicycle

import (
	"flag"
	"math/rand"
	"time"

	"github.com/zkmrgirish/iob-dtn/env/comdevice"
	"github.com/zkmrgirish/iob-dtn/env/sensor"
	"github.com/zkmrgirish/iob-dtn/env/util"
)

var speed uint
var sensorRange float64

func init() {
	flag.UintVar(&speed, "speed", 1, "speed of the bicycle, e.g. unit movement in per second")
	flag.Float64Var(&sensorRange, "range", 3.0, "bike sensor range")
}

type bicycle struct {
	pos      util.Position
	dest     util.Position
	receiver chan comdevice.Message
	comdevice.Comdevice
	sensor sensor.Sensor

	moving bool
}

// New bicycle with current position and destionation as pos
func New(pos util.Position, sensor sensor.Sensor) bicycle {
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
func (b bicycle) SetDestination(dest util.Position) {
	if b.Reached() {
		b.dest = dest
	}
}

// Move moves the bicycle randomly towards the destination
func (b bicycle) Move() {
	if b.Reached() {
		b.moving = false
		time.Sleep(100 * time.Millisecond)
		// TODO: request destination from environment
		// dest := env.Request()
		// b.SetDestination(dest)
		return
	}

	b.moving = true
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
