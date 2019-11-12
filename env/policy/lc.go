package Policy

import (
	"github.com/iob-dtn/env/sensor/buffer"
	"github.com/iob-dtn/env/sensor/buffer/packet"
)

// LC Lesser Copy
// When a packet is receivedor generated, it replaces the packet having the smallestnumber of copies in the buffer.
type LC struct {
	basePolicy
}

func (l LC) CreateSlot(b buffer.Buffer, p packet.Packet) (int, policyError) {
	// TODO:
}
