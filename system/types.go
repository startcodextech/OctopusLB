package system

type (
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
)
