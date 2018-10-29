package main 

import (
"fmt"
"time"
"net"
"sort"
)

//Device Structs

type IOSRouterStruct {
	hostname string
	ipAddress string
	managementInterface string
	poweredOn bool
	codeVersion string
}

type IOSXrRouterStruct {
	hostname string
	ipAddress string
	managementInterface string
	poweredOn bool
	codeVersion string
}

type IOSxeRouterStruct {
	hostname string
	ipAddress string
	managementInterface string
	poweredOn bool
	codeVersion string
}

type NexusRouterStruct {
	hostname string
	ipAddress string
	managementInterface string
	poweredOn bool
	codeVersion string
}

type IOSSwitchStruct {
	hostname string
	ipAddress string
	managementInterface string
	poweredOn bool
	codeVersion string
}

type IOSXrSwitchStruct {
	hostname string
	ipAddress string
	managementInterface string
	poweredOn bool
	codeVersion string
}

type IOSxeSwitchStruct {
	hostname string
	ipAddress string
	managementInterface string
	poweredOn bool
	codeVersion string
}

type NexusSwtichStruct {
	hostname string
	ipAddress string
	managementInterface string
	poweredOn bool
	codeVersion string
}


nexusrouter1 := NexusRouterStruct{
	name: "AS-RTP-TS-R-114"
	ipAddress: "10.122.24.36"
	managementInterface: "mgmt0"
	poweredOn:True
}

IpAddressList := []string{
	"10.122.16.142"
	"10.122.15.133"
	"10.122.12.142"
	"10.122.21.224"
}

func main() {

	fmt.Println(time.Now().Format(time.RFC850))

	byteIP := make([]net.IP, 0, len(IpAddressList))

	//IP String to bytes
	
	for _, ip := range IpAddressList {
		byteIP = append(byteIP, net.ParseIP(ip))
	}

	sort.Slice(byteIP, func(i, j int) bool {
		return bytes.Compare(byteIP[i], byteIP[j]) < 0
		})
	for _, ip := range byteIP{
		fmt.Println("%s\n", ip)
	}
}
