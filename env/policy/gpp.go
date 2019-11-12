package Policy

import (
	"github.com/iob-dtn/env/sensor/buffer"
	"github.com/iob-dtn/env/sensor/buffer/packet"
)

// GPP Generated Packet Priority
// whena packet is generated, it replaces the oldest received packet.
// But if there are only generated packets, then it replaces theoldest one.
// If a packet is received while the buffer is full, itis discarded.
type GPP struct {
	basePolicy
}

func (g GPP) CreateSlot(b buffer.Buffer, p packet.Packet) (int, policyError) {
	// TODO:
}
