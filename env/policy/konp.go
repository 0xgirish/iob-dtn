package policy

import (
	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer"
	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer/packet"
)

// KNOP Keep Oldest No Priority
// This policy is an usual"First In First Served" buffer, such as in basic network routerbuffers.
// if the buffer if full, the new packet is discarded
type KONP struct {
	basePolicy
}

func (k KONP) CreateSlot(b buffer.Buffer, p packet.Packet, sensor_id int) (int, policyError) {
	index, err := k.getFreeSlot(b)
	if err != nil {
		return 0, CAN_NOT_CREATE_SLOT_ERROR
	}
	return index, nil
}
