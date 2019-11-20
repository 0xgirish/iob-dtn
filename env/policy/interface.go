// Package policy provides differente buffer management policies
// and implements KNOP, NP, GPP and LC policies
package policy

import (
	"errors"

	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer"
	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer/packet"
)

var FREE_SLOT_NOT_FOUND_ERROR = errors.New("FREE_SLOT_NOT_FOUND_ERROR")
var CAN_NOT_CREATE_SLOT_ERROR = errors.New("CAN_NOT_CREATE_SLOT_ERROR")

type Policy interface {
	// GetFreeSlot in the buffer, if there is no free slot then returns policyError
	getFreeSlot(b buffer.Buffer) (int, error)
	// CreateSlot removes a slot for the packet if possible else returns policyError
	CreateSlot(b buffer.Buffer, p packet.Packet, sensor_id int) (int, error)
}

type basePolicy struct {
	Name string
}

func (bp basePolicy) getFreeSlot(b buffer.Buffer) (int, error) {
	for i, p := range b.Packets {
		if !p.Exists() {
			return i, nil
		}
	}
	return -1, FREE_SLOT_NOT_FOUND_ERROR
}
