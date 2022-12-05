package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
)

// cpu info
func getCpuInfo() {
	cpuInfos, err := cpu.Info()
	if err != nil {
		fmt.Printf("get cpu info failed, err:%v", err)
	}

	for _, ci := range cpuInfos {
		fmt.Println(ci)
	}

	// 获取CPU逻辑核心数
	cores, err := cpu.Counts(true)
	if err != nil {
		fmt.Printf("get cpu cores failed, err:%v\n", err)
		return
	}
	fmt.Println(cores)

	// CPU使用率
	for {
		percent, _ := cpu.Percent(time.Second, false)
		fmt.Printf("cpu percent:%v\n", percent)
	}
}

// cpu load
func getLoad() {
	info, err := load.Avg()
	if err != nil {
		fmt.Printf("get load failed, err:%v\n", err)
		return
	}
	fmt.Println(info)
	fmt.Println(info.Load1)
}

// memory info
func getMemInfo() {
	info, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("get mem info failed, err:%v\n", err)
		return
	}
	fmt.Println(info)
}

// host info
func getHostInfo() {
	info, err := host.Info()
	if err != nil {
		fmt.Printf("get host info failed, err:%v\n", err)
		return
	}
	fmt.Println(info)
}

// Disk info
func getDiskInfo() {
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get disk partitions failed, err:%v\n", err)
		return
	}

	for _, part := range parts {
		fmt.Printf("part:%v\n", part.String())
		diskInfo, _ := disk.Usage(part.Mountpoint)
		fmt.Printf("disk info:used:%v free:%v\n", diskInfo.UsedPercent, diskInfo.Free)
	}
}

// Network info
func getNetInfo() {
	info, err := net.IOCounters(true)
	if err != nil {
		fmt.Printf("get net io counters failed, err:%v\n", err)
		return
	}
	for index, v := range info {
		fmt.Printf("%v:%v send:%v recv:%v\n", index, v, v.BytesSent, v.BytesRecv)
	}
}

func main() {
	// getCpuInfo()
	// getLoad()
	// getMemInfo()
	// getHostInfo()
	getDiskInfo()
	// getNetInfo()
}
