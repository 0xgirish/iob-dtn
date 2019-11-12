// Package buffer implements buffer for the sensor
package buffer

import (
	"errors"

	"github.com/iob-dtn/env/sensor/buffer/packet"
)

const BUFFER_PACKET_REMOVED_ERROR = errors.New("BUFFER_PACKET_REMOVED_ERROR")
const BUFFER_INVALID_INDEX_ERROR = errors.New("BUFFER_INVALID_INDEX_ERROR")

// Buffer is simulated buffer for the sensor
type Buffer struct {
	packets []packet.Packet
	size    uint
}

func New(size uint) Buffer {
	return Buffeer{
		size:    size,
		packets: make([]packet.Packet, size),
	}
}

// Add adds the packet at index in the buffer
// if there is an another non-zero packet then it replaces it with current packet
func (b *Buffer) Add(p packet.Packet, index int) error {
	if index < int(b.size) && index >= 0 {
		if !b.packets[index].Exists() {
			b.packets[index] = p
			return nil
		}
		b.packets[index] = p
		return BUFFER_PACKET_REMOVED_ERROR
	}
	return BUFFER_INVALID_INDEX_ERROR
}

// Remove zeros the packet at the index
func (b *Buffer) Remove(index int) error {
	if index < int(b.size) && index >= 0 {
		b.packets[index].Zero()
		return nil
	}
	return BUFFER_INVALID_INDEX_ERROR
}
