package dhcp

import "net"

type (
	InterfaceConfigNetwork struct {
	}

	Interface struct {
		Name       string           `json:"name"`
		MacAddress net.HardwareAddr `json:"mac_address"`
		IP         net.IP           `ip:"ip"`
		Mask       net.IPMask       `json:"mask"`
		Gateway    net.IP           `json:"gateway"`
		DNS        []string         `json:"dns"`
		Up         bool             `json:"up"`
	}
)

func (i *Interface) Change(config InterfaceConfig) error {
	return nil
}
