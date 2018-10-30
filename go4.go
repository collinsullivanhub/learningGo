package main 

import (
"fmt"
"time"
)

//Device Structs

type IOSRouterStruct struct{
	hostname string
	ipAddress string
	managementInterface string
	poweredOn bool
	codeVersion string
}

type IOSXrRouterStruct struct{
	hostname string
	ipAddress string
	managementInterface string
	poweredOn bool
	codeVersion string
}

type IOSxeRouterStruct struct{
	hostname string
	ipAddress string
	managementInterface string
	poweredOn bool
	codeVersion string
}

type NexusRouterStruct struct{
	hostname string
	ipAddress string
	managementInterface string
	poweredOn bool
	codeVersion string
}

type IOSSwitchStruct struct{
	hostname string
	ipAddress string
	managementInterface string
	poweredOn bool
	codeVersion string
}

type IOSXrSwitchStruct struct{
	hostname string
	ipAddress string
	managementInterface string
	poweredOn bool
	codeVersion string
}

type IOSxeSwitchStruct struct{
	hostname string
	ipAddress string
	managementInterface string
	poweredOn bool
	codeVersion string
}

type NexusSwtichStruct struct{
	hostname string
	ipAddress string
	managementInterface string
	poweredOn bool
	codeVersion string
}

func main() {

	nexusrouter1 := NexusRouterStruct{hostname: "AS-RTP-TS-R-114",
	ipAddress: "10.122.24.36",
	managementInterface: "mgmt0",
	poweredOn:true,
	codeVersion:"19.4.23"}

	fmt.Println(nexusrouter1.hostname)

	IpAddressList := []string{
	"10.122.16.142",
	"10.122.15.133",
	"10.122.12.142",
	"10.122.21.224"}

	fmt.Println(time.Now().Format(time.RFC850))
	fmt.Println(IpAddressList)

	//infinite loop
	//for {
	//	fmt.Println("hi")
	//}
	
}




