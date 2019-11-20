// Package comdevice simulate comunication device
package comdevice

// msgType insures that
// the user can only send 'ACK' and 'PacketTrasfer' type messages
type msgType string

const ACK = msgType("ACK")
const PacketTransfer = msgType("PacketTransfer")

// Comdevice is the concept of a device for the sensor
type Comdevice interface {
	Send(msg Message, dvc Comdevice)
	Receive(msg Message)
}

// Message describes the underlying message and its type
type Message struct {
	Type msgType
	Msg  interface{}
	From int
	To   int
}
