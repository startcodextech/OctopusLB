package system

import (
	"errors"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/disk"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/load"
	"github.com/shirou/gopsutil/v4/mem"
	"github.com/shirou/gopsutil/v4/net"
	"github.com/shirou/gopsutil/v4/process"
	"math"
)

var (
	netPrevBytesSent = uint64(0)
	netPrevBytesRecv = uint64(0)

	ErrUnsupportedSystemFamily = errors.New("Unsupported system family")
)

const (
	Debian OSFamily = "debian"
	Ubuntu OSFamily = "ubuntu"
	Rhel   OSFamily = "rhel"
)

type (
	OSFamily string

	InformationUnit struct {
		Value float64 `json:"value"`
		Unit  string  `json:"unit"`
	}

	ResourceConsumptionRam struct {
		Total       InformationUnit `json:"total"`
		Used        InformationUnit `json:"used"`
		Free        InformationUnit `json:"free"`
		Share       InformationUnit `json:"share"`
		Buffers     InformationUnit `json:"buffers"`
		Cached      InformationUnit `json:"cached"`
		Available   InformationUnit `json:"available"`
		UsedPercent float64         `json:"used_percent"`
	}

	ResourceConsumptionDisk struct {
		Total       InformationUnit `json:"total"`
		Used        InformationUnit `json:"used"`
		Free        InformationUnit `json:"free"`
		UsedPercent float64         `json:"used_percent"`
	}

	ResourceConsumptionNet struct {
		Sent     InformationUnit `json:"sent"`
		Received InformationUnit `json:"received"`
	}

	ResourceConsumption struct {
		RAM  ResourceConsumptionRam  `json:"ram"`
		CPU  float64                 `json:"cpu"`
		Disk ResourceConsumptionDisk `json:"disk"`
		Net  ResourceConsumptionNet  `json:"net"`
	}

	SysInfo struct {
		host *host.InfoStat
		cpu  []cpu.InfoStat
		mem  *mem.VirtualMemoryStat
		disk *disk.UsageStat
		proc []*process.Process
		net  []net.InterfaceStat
		load *load.AvgStat

		Os                 string          `json:"os"`
		Platform           string          `json:"platform"`
		PlatformFamily     OSFamily        `json:"platform_family"`
		PlatformVersion    string          `json:"platform_version"`
		Architecture       string          `json:"architecture"`
		Virtualization     string          `json:"virtualization"`
		VirtualizationRole string          `json:"virtualization_role"`
		CPUs               int             `json:"cpus"`
		CPUModel           string          `json:"cpu_model"`
		RAM                InformationUnit `json:"ram"`
		DiskSize           InformationUnit `json:"disk_size"`
	}
)

func Get() (*SysInfo, error) {
	hostInfo, _ := host.Info()
	cpuInfo, _ := cpu.Info()
	menInfo, _ := mem.VirtualMemory()
	diskInfo, _ := disk.Usage("/")
	procInfo, _ := process.Processes()
	netInfo, _ := net.Interfaces()
	loadInfo, _ := load.Avg()

	osFamily := OSFamily(hostInfo.PlatformFamily)

	validOsFamily := map[OSFamily]bool{
		Debian: true,
		Ubuntu: true,
		Rhel:   true,
	}

	if !validOsFamily[OSFamily(osFamily)] {
		return nil, ErrUnsupportedSystemFamily
	}

	info := &SysInfo{
		host: hostInfo,
		cpu:  cpuInfo,
		mem:  menInfo,
		disk: diskInfo,
		proc: procInfo,
		net:  netInfo,
		load: loadInfo,

		Os:                 hostInfo.OS,
		Platform:           hostInfo.Platform,
		PlatformFamily:     osFamily,
		PlatformVersion:    hostInfo.PlatformVersion,
		Architecture:       hostInfo.KernelArch,
		Virtualization:     hostInfo.VirtualizationSystem,
		VirtualizationRole: hostInfo.VirtualizationRole,
		CPUs:               len(cpuInfo),
		CPUModel:           cpuInfo[0].ModelName,
		RAM:                BytesToUnit(menInfo.Total),
		DiskSize:           BytesToUnit(diskInfo.Total),
	}

	return info, nil
}

func (s *SysInfo) CurrentResourcesConsumption() ResourceConsumption {
	cpuPercent, _ := cpu.Percent(0, false)
	netCounter, _ := net.IOCounters(false)

	netSent := netPrevBytesSent
	netRecv := netPrevBytesRecv

	if netCounter[0].BytesSent > netPrevBytesSent {
		netPrevBytesSent = netCounter[0].BytesSent
	}
	if netCounter[0].BytesRecv > netPrevBytesRecv {
		netPrevBytesRecv = netCounter[0].BytesRecv
	}

	return ResourceConsumption{
		RAM: ResourceConsumptionRam{
			Total:       BytesToUnit(s.mem.Total),
			Used:        BytesToUnit(s.mem.Used),
			Free:        BytesToUnit(s.mem.Free),
			Available:   BytesToUnit(s.mem.Available),
			Share:       BytesToUnit(s.mem.Shared),
			Buffers:     BytesToUnit(s.mem.Buffers),
			Cached:      BytesToUnit(s.mem.Cached),
			UsedPercent: math.Round(s.mem.UsedPercent*10) / 10,
		},
		CPU: math.Round(cpuPercent[0]*10) / 10,
		Disk: ResourceConsumptionDisk{
			Total:       BytesToUnit(s.disk.Total),
			Used:        BytesToUnit(s.disk.Used),
			Free:        BytesToUnit(s.disk.Free),
			UsedPercent: math.Round(s.disk.UsedPercent*10) / 10,
		},
		Net: ResourceConsumptionNet{
			Sent:     BytesToUnit(netPrevBytesSent - netSent),
			Received: BytesToUnit(netPrevBytesRecv - netRecv),
		},
	}
}
