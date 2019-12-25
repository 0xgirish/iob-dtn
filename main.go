package main

import (
	"flag"
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
	var simulationTime int
	flag.IntVar(&simulationTime, "env-time", 30, "simulation time in second")
	flag.Parse()

	policies := []policy.Policy{
		policy.New(policy.KONP_POLICY),
		policy.New(policy.NP_POLICY),
		policy.New(policy.GPP_POLICY),
		policy.New(policy.LC_POLICY),
	}

	names := []string{"konp", "np", "gpp", "lc"}

	for i, p := range policies {
		RunSimulationWithPolicy(p, simulationTime, names[i])
	}

}

func RunSimulationWithPolicy(p policy.Policy, simulationTime int, name string) {
	manager.Erase()

	spos := []util.Position{
		util.Position{X: 0, Y: 0},
		util.Position{X: 25, Y: 25},
		util.Position{X: 0, Y: 25},
		util.Position{X: 25, Y: 0},
		util.Position{X: 5, Y: 15},
		util.Position{X: 12, Y: 25},
		util.Position{X: 25, Y: 12},
		util.Position{X: 13, Y: 6},
		util.Position{X: 10, Y: 14},
	}

	stop := make(chan bool)
	var sensors []sensor.Sensor

	for i := 1; i <= len(spos)*int(env.Num_cycles_per_station); i++ {
		manager.Register(i)
		sensors = append(sensors, sensor.New(i, p))
	}

	envr := env.New(spos, sensors, stop)
	envr.StartSimulation()

	time.Sleep(time.Duration(simulationTime) * time.Second)
	for i := 1; i <= 7*int(env.Num_cycles_per_station); i++ {
		stop <- true
	}

	filename := fmt.Sprintf("result.%s.csv", name)
	file, err := os.Create(filename)
	if err != nil {
		log.Printf("can not create %s\n", filename)
		return
	}

	defer file.Close()
	fmt.Fprint(file, manager.Manager)
	fmt.Printf("result for %s written\n", name)
}
