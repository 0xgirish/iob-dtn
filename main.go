package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/zkmrgirish/iob-dtn/env"
	"github.com/zkmrgirish/iob-dtn/env/manager"
	"github.com/zkmrgirish/iob-dtn/env/policy"
	"github.com/zkmrgirish/iob-dtn/env/sensor"
	"github.com/zkmrgirish/iob-dtn/env/util"
)

func main() {
	spos := []util.Position{
		util.Position{X: 0, Y: 0},
		util.Position{X: 25, Y: 25},
		util.Position{X: 5, Y: 15},
		util.Position{X: 15, Y: 6},
		util.Position{X: 10, Y: 14},
		util.Position{X: 0, Y: 25},
		util.Position{X: 25, Y: 0},
	}

	stop := make(chan bool)
	var sensors []sensor.Sensor

	for i := 1; i <= 7*int(env.Num_cycles_per_station); i++ {
		sensors = append(sensors, sensor.New(i, policy.New(policy.LC_POLICY)))
	}

	envr := env.New(spos, sensors, stop)
	envr.StartSimulation()

	time.Sleep(10 * time.Second)

	file, err := os.Create("result_lc.txt")
	if err != nil {
		log.Println("can not create result_lc.txt")
		return
	}

	defer file.Close()
	fmt.Fprint(file, manager.Manager)
}
