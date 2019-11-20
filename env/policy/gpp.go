package policy

import (
	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer"
	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer/packet"
)

// GPP Generated Packet Priority
// whena packet is generated, it replaces the oldest received packet.
// But if there are only generated packets, then it replaces theoldest one.
// If a packet is received while the buffer is full, itis discarded.
type GPP struct {
	basePolicy
}

func (g GPP) CreateSlot(b buffer.Buffer, p packet.Packet, sensor_id int) (int, error) {
	index, err := g.getFreeSlot(b)
	if err == nil {
		return index, nil
	}

	// if the packet is received not generated then, do not allocate slot
	if sensor_id != p.Parent_id {
		return 0, CAN_NOT_CREATE_SLOT_ERROR
	}

	index = 0
	min_time := b.Packets[index].GetTimestamp()
	for i, pac := range b.Packets {
		if min_time.After(pac.GetTimestamp()) {
			min_time, index = pac.GetTimestamp(), i
		}
	}

	return index, err
}
