// Package env provide environment simulation
package env

import (
	"flag"
	"math/rand"

	"github.com/zkmrgirish/iob-dtn/env/bicycle"
	"github.com/zkmrgirish/iob-dtn/env/comdevice"
	"github.com/zkmrgirish/iob-dtn/env/sensor"
	"github.com/zkmrgirish/iob-dtn/env/station"
	"github.com/zkmrgirish/iob-dtn/env/util"
)

var (
	DeviceRange            float64
	Num_cycles_per_station uint
)

func init() {
	flag.Float64Var(&DeviceRange, "range", 3.0, "bike sensor range")
	flag.UintVar(&Num_cycles_per_station, "num-cycles", 8, "number of cycles per station")
	flag.Parse()
}

type env struct {
	bicycles []bicycle.Bicycle
	stations []station.Station
}

// RequestDestination is required by the bicycle
func (e env) RequestDestination() util.Position {
	ind := rand.Int() % len(e.stations)
	return e.stations[ind].GetPosition()
}

// Range is required by the bicycle
func (e env) Range(pos util.Position) []comdevice.Comdevice {
	for _, s := range e.stations {
		if pos.Distance(s.GetPosition()) <= DeviceRange {
			return []comdevice.Comdevice{
				s.Comdevice,
			}
		}
	}

	var devices []comdevice.Comdevice
	for _, bc := range e.bicycles {
		if bc.Moving() && pos.Distance(bc.GetPosition()) <= DeviceRange {
			devices = append(devices, bc.Comdevice)
		}
	}
	return devices
}

func (e *env) StartSimulation() {
	for i := 0; i < len(e.bicycles); i++ {
		go e.bicycles[i].Initiate()
	}
}

func New(spos []util.Position, sensors []sensor.Sensor, stop chan bool) env {
	ssize := len(spos)
	bicycles := make([]bicycle.Bicycle, ssize*int(Num_cycles_per_station))
	stations := make([]station.Station, ssize)

	envr := env{
		bicycles: bicycles,
		stations: stations,
	}

	for i, pos := range spos {
		envr.stations[i] = station.New(pos)
		for k := 0; k < int(Num_cycles_per_station); k++ {
			ind := i*int(Num_cycles_per_station) + k
			envr.bicycles[ind] = bicycle.New(pos, sensors[ind], envr, stop)
		}
	}

	return envr
}
