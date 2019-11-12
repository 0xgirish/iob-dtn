package Policy

import (
	"github.com/iob-dtn/env/sensor/buffer"
	"github.com/iob-dtn/env/sensor/buffer/packet"
)

// KNOP Keep Oldest No Priority
// This policy is an usual"First In First Served" buffer, such as in basic network routerbuffers.
// if the buffer if full, the new packet is discarded
type KONP struct {
	basePolicy
}

func (k KNOP) CreateSlot(b buffer.Buffer, p packet.Packet) (int, policyError) {
	// TODO:
}
