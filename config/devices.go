package config

import (
	"github.com/StackExchange/wmi"
	"net"
	"os"
	"strings"
	"syscall"
	"unsafe"
)

//cpux信息
type CpuInfo struct {
	Name          string
	NumberOfCores uint32
	ThreadCount   uint32
}

//gpu信息
type GpuInfo struct {
	Name string
}

//操作系统
type OsInfo struct {
	Name    string
	Version string
}

//内存信息
type MemoryInfo struct {
	cbSize                  uint32
	dwMemoryLoad            uint32
	ullTotalPhys            uint64 // in bytes 内存大小
	ullAvailPhys            uint64
	ullTotalPageFile        uint64
	ullAvailPageFile        uint64
	ullTotalVirtual         uint64
	ullAvailVirtual         uint64
	ullAvailExtendedVirtual uint64
}

//网卡信息
type NetworkInfo struct {
	Name       string
	IP         string
	MACAddress string
	RealIp     string
}

//磁盘信息
type Storage struct {
	Name       string
	FileSystem string
	Total      uint64
	Free       uint64
}

type storageInfo struct {
	Name       string
	Size       uint64
	FreeSpace  uint64
	FileSystem string
}

type Devices struct {
	CpuInfo     []CpuInfo   `json:"cpu_info"`
	GpuInfo     []GpuInfo   `json:"gpu_info"`
	OsInfo      []OsInfo    `json:"os_info"`
	MemoryInfo  MemoryInfo  `json:"memory_info"`
	NetworkInfo NetworkInfo `json:"network_info"`
	Storage     []Storage   `json:"storage"`
	UserName    string      `json:"user_name"`
	HostName    string      `json:"host_name"`
}

type IpGetter interface {
	GetRealIp() (string, error)
}

var ipGetter IpGetter

func (self *Devices) DevicesInit(getter IpGetter) {
	ipGetter = getter
	self.CpuInfo = self.getCPUInfo()
	self.GpuInfo = self.getGPUInfo()
	self.OsInfo = self.getOSInfo()
	self.MemoryInfo = self.getMemoryInfo()
	self.NetworkInfo = self.getNetworkInfo()
	self.Storage = self.getStorageInfo()
	self.UserName = self.getUserName()
	self.HostName = self.getHostName()
}

func (self *Devices) getCPUInfo() []CpuInfo {

	var cpuinfo []CpuInfo

	err := wmi.Query("Select * from Win32_Processor", &cpuinfo)
	if err != nil {
		return nil
	}
	//fmt.Printf("Cpu info =", cpuinfo)
	return cpuinfo
}

func (self *Devices) getGPUInfo() []GpuInfo {

	var gpuinfo []GpuInfo
	err := wmi.Query("Select * from Win32_VideoController", &gpuinfo)
	if err != nil {
		return nil
	}
	//fmt.Printf("GPU:=", gpuinfo[0].Name)
	return gpuinfo
}

func (self *Devices) getOSInfo() []OsInfo {
	var osInfo []OsInfo
	err := wmi.Query("Select * from Win32_OperatingSystem", &osInfo)
	if err != nil {
		return nil
	}
	//fmt.Printf("OS info =", osInfo)
	return osInfo
}

func (self *Devices) getMemoryInfo() MemoryInfo {
	var kernel = syscall.NewLazyDLL("Kernel32.dll")

	GlobalMemoryStatusEx := kernel.NewProc("GlobalMemoryStatusEx")
	var memInfo MemoryInfo
	memInfo.cbSize = uint32(unsafe.Sizeof(memInfo))
	_, _, _ = GlobalMemoryStatusEx.Call(uintptr(unsafe.Pointer(&memInfo)))
	//fmt.Printf("total=:", memInfo.ullTotalPhys)
	//fmt.Printf("free=:", memInfo.ullAvailPhys)
	//fmt.Printf("mem=:", mem)
	return memInfo
}

func (self *Devices) getNetworkInfo() NetworkInfo {
	intf, err := net.Interfaces()
	var network NetworkInfo
	if ipGetter != nil {
		network.RealIp, err = ipGetter.GetRealIp()
	}
	if err != nil {
		//fmt.Printf("get network info failed: ", err)
		return network
	}
	for _, v := range intf {
		addrs, err := v.Addrs()
		if err != nil {
			//fmt.Printf("get network addr failed: ", err)
			return network
		}
		//此处过滤loopback（本地回环）和isatap（isatap隧道）
		if !strings.Contains(v.Name, "Loopback") && !strings.Contains(v.Name, "isatap") {
			network.Name = v.Name
			network.MACAddress = v.HardwareAddr.String()
			for _, addr := range addrs {
				ip := getIpFromAddr(addr)
				if ip == nil {
					continue
				}
				network.IP = ip.String()

				return network
			}
			//fmt.Printf("network:=", network)
		}
	}
	return network
}

func getIpFromAddr(addr net.Addr) net.IP {
	var ip net.IP
	switch v := addr.(type) {
	case *net.IPNet:
		ip = v.IP
	case *net.IPAddr:
		ip = v.IP
	}
	if ip == nil || ip.IsLoopback() {
		return nil
	}
	ip = ip.To4()
	if ip == nil {
		return nil // not an ipv4 address
	}

	return ip
}

func (self *Devices) getStorageInfo() []Storage {
	var storageinfo []storageInfo
	var localStorages []Storage
	err := wmi.Query("Select * from Win32_LogicalDisk", &storageinfo)
	if err != nil {
		return localStorages
	}

	for _, storage := range storageinfo {
		info := Storage{
			Name:       storage.Name,
			FileSystem: storage.FileSystem,
			Total:      storage.Size,
			Free:       storage.FreeSpace,
		}
		localStorages = append(localStorages, info)
	}

	//fmt.Printf("localStorages:=", localStorages)

	return localStorages
}

func (self *Devices) getUserName() string {
	var size uint32 = 128
	var buffer = make([]uint16, size)
	user := syscall.StringToUTF16Ptr("USERNAME")
	domain := syscall.StringToUTF16Ptr("USERDOMAIN")
	r, err := syscall.GetEnvironmentVariable(user, &buffer[0], size)
	if err != nil {
		return ""
	}
	buffer[r] = '@'
	old := r + 1
	if old >= size {
		return syscall.UTF16ToString(buffer[:r])
	}
	r, err = syscall.GetEnvironmentVariable(domain, &buffer[old], size-old)
	return syscall.UTF16ToString(buffer[:old+r])
}

func (self *Devices) getHostName() string {
	ret := ""
	hostname, err := os.Hostname()
	if err == nil {
		ret += hostname
	}
	homedir, err := os.UserHomeDir()
	if err == nil {
		ret += "__"
		ret += homedir
	}

	return ret
}
