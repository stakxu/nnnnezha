package utils

import (
	"fmt"
	"time"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/net"
)

// ANSI escape codes for colors
const (
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorReset  = "\033[0m"
)

// PrintTime 打印当前时间
func PrintTime() {
	fmt.Print("\033[u") // 移动光标至行首
	fmt.Print("\033[K") // 清除当前行

	fmt.Printf("Time: %s%s%s\n", colorYellow, time.Now().Format("2006-01-02 15:04:05"), colorReset)
}

// PrintSystemInfo 打印系统信息
func PrintSystemInfo() {
	// 获取 CPU 使用率
	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		fmt.Printf("Failed to get CPU usage: %v\n", err)
	} else {
		fmt.Printf("CPU Usage: %s%.2f%%%s\n", colorRed, cpuPercent[0], colorReset)
	}

	// 获取内存使用情况
	memInfo, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("Failed to get memory info: %v\n", err)
	} else {
		fmt.Printf("Memory Usage: %s%.2f%%%s\n", colorRed, memInfo.UsedPercent, colorReset)
	}

	// 获取网络上行速度
	netSpeed, err := GetNetworkSpeed()
	if err != nil {
		fmt.Printf("Failed to get network speed: %v\n", err)
	} else {
		fmt.Printf("Network Sent Speed: %s%.2f Mbps%s\n", colorRed, netSpeed, colorReset)
	}
}

// GetNetworkSpeed 获取网络上行速度
func GetNetworkSpeed() (float64, error) {
	netIO, err := net.IOCounters(false)
	if err != nil {
		return 0, err
	}
	var lastSent uint64
	var lastTime time.Time
	now := time.Now()

	if lastTime.IsZero() {
		lastTime = now
		lastSent = netIO[0].BytesSent
	} else {
		interval := now.Sub(lastTime).Seconds()
		netSpeed := float64(netIO[0].BytesSent-lastSent) / interval / 1024 / 1024 // Convert bytes to MB
		return netSpeed, nil
	}
	return 0, nil
}
