package main

import (
	"fmt"
	"time"

	client "github.com/influxdata/influxdb1-client/v2"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/net"
)

// 获取系统资源使用情况，存入到influxdb中

var (
	cli                    client.Client
	lastNetIoStatTimeStamp int64    // 上一次获取网络IO数据的时间点
	lastNetInfo            *NetInfo // 上一次的网络IO数据
)


func initConnInflux() (err error) {
	cli, err = client.NewHTTPClient(client.HTTPConfig{
		Addr:     "http://120.24.222.91:8086",
		Username: "admin",
		Password: "admin@123",
	})
	return
}

// cpu info
func getCpuInfo() {
	var cpuInfo = new(CpuInfo)
	// CPU 使用率
	percent, _ := cpu.Percent(time.Second, false)
	fmt.Printf("cpu percent:%v\n", percent)
	// 写入到Influxdb
	cpuInfo.CpuPercent = percent[0]
	writesCPUPoints(cpuInfo)

}

// memory info
func getMemInfo() {
	var memInfo = new(MemInfo)
	info, err := mem.VirtualMemory()
	if err != nil {
		fmt.Printf("get mem info failed, err:%v\n", err)
		return
	}
	memInfo.Total = info.Total
	memInfo.Available = info.Available
	memInfo.Used = info.Used
	memInfo.UsedPercent = info.UsedPercent
	memInfo.Free = info.Free
	memInfo.Buffers = info.Buffers
	memInfo.Cached = info.Cached
	writesMemPoints(memInfo)
}

func getDiskInfo() {
	var diskInfo = &DiskInfo{
		PartitionUsageStat: make(map[string]*disk.UsageStat, 32),
	}
	parts, err := disk.Partitions(true)
	if err != nil {
		fmt.Printf("get disk partitions failed, err:%v\n", err)
		return
	}
	for _, part := range parts {
		// 拿到每一个分区
		usageStat, err := disk.Usage(part.Mountpoint) // 传挂载点进去
		if err != nil {
			fmt.Printf("get %s usage stat failed, err:%v\n", part.Mountpoint, err)
			continue
		}
		diskInfo.PartitionUsageStat[part.Mountpoint] = usageStat
	}
	writesDiskPoints(diskInfo)
}

// Network info
func getNetInfo() {
	var netInfo = &NetInfo{
		NetIOCountersStat: make(map[string]*IOStat, 16),
	}
	currentTimeStamp := time.Now().Unix()
	netIos, err := net.IOCounters(true)
	if err != nil {
		fmt.Printf("get net io counters failed, err:%v\n", err)
		return
	}
	for _, netIO := range netIos {
		var ioStat = new(IOStat)
		ioStat.BytesRecv = netIO.BytesRecv
		ioStat.BytesSent = netIO.BytesSent
		ioStat.PacketsSent = netIO.PacketsSent
		ioStat.PacketsRecv = netIO.PacketsRecv
		// 将具体网卡数据的ioStat变量添加到map中
		netInfo.NetIOCountersStat[netIO.Name] = ioStat // 不要放到continue下面
		// 开始计算网卡相关速率
		if lastNetIoStatTimeStamp == 0 || lastNetInfo == nil {
			continue
		}
		// 计算时间间隔
		interval := currentTimeStamp - lastNetIoStatTimeStamp
		// 计算速率
		ioStat.BytesSentRate = float64((ioStat.BytesSent - lastNetInfo.NetIOCountersStat[netIO.Name].BytesSent)) / float64(interval)
		ioStat.BytesRecvRate = float64((ioStat.BytesRecv - lastNetInfo.NetIOCountersStat[netIO.Name].BytesRecv)) / float64(interval)
		ioStat.PacketsSentRate = float64((ioStat.PacketsSent - lastNetInfo.NetIOCountersStat[netIO.Name].PacketsSent)) / float64(interval)
		ioStat.PacketsRecvRate = float64((ioStat.PacketsRecv - lastNetInfo.NetIOCountersStat[netIO.Name].PacketsRecv)) / float64(interval)

	}
	// 更新全局记录的上一次采集网卡的时间点和数据
	lastNetIoStatTimeStamp = currentTimeStamp // 更新时间
	lastNetInfo = netInfo

	// 发送至Influxdb
	writesNetPoints(netInfo)
}

func run(interval time.Duration) {
	ticker := time.Tick(interval)
	for range ticker {
		getCpuInfo()
		getMemInfo()
		getDiskInfo()
		getNetInfo()
		// fmt.Println(t)
	}
}

func main() {
	// 连接influxdb
	err := initConnInflux()
	if err != nil {
		fmt.Printf("connect to influxdb failed, err:%v\n", err)
		return
	}
	run(time.Second)
	
}
