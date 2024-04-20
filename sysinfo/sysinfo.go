package sysinfo

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
	"time"
)

// PrintTime prints the current time to the console
func PrintTime() {
	fmt.Printf("Time: %s\n", time.Now().Format("2006-01-02 15:04:05"))
}

// PrintSystemInfo prints CPU usage, memory usage, and network speed to the console
func PrintSystemInfo() {
	cpuPercent, _ := cpu.Percent(time.Second, false)
	memInfo, _ := mem.VirtualMemory()
	netSpeed, _ := getNetworkSpeed()

	fmt.Printf("CPU Usage: %.2f%%\n", cpuPercent[0])
	fmt.Printf("Memory Usage: %.2f%%\n", memInfo.UsedPercent)
	fmt.Printf("Network Sent Speed: %.2f Mbps\n", netSpeed)
}

// getNetworkSpeed returns the network sending speed in Mbps
func getNetworkSpeed() (float64, error) {
	netIO, err := net.IOCounters(false)
	if err != nil {
		return 0, err
	}

	now := time.Now()
	interval := now.Sub(lastTime).Seconds()
	netSpeed := float64(netIO[0].BytesSent-lastSent) / interval / 1024 / 1024 // Convert bytes to MB
	lastSent = netIO[0].BytesSent
	lastTime = now

	return netSpeed, nil
}
