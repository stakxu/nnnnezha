// monitor/system_info.go
package monitor

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
)

const (
	colorRed    = "\033[31m"
	colorReset  = "\033[0m"
)

func PrintSystemInfo() {
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Printf("Failed to get CPU usage: %v\n", err)
	} else {
		fmt.Printf("CPU Usage: %s%.2f%%%s\n", colorRed, cpuPercent[0], colorReset)
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Failed to get memory info: %v\n", err)
	} else {
		fmt.Printf("Memory Usage: %s%.2f%%%s\n", colorRed, memInfo.UsedPercent, colorReset)
	}
}
