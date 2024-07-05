package system

import (
	"github.com/phuslu/log"
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
)

type (
	System struct {
		host *host.InfoStat
		cpu  []cpu.InfoStat
		mem  *mem.VirtualMemoryStat
		disk *disk.UsageStat
		proc []*process.Process
		net  []net.InterfaceStat
		load *load.AvgStat

		OS                 OS              `json:"os"`
		Architecture       string          `json:"architecture"`
		Virtualization     string          `json:"virtualization"`
		VirtualizationRole string          `json:"virtualization_role"`
		CPUs               int             `json:"cpus"`
		CPUModel           string          `json:"cpu_model"`
		RAM                InformationUnit `json:"ram"`
		DiskSize           InformationUnit `json:"disk_size"`
	}
)

func Init() *System {
	hostInfo, _ := host.Info()
	cpuInfo, _ := cpu.Info()
	menInfo, _ := mem.VirtualMemory()
	diskInfo, _ := disk.Usage("/")
	procInfo, _ := process.Processes()
	netInfo, _ := net.Interfaces()
	loadInfo, _ := load.Avg()

	//log.Info().Str("distro linux", hostInfo.PlatformFamily).Msg("Distro Linux")

	pack, err := detectPackageManager()
	if err != nil {
		panic(err)
	}

	packs, err := GetPackages(hostInfo.PlatformFamily)
	if err != nil {
		panic(err)
	}

	return &System{
		host: hostInfo,
		cpu:  cpuInfo,
		mem:  menInfo,
		disk: diskInfo,
		proc: procInfo,
		net:  netInfo,
		load: loadInfo,

		OS: OS{
			Name:            hostInfo.OS,
			Platform:        hostInfo.Platform,
			PlatformFamily:  hostInfo.PlatformFamily,
			PlatformVersion: hostInfo.PlatformVersion,
			PackManager:     pack,
			Packages:        packs,
		},
		Architecture:       hostInfo.KernelArch,
		Virtualization:     hostInfo.VirtualizationSystem,
		VirtualizationRole: hostInfo.VirtualizationRole,
		CPUs:               len(cpuInfo),
		CPUModel:           cpuInfo[0].ModelName,
		RAM:                BytesToUnit(menInfo.Total),
		DiskSize:           BytesToUnit(diskInfo.Total),
	}
}

func (s *System) CurrentResourcesConsumption() ResourceConsumption {
	cpuPercent, _ := cpu.Percent(0, false)
	netCounter, _ := net.IOCounters(false)

	log.Printf("\n")
	log.Printf("cpuPercent: %v\n", netCounter)

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
