package policy

import (
	"math"
	"math/rand"
	"time"

	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer"
	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer/packet"
)

// LC Lesser Copy
// When a packet is receivedor generated, it replaces the packet having the smallestnumber of copies in the buffer.
type LC struct {
	basePolicy
}

func (l LC) CreateSlot(b buffer.Buffer, p packet.Packet, sensor_id int) (int, error) {
	index, err := l.getFreeSlot(b)
	if err == nil {
		return index, nil
	}

	min_copies := math.MaxInt32
	for i, pac := range b.Packets {
		if min_copies > pac.GetCopies() {
			min_copies = pac.GetCopies()
			index = i
		} else if min_copies == pac.GetCopies() {
			rand.Seed(int64(time.Now().Nanosecond()))
			if rand.Int()%2 == 0 {
				index = i
			}
		}
	}
	return index, err
}
