package main

/*
	#include <stdlib.h>
	#include "../PortOwnerCpp/PortOwner.h"

	#cgo LDFLAGS: -L../PortOwnerCpp -lPortOwner64

*/
import "C"
import (
	"fmt"
	"time"
	"unsafe"
)

type Port struct {
	Port uint32
	Pid  uint32
}

func GetUdp4Ports(maxCount int) []Port {
	var ports []Port = nil
	out := C.malloc((C.ulonglong)(maxCount)*8 + 4)
	result := C.GetUdp4Ports(out, 256*8+4)
	if result == 0 {
		buf := out
		count := int(*(*uint32)(buf))
		if count > maxCount {
			count = maxCount
		}
		ports = make([]Port, count)
		for i := 0; i < count; i++ {
			buf = unsafe.Add(buf, 4)
			port := *(*uint32)(buf)
			buf = unsafe.Add(buf, 4)
			pid := *(*uint32)(buf)
			ports[i] = Port{Port: port, Pid: pid}
		}
	}
	C.free(out)
	return ports
}

func GetUdp6Ports(maxCount int) []Port {
	var ports []Port = nil
	out := C.malloc((C.ulonglong)(maxCount)*8 + 4)
	result := C.GetUdp6Ports(out, 256*8+4)
	if result == 0 {
		buf := out
		count := int(*(*uint32)(buf))
		if count > maxCount {
			count = maxCount
		}
		ports = make([]Port, count)
		for i := 0; i < count; i++ {
			buf = unsafe.Add(buf, 4)
			port := *(*uint32)(buf)
			buf = unsafe.Add(buf, 4)
			pid := *(*uint32)(buf)
			ports[i] = Port{Port: port, Pid: pid}
		}
	}
	C.free(out)
	return ports
}

func GetTcp4Ports(maxCount int) []Port {
	var ports []Port = nil
	out := C.malloc((C.ulonglong)(maxCount)*8 + 4)
	result := C.GetTcp4Ports(out, 256*8+4)
	if result == 0 {
		buf := out
		count := int(*(*uint32)(buf))
		if count > maxCount {
			count = maxCount
		}
		ports = make([]Port, count)
		for i := 0; i < count; i++ {
			buf = unsafe.Add(buf, 4)
			port := *(*uint32)(buf)
			buf = unsafe.Add(buf, 4)
			pid := *(*uint32)(buf)
			ports[i] = Port{Port: port, Pid: pid}
		}
	}
	C.free(out)
	return ports
}

func GetTcp6Ports(maxCount int) []Port {
	var ports []Port = nil
	out := C.malloc((C.ulonglong)(maxCount)*8 + 4)
	result := C.GetTcp6Ports(out, 256*8+4)
	if result == 0 {
		buf := out
		count := int(*(*uint32)(buf))
		if count > maxCount {
			count = maxCount
		}
		ports = make([]Port, count)
		for i := 0; i < count; i++ {
			buf = unsafe.Add(buf, 4)
			port := *(*uint32)(buf)
			buf = unsafe.Add(buf, 4)
			pid := *(*uint32)(buf)
			ports[i] = Port{Port: port, Pid: pid}
		}
	}
	C.free(out)
	return ports
}

func GetPidName(pid uint32) *string {
	buf := C.malloc(512)
	size := C.GetProcessName((C.uint)(pid), (*C.char)(buf), 512)
	if size < 1 {
		return nil
	}
	str := C.GoString((*C.char)(buf))
	C.free(buf)
	return &str
}

func main() {
	//maxCount := 256
	start := time.Now()
	pName := GetPidName(18400)
	//ports := GetUdp4Ports(maxCount)
	end := time.Now()
	if pName != nil {
		fmt.Printf("App name = %s\n", *pName)
	} else {
		fmt.Println("NIL")
	}
	//if ports != nil {
	//	count := len(ports)
	//	fmt.Printf("Found %4d\n", count)
	//	for i := 0; i < count; i++ {
	//		fmt.Printf("PID %6d -> PORT %6d\n", ports[i].Pid, ports[i].Port)
	//	}
	//} else {
	//	fmt.Println("No ports found")
	//}
	dur := end.Sub(start)
	fmt.Printf("Time taken = %d ms\n", dur.Milliseconds())
}
