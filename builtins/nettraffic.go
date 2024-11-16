package builtins

import (
	"dwl_asyncbar/pkg"
	"fmt"
	"github.com/shirou/gopsutil/v3/net"
	"time"
)

var prevRx uint64
var prevTx uint64

func formatBytes(value uint64) string {
	if value > 1024*1024 {
		return fmt.Sprintf("%.1fM", float64(value)/1024/1024)
	}
	if value > 1024 {
		return fmt.Sprintf("%.1fK", float64(value)/1024)
	}
	return fmt.Sprintf("%d", value)
}

func netTrafficString() string {
	ioCounters, err := net.IOCounters(false)
	if err != nil || len(ioCounters) == 0 {
		return "↓   0B ↑   0B "
	}

	currentRx := ioCounters[0].BytesRecv
	currentTx := ioCounters[0].BytesSent

	// Calculate the difference since the last call
	rxDelta := currentRx - prevRx
	txDelta := currentTx - prevTx

	// Update previous values for the next interval
	prevRx = currentRx
	prevTx = currentTx

	// Fixed width formatting with padding
	return fmt.Sprintf("↓%6sB ↑%6sB", formatBytes(rxDelta), formatBytes(txDelta))
}

// NetTrafficModule returns a prebuilt NetTraffic module.
//
// Displays:  ↓4B ↑4B
func NetTrafficModule() pkg.Module {
	return pkg.Module{
		Name:     "NetTraffic",
		Interval: 1 * time.Second,
		Update:   netTrafficString,
	}
}
