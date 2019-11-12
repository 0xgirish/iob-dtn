// Package policy provides differente buffer management policies
// and implements KNOP, NP, GPP and LC policies
package policy

import (
	"errors"

	"github.com/iob-dtn/env/sensor/buffer"
	"github.com/iob-dtn/env/sensor/buffer/packet"
)

type policyError error

const FREE_SLOT_NOT_FOUND_ERROR = policyError(errors.New("FREE_SLOT_NOT_FOUND_ERROR"))
const CAN_NOT_CREATE_SLOT_ERROR = policyError(errors.New("CAN_NOT_CREATE_SLOT_ERROR"))

type Policy interface {
	// GetFreeSlot in the buffer, if there is no free slot then returns policyError
	GetFreeSlot(b buffer.Buffer) (int, policyError)
	// CreateSlot removes a slot for the packet if possible else returns policyError
	CreateSlot(b buffer.Buffer, p packet.Packet) (int, policyError)
}

type basePolicy struct {
	Name string
}

func (b basePolicy) GetFreeSlot(b buffer.Buffer) (int, policyError) {
	for i, p := range b.packets {
		if !p.Exists() {
			return i, nil
		}
	}
	return -1, FREE_SLOT_NOT_FOUND_ERROR
}
