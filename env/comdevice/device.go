package comdevice

import (
	"errors"

	"github.com/zkmrgirish/iob-dtn/env/manager"
	"github.com/zkmrgirish/iob-dtn/env/policy"
	"github.com/zkmrgirish/iob-dtn/env/sensor"
	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer"
	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer/packet"
	"github.com/zkmrgirish/iob-dtn/env/util"
)

// device is the physical hardware for the sensor
type device struct {
	s   *sensor.Sensor
	Pos *util.Position
}

func (d device) ID() int {
	return d.s.Id
}

// Receive sends the message to device's receiver channel
func (d device) Receive(msg Message) Message {
	if d.s.Id != msg.To {
		return Message{
			Type: ERR,
			Msg:  0,
			From: d.s.Id,
			To:   msg.From,
		}
	}

	if msg.Type == ACK {
		return Message{
			Type: OK,
			From: d.s.Id,
			To:   msg.From,
		}
	}

	pkt, ok := msg.Msg.(packet.Packet)
	if !ok {
		return Message{
			Type: ERR,
			From: d.s.Id,
			To:   msg.From,
		}
	}

	if d.s.B.InBuffer(pkt) {
		return Message{
			Type: ACK,
			Msg:  0,
			From: d.s.Id,
			To:   msg.From,
		}
	}

	ind, err := d.s.P.CreateSlot(d.s.B, pkt, d.s.Id)
	if err != nil && errors.As(err, policy.CAN_NOT_CREATE_SLOT_ERROR) {
		return Message{
			Type: ACK,
			Msg:  0,
			From: d.s.Id,
			To:   msg.From,
		}
	}
	n := pkt.GetCopies()
	pkt.DecreaseCopies(n / 2)
	err = d.s.B.Add(pkt, ind)
	if err != nil && !errors.As(err, buffer.BUFFER_PACKET_REMOVED_ERROR) {
		return Message{
			Type: ACK,
			Msg:  0,
			From: d.s.Id,
			To:   msg.From,
		}
	}

	return Message{
		Type: ACK,
		Msg:  n - pkt.GetCopies(),
		From: d.s.Id,
		To:   msg.From,
	}
}

// Send sends the message to device's receiver
func (d device) Send(msg Message, dvc Comdevice) Message {
	return dvc.Receive(msg)
}

// New Comdevice for a sensor
func New(s *sensor.Sensor, pos *util.Position) Comdevice {
	return device{
		s:   s,
		Pos: pos,
	}
}

type stationDevice struct {
	id  int
	Pos *util.Position
}

func (sd stationDevice) ID() int {
	return sd.id
}

func (sd stationDevice) Receive(msg Message) Message {
	if sd.id != msg.To {
		return Message{
			Type: ERR,
			Msg:  0,
			From: sd.id,
			To:   msg.From,
		}
	}

	if msg.Type != PacketTransfer {
		return Message{
			Type: ERR,
			Msg:  0,
			From: sd.id,
			To:   msg.From,
		}
	}

	pkt, ok := msg.Msg.(packet.Packet)
	if !ok {
		return Message{
			Type: ERR,
			Msg:  0,
			From: sd.id,
			To:   msg.From,
		}
	}
	manager.MarkSuccess(pkt.Parent_id, pkt.Id)
	return Message{
		Type: ACK,
		Msg:  pkt.GetCopies(),
		From: sd.id,
		To:   msg.From,
	}
}

func (sd stationDevice) Send(msg Message, dvc Comdevice) Message {
	return dvc.Receive(msg)
}

func NewStationDevice(id int, pos *util.Position) Comdevice {
	return stationDevice{
		id:  id,
		Pos: pos,
	}
}
