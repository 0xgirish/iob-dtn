// Package manager implements manager for recording generation and deliveries of packets
package manager

import (
	"fmt"
	"sync"
)

var Manager = manager{
	successCount: make(map[int]int),
	packetCount:  make(map[int]int),
	delivered:    make(map[int]bool),
}

// manager keep the count of successfully delivered packets and total number of generated packets for a parent
// and set of delivered packets. manager implements stringer interface
type manager struct {
	successCount map[int]int
	packetCount  map[int]int
	delivered    map[int]bool
	mux          sync.Mutex
}

// Register the parent for deliveries
func Register(parent int) {
	Manager.mux.Lock()
	defer Manager.mux.Unlock()

	Manager.packetCount[parent] = 0
}

// IncrCunter increases the total packetCount for parent by one
func IncrCounter(parent int) {
	Manager.mux.Lock()
	defer Manager.mux.Unlock()

	Manager.packetCount[parent]++
}

// MarkSuccess if the packet have not been delivered before
func MarkSuccess(parent, pkt int) {
	Manager.mux.Lock()
	defer Manager.mux.Unlock()

	if !Manager.delivered[pkt] {
		Manager.successCount[parent]++
		Manager.delivered[pkt] = true
	}
}

// String returns the manager data of delivered packets of bicycles in string form
// e.g. fmt.Fprint(os.Create("result.txt"), manager)
func (m manager) String() string {
	result := "id,success,packetCount"
	for parent, packetCount := range m.packetCount {
		result += fmt.Sprintf("%d,%d,%d\n", parent, m.successCount[parent], packetCount)
	}
	return result
}
