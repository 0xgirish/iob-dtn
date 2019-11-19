// Package sensor implements sensor for IoB
package sensor

import (
	"github.com/zkmrgirish/iob-dtn/env/policy"
	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer"
	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer/packet"
)

type Sensor struct {
	Id     int
	b      buffer.Buffer
	p      policy.Policy
	failes int
}

func New(id int, p policy.Policy) Sensor {
	return Sensor{
		b:  buffer.New(),
		p:  p,
		Id: id,
	}
}

func (s Sensor) GeneratePacket() Packet {
	return packet.New(s.Id)
}
