package main

import (
	"fmt"
	"net"
	"strings"
)

// GetOutBoundIP 获取本机对外的IP
func GetOutBoundIP() (ip string, err error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return
	}
	defer conn.Close()
	localAddr := conn.LocalAddr().(*net.UDPAddr)
	// fmt.Println(localAddr.String())
	ip = strings.Split(localAddr.IP.String(), ":")[0]
	return
}

func main() {
	ip, err := GetOutBoundIP()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ip)
}
