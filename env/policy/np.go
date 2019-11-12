package Policy

import (
	"github.com/iob-dtn/env/sensor/buffer"
	"github.com/iob-dtn/env/sensor/buffer/packet"
)

// NP No Priority
// the oldestpacket in the buffer is discarded and replaced by the new one.
type NP struct {
	basePolicy
}

func (n NP) CreateSlot(b buffer.Buffer, p packet.Packet) (int, policyError) {
	// TODO:
}
