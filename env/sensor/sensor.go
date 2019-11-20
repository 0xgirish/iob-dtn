// Package sensor implements sensor for IoB
package sensor

import (
	"errors"
	"flag"
	"time"

	"github.com/zkmrgirish/iob-dtn/env/manager"
	"github.com/zkmrgirish/iob-dtn/env/policy"
	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer"
	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer/packet"
)

var generation_frequency uint

func init() {
	flag.UintVar(&generation_frequency, "freq", 4, "number of packets to be generated in a second")
}

type Sensor struct {
	Id      int
	b       buffer.Buffer
	p       policy.Policy
	started bool
	stop    chan bool
}

// New creates a sensor and registers it with manager
func New(id int, p policy.Policy) Sensor {
	manager.Register(id)
	return Sensor{
		b:  buffer.New(),
		p:  p,
		Id: id,
	}
}

// GeneratePacket generates a packet
func (s Sensor) GeneratePacket() packet.Packet {
	return packet.New(s.Id)
}

// Start generating packets
func (s Sensor) Start() {
	if s.started {
		return
	}
	s.started = true
	wait_duration := time.Duration(1000/int(generation_frequency)) * time.Millisecond
	tick := time.Tick(wait_duration)
	go func() {
		for {
			select {
			case <-tick:
				pkt := s.GeneratePacket()
				ind, err := s.p.CreateSlot(s.b, pkt, s.Id)
				if err != nil && errors.As(err, policy.CAN_NOT_CREATE_SLOT_ERROR) {
					continue
				}
				s.b.Add(pkt, ind)

			case <-s.stop:
				s.started = false
				return
			}
		}

	}()
}

// Stop sensor from generating packets
func (s Sensor) Stop() {
	s.stop <- true
}
