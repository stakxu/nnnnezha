// network/speed.go
package network

import (
	"time"

	"github.com/shirou/gopsutil/net"
)
var lastTime time.Time
var lastSent uint64
func GetNetworkSpeed() (float64, error) {
	netIO, err := net.IOCounters(true)
	if err != nil {
		return 0, err
	}

	var totalSent uint64
	now := time.Now()

	for _, io := range netIO {
		totalSent += io.BytesSent
	}

	interval := now.Sub(lastTime).Seconds()
	netSpeed := float64(totalSent-lastSent) / interval / 1024 / 1024
	lastTime = now
	lastSent = totalSent

	return netSpeed, nil
}
