// Package buffer implements buffer for the sensor
package buffer

import (
	"errors"
	"flag"

	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer/packet"
)

// BUFFER_PACKET_REMOVED_ERROR when some packet is removed from the buffer
var BUFFER_PACKET_REMOVED_ERROR = errors.New("BUFFER_PACKET_REMOVED_ERROR")

// BUFFER_INVALID_INDEX_ERROR when when the index is out of range
var BUFFER_INVALID_INDEX_ERROR = errors.New("BUFFER_INVALID_INDEX_ERROR")

var buffer_size uint

func init() {
	flag.UintVar(&buffer_size, "buffer-size", 20, "sensor buffer size")
}

// Buffer is simulated buffer for the sensor
type Buffer struct {
	Packets []packet.Packet
	size    uint
}

func New() Buffer {
	return Buffer{
		size:    buffer_size,
		Packets: make([]packet.Packet, buffer_size),
	}
}

func (b *Buffer) InBuffer(pkt packet.Packet) bool {
	for _, p := range b.Packets {
		if pkt.Id == p.Id {
			return true
		}
	}
	return false
}

// Add adds the packet at index in the buffer
// if there is an another non-zero packet then it replaces it with current packet
func (b *Buffer) Add(p packet.Packet, index int) error {
	if index < int(b.size) && index >= 0 {
		if !b.Packets[index].Exists() {
			b.Packets[index] = p
			return nil
		}
		b.Packets[index] = p
		return BUFFER_PACKET_REMOVED_ERROR
	}
	return BUFFER_INVALID_INDEX_ERROR
}

// Remove zeros the packet at the index
func (b *Buffer) Remove(index int) error {
	if index < int(b.size) && index >= 0 {
		b.Packets[index].Zero()
		return nil
	}
	return BUFFER_INVALID_INDEX_ERROR
}
