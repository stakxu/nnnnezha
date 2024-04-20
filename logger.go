// utils/utils.go

package utils

// ANSI escape codes for colors
const (
	ColorRed    = "\033[31m"
	ColorGreen  = "\033[32m"
	ColorYellow = "\033[33m"
	ColorReset  = "\033[0m"
)

// PrintTime prints the current time with color
func PrintTime() {
	fmt.Print("\033[u") // Move cursor to the beginning of the line
	fmt.Print("\033[K") // Clear current line

	fmt.Printf("Time: %s%s%s\n", ColorYellow, time.Now().Format("2006-01-02 15:04:05"), ColorReset)
}

// PrintSystemInfo prints system information with color
func PrintSystemInfo() {
	// Get CPU usage
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Printf("Failed to get CPU usage: %v\n", err)
	} else {
		fmt.Printf("CPU Usage: %s%.2f%%%s\n", utils.ColorRed, cpuPercent[0], utils.ColorReset)
	}

	// Get memory usage
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Failed to get memory info: %v\n", err)
	} else {
		fmt.Printf("Memory Usage: %s%.2f%%%s\n", utils.ColorRed, memInfo.UsedPercent, utils.ColorReset)
	}

	// Get network upload speed
	netSpeed, err := getNetworkSpeed()
	if err != nil {
		fmt.Printf("Failed to get network speed: %v\n", err)
	} else {
		fmt.Printf("Network Sent Speed: %s%.2f Mbps%s\n", utils.ColorRed, netSpeed, utils.ColorReset)
	}
}
