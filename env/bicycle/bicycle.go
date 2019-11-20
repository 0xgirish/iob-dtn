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

func init() {
	flag.UintVar(&speed, "speed", 1, "speed of the bicycle, e.g. unit movement in per second")
}

type Bicycle struct {
	pos  util.Position
	dest util.Position
	comdevice.Comdevice
	sensor *sensor.Sensor
	env    requestRanger

	stop   chan bool
	moving bool
}

type requestRanger interface {
	// Range returns slice of devices which are in the range
	Range(pos util.Position) []comdevice.Comdevice
	RequestDestination() util.Position
}

// New bicycle with current position and destionation as pos
func New(pos util.Position, s sensor.Sensor, r requestRanger, stop chan bool) Bicycle {
	return Bicycle{
		pos:       pos,
		dest:      pos,
		env:       r,
		Comdevice: comdevice.New(&s, &pos),
		sensor:    &s,
		stop:      stop,
	}
}

func (b Bicycle) GetPosition() util.Position {
	return b.pos
}

// SetDestination of the bicycle if the bicycle has reached the station
func (b Bicycle) SetDestination(dest util.Position) {
	if b.Reached() {
		b.dest = dest
	}
}

func (b Bicycle) Initiate() {
	move := time.Tick(time.Duration(1000/int(speed)) * time.Millisecond)
	duty := time.Tick(time.Duration(1000/int(sensor.Generation_frequency)) * time.Millisecond)

	generate := true
	for {
		select {
		case <-move:
			b.Move()
		case <-duty:
			if generate {
				go b.sensor.GeneratePacket()
			} else {
				go b.SendPackets()
			}
			generate = (!generate)
		case <-b.stop:
			return
		}
	}
}

func (b Bicycle) SendPackets() {
	for ind, pkt := range b.sensor.B.Packets {
		if pkt.GetCopies() > 1 {
			for _, dvc := range b.env.Range(b.pos) {
				msg := b.Comdevice.Send(comdevice.Message{
					Type: comdevice.PacketTransfer,
					Msg:  b.sensor.B.Packets[ind],
					From: b.sensor.Id,
					To:   dvc.ID(),
				}, dvc)
				if msg.Type == comdevice.ACK {
					copies, ok := msg.Msg.(int)
					if !ok {
						continue
					}
					b.sensor.B.Packets[ind].SetCopies(copies)
				}
			}
		}
	}
}

// Move moves the bicycle randomly towards the destination
func (b Bicycle) Move() {
	if b.Reached() {
		b.moving = false
		time.Sleep(100 * time.Millisecond)
		dest := b.env.RequestDestination()
		b.SetDestination(dest)
		return
	}

	b.moving = true
	vertical := rand.Int()%2 == 0
	if vertical {
		if b.dest.Y-b.pos.Y > 0 {
			b.pos.Y++
			return
		} else if b.dest.Y-b.pos.Y < 0 {
			b.pos.Y--
			return
		}
	}

	if b.dest.X-b.pos.X > 0 {
		b.pos.X++
	} else if b.dest.X-b.pos.X < 0 {
		b.pos.X--
	}
}

func (b Bicycle) Moving() bool {
	return b.moving
}

// Reached checks if the bicycle has reached the destination
func (b Bicycle) Reached() bool {
	return (b.dest.X == b.pos.X) && (b.dest.Y == b.pos.Y)
}
