package Policy

import (
	"time"

	"github.com/iob-dtn/env/sensor/buffer"
	"github.com/iob-dtn/env/sensor/buffer/packet"
)

// NP No Priority
// the oldest packet in the buffer is discarded and replaced by the new one.
type NP struct {
	basePolicy
}

func (n NP) CreateSlot(b buffer.Buffer, p packet.Packet, sensor_id int) (int, policyError) {
	index, err := l.getFreeSlot(b)
	if err == nil {
		return index, nil
	}

	var min_time time.Time
	var first_time = true

	for i, pac := range b.Packets {
		if first_time {
			min_time, index = pac.GetTimestamp(), i
			first_time = false
			continue
		}
		if min_time.After(pac.GetTimestamp()) {
			min_time, index = pac.GetTimestamp(), i
		}
	}

	return index, err
}
