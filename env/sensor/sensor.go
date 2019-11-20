// Package sensor implements sensor for IoB
package sensor

import (
	"errors"
	"flag"

	"github.com/zkmrgirish/iob-dtn/env/manager"
	"github.com/zkmrgirish/iob-dtn/env/policy"
	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer"
	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer/packet"
)

var Generation_frequency uint

func init() {
	flag.UintVar(&Generation_frequency, "freq", 4, "number of packets to be generated in a second")
}

// Sensor simulation
type Sensor struct {
	Id      int
	B       buffer.Buffer
	P       policy.Policy
	started bool
}

// New creates a sensor and registers it with manager
func New(id int, p policy.Policy) Sensor {
	manager.Register(id)
	return Sensor{
		B:  buffer.New(),
		P:  p,
		Id: id,
	}
}

// GeneratePacket generates a packet
func (s *Sensor) GeneratePacket() {
	pkt := packet.New(s.Id)
	ind, err := s.P.CreateSlot(s.B, pkt, s.Id)
	if err != nil && errors.Is(err, policy.CAN_NOT_CREATE_SLOT_ERROR) {
		return
	}
	s.B.Add(pkt, ind)
}
