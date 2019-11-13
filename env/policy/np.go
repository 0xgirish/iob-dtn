package Policy

import (
	"github.com/iob-dtn/env/sensor/buffer"
	"github.com/iob-dtn/env/sensor/buffer/packet"
)

// NP No Priority
// the oldest packet in the buffer is discarded and replaced by the new one.
type NP struct {
	basePolicy
}

func (n NP) CreateSlot(b buffer.Buffer, p packet.Packet, sensor_id int) (int, policyError) {
	index, err := n.getFreeSlot(b)
	if err == nil {
		return index, nil
	}

	min_time := b.Packets[0].GetTimestamp()
	for i, pac := range b.Packets {
		if min_time.After(pac.GetTimestamp()) {
			min_time, index = pac.GetTimestamp(), i
		}
	}

	return index, err
}
