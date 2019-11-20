// Package station provides simulation for stations for bicycles
package station

import (
	"sync"

	"github.com/zkmrgirish/iob-dtn/env/comdevice"
	"github.com/zkmrgirish/iob-dtn/env/util"
)

var (
	id  int
	mux sync.Mutex
)

type Station struct {
	Id       int
	pos      util.Position
	receiver chan comdevice.Message
	comdevice.Comdevice
}

func New(pos util.Position) Station {
	mux.Lock()
	defer mux.Unlock()

	id--
	receiver := make(chan comdevice.Message)
	return Station{
		Id:        id,
		pos:       pos,
		receiver:  receiver,
		Comdevice: comdevice.NewStationDevice(id, &pos),
	}
}

func (s Station) GetPosition() util.Position {
	return s.pos
}
