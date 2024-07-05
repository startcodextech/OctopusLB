package system

import "math"

func BytesToKB(b uint64) float64 {
	return math.Round(float64(b)/1024*10) / 10
}

func BytesToMB(b uint64) float64 {
	return math.Round(float64(b)/1024/1024*10) / 10
}

func BytesToGB(b uint64) float64 {
	return math.Round(float64(b)/1024/1024/1024*10) / 10
}

func BytesToTB(b uint64) float64 {
	return math.Round(float64(b)/1024/1024/1024/1024*10) / 10
}

func BytesToUnit(b uint64) InformationUnit {
	if b < 1024 {
		return InformationUnit{
			Value: float64(b),
			Unit:  "B",
		}
	} else if b < 1024*1024 {
		return InformationUnit{
			Value: BytesToKB(b),
			Unit:  "KB",
		}
	} else if b < 1024*1024*1024 {
		return InformationUnit{
			Value: BytesToMB(b),
			Unit:  "MB",
		}
	} else if b < 1024*1024*1024*1024 {
		return InformationUnit{
			Value: BytesToGB(b),
			Unit:  "GB",
		}
	} else {
		return InformationUnit{
			Value: BytesToTB(b),
			Unit:  "TB",
		}
	}
}
