package dhcp

type (
	InterfaceConfigNetwork struct {
	}

	Interface struct {
		Name       string   `json:"name"`
		Alias      string   `json:"alias,omitempty"`
		MacAddress string   `json:"mac_address"`
		IP         string   `ip:"ip"`
		Mask       string   `json:"mask"`
		Gateway    string   `json:"gateway"`
		DNS        []string `json:"dns"`
		Up         bool     `json:"up"`
	}
)
