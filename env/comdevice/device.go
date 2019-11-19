package comdevice

// device is the physical hardware for the sensor
type device struct {
	receiver chan Message
}

// Receive sends the message to device's receiver channel
func (d device) Receive(msg Message) {
	d.reveiver <- msg
}

// Send sends the message to device's receiver
func (d device) Send(msg Message, dvc Comdevice) {
	dvc.Receive(msg)
}

// New Comdevice for a sensor
func New() Comdevice {
	return device {
		receiver: make(chan Message)
	}
}
