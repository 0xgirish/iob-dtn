package Policy

import (
	"crypto/rand"
	"math"

	"github.com/iob-dtn/env/sensor/buffer"
	"github.com/iob-dtn/env/sensor/buffer/packet"
)

// LC Lesser Copy
// When a packet is receivedor generated, it replaces the packet having the smallestnumber of copies in the buffer.
type LC struct {
	basePolicy
}

func (l LC) CreateSlot(b buffer.Buffer, p packet.Packet) (int, policyError) {
	index, err := l.getFreeSlot(b)
	if err == nil {
		return index, nil
	}

	min_copies := math.MaxInt32
	for i, p := range b.packets {
		if min_copies > p.GetCopies() {
			min_copies = p.GetCopies()
			index = i
		} else if min_copies == p.GetCopies() {
			if rand.Int()%2 == 0 {
				index = i
			}
		}
	}
	return index, nil
}
