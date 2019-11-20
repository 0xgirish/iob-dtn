// Package station provides simulation for stations for bicycles
package station

import (
	"log"
	"sync"

	"github.com/zkmrgirish/iob-dtn/env/comdevice"
	"github.com/zkmrgirish/iob-dtn/env/manager"
	"github.com/zkmrgirish/iob-dtn/env/sensor/buffer/packet"
	"github.com/zkmrgirish/iob-dtn/env/util"
)

var (
	id  int
	mux sync.Mutex
)

type station struct {
	Id       int
	pos      util.Position
	receiver chan comdevice.Message
	comdevice.Comdevice
}

func New(pos util.Position) station {
	mux.Lock()
	defer mux.Unlock()

	id--
	receiver := make(chan comdevice.Message)
	return station{
		Id:        id,
		pos:       pos,
		receiver:  receiver,
		Comdevice: comdevice.New(receiver),
	}
}

func (s station) Listen() {
	for {
		select {
		case msg := <-s.receiver:
			if msg.To != s.Id && msg.Type != comdevice.PacketTransfer {
				continue
			}

			pkt, ok := msg.Msg.(packet.Packet)
			if !ok {
				log.Printf("message type shoud be packet.Packet with message singal: %s\n", msg.Type)
				continue
			}
			manager.MarkSuccess(pkt.Parent_id, pkt.Id)
			// TODO: send ACK to msg.From
		}
	}
}
