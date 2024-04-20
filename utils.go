package utils

import (
	"fmt"
	"math/rand"
	"time"
)

// ANSI escape codes for colors
const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorReset  = "\033[0m"
)

// RandomizeStatus randomizes the values in the state within 10%
func RandomizeStatus(status *proto.State) {
	RandomizeFloat64(&status.Cpu)
	RandomizeFloat64(&status.Load1)
	RandomizeFloat64(&status.Load5)
	RandomizeFloat64(&status.Load15)
	RandomizeUint64(&status.MemUsed)
	//	RandomizeUint64(&status.NetInTransfer)
	//	RandomizeUint64(&status.NetOutTransfer)
	RandomizeUint64(&status.NetInSpeed)
	RandomizeUint64(&status.NetOutSpeed)
	RandomizeUint64(&status.ProcessCount)
	RandomizeUint64(&status.SwapUsed)
	RandomizeUint64(&status.TcpConnCount)
	RandomizeUint64(&status.UdpConnCount)
}

// RandomizeFloat64 randomizes a float64 value within 10%
func RandomizeFloat64(val *float64) {
	*val = *val * (1 + (rand.Float64()*0.2 - 0.1)) // 10% range of random fluctuation
}

// RandomizeUint64 randomizes a uint64 value within 10%
func RandomizeUint64(val *uint64) {
	*val = uint64(float64(*val) * (1 + (rand.Float64()*0.2 - 0.1))) // 10% range of random fluctuation
}

// PrintTime prints the current time
func PrintTime() {
	fmt.Print("\033[u") // Move cursor to beginning of line
	fmt.Print("\033[K") // Clear current line

	fmt.Printf("Time: %s%s%s\n", ColorYellow, time.Now().Format("2006-01-02 15:04:05"), ColorReset)
}

// PrintSystemInfo prints system information
func PrintSystemInfo() {
	// Print CPU usage
	cpuPercent := 10.0 // Placeholder value
	fmt.Printf("CPU Usage: %s%.2f%%%s\n", ColorRed, cpuPercent, ColorReset)

	// Print memory usage
	memUsedPercent := 20.0 // Placeholder value
	fmt.Printf("Memory Usage: %s%.2f%%%s\n", ColorRed, memUsedPercent, ColorReset)

	// Print network speed
	netSpeed := 30.0 // Placeholder value
	fmt.Printf("Network Sent Speed: %s%.2f Mbps%s\n", ColorRed, netSpeed, ColorReset)
}


// RandomizeFloat64 randomizes a float64 value within 10%
func RandomizeFloat64(val *float64) {
	*val = *val * (1 + (rand.Float64()*0.2 - 0.1)) // 10% range of random fluctuation
}

// RandomizeUint64 randomizes a uint64 value within 10%
func RandomizeUint64(val *uint64) {
	*val = uint64(float64(*val) * (1 + (rand.Float64()*0.2 - 0.1))) // 10% range of random fluctuation
}
