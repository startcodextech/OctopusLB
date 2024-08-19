package dhcp

import (
	"github.com/insomniacslk/dhcp/dhcpv4"
	"net"
	"sync"
	"time"
)

var (
	// DefaultLeaseDuration is the default duration 1 week of the IP lease.
	DefaultLeaseDuration = 7 * 24 * time.Hour

	// DefaultRangeStartIp is the default start IP 192.168.200.2 address for the DHCP server.
	DefaultRangeStartIp = net.ParseIP("192.168.200.2")

	// DefaultRangeEndIp is the default end IP 192.168.200.254 address for the DHCP server
	DefaultRangeEndIp = net.ParseIP("192.168.200.254")

	// DefaultSubnetMask is the default 255.255.255.0 subnet mask
	DefaultSubnetMask = net.CIDRMask(24, 32)

	// DefaultRouter is the default IP 192.168.200.1 address for the router/gateway
	DefaultRouter = net.ParseIP("192.168.200.1")

	// DNSGoogle1 and DNSGoogle2 are the default Google DNS servers
	DNSGoogle1 = net.ParseIP("8.8.8.8")
	DNSGoogle2 = net.ParseIP("8.8.4.4")

	// DefaultDNS is the default list of DNS servers
	DefaultDNS = []net.IP{DefaultRouter, DNSGoogle1, DNSGoogle2}
)

type (
	// Lease represents an IP lease.
	Lease struct {
		IP     net.IP    `json:"ip"`
		MAC    string    `json:"mac_address"`
		Expire time.Time `json:"expire"`
	}

	// Config is the configuration optional for the DHCPHandler.
	Config struct {
		// LeaseDuration is the duration of the IP lease.
		LeaseDuration time.Duration `json:"lease_duration"`

		// IPRange is the range of IP addresses that the DHCP server can assign.
		IPRange [2]net.IP `json:"ip_range"`

		// SubnetMask is the subnet mask of the network.
		SubnetMask net.IPMask `json:"subnet_mask"`

		// Router is the IP address of the router/gateway.
		Router net.IP `json:"router"`

		// DNS is the list of DNS servers.
		DNS []net.IP `json:"dns"`
	}

	// State is the combined structure that contains the Configuration and Leases
	State struct {
		Config *Config           `json:"config"`
		Leases map[string]*Lease `json:"leases"`
	}

	// DHCPHandler handles DHCP requests.
	DHCPHandler struct {
		state      *State
		leaseMutex sync.Mutex
	}
)

// NewDHCPHandler creates a new DHCPHandler with the given configuration.
// If the configuration is nil, the default configuration will be used.
//
// The default configuration is:
//   - LeaseDuration: 7 days
//   - IPRange: 192.168.200.2 - 192.168.200.254
//   - SubnetMask: 255.255.255.0
//   - Router: 192.168.200.1
//   - DNS: 192.168.200.1 - 8.8.8.8 - 8.8.4.4
func NewDHCPHandler(state *State) *DHCPHandler {
	handler := &DHCPHandler{}
	handler.state = handler.initializeState(state)
	return handler
}

// serveDHCP serves a DHCP request and returns a DHCP response.
func (h *DHCPHandler) serveDHCP(request *dhcpv4.DHCPv4) (*dhcpv4.DHCPv4, error) {
	switch request.MessageType() {
	case dhcpv4.MessageTypeDiscover:
		return h.handleDiscover(request)
	case dhcpv4.MessageTypeRequest:
		return h.handleRequest(request)
	default:
		return nil, nil
	}
}

func (h *DHCPHandler) Handle(conn net.PacketConn, peer net.Addr, m *dhcpv4.DHCPv4) {
	response, err := h.serveDHCP(m)
	if err != nil {
		return
	}

	if response != nil {
		_, err = conn.WriteTo(response.ToBytes(), peer)
		if err != nil {

		}
	}
}

// handleDiscover handles a DHCP Discover request.
func (h *DHCPHandler) handleDiscover(request *dhcpv4.DHCPv4) (*dhcpv4.DHCPv4, error) {
	h.leaseMutex.Lock()
	defer h.leaseMutex.Unlock()

	macAddress := request.ClientHWAddr.String()

	offerIP := h.getAvailableIP(macAddress)
	if offerIP == nil {
		return nil, nil
	}

	expire := time.Now().Add(h.state.Config.LeaseDuration)

	response, err := dhcpv4.NewReplyFromRequest(request,
		dhcpv4.WithMessageType(dhcpv4.MessageTypeOffer),
		dhcpv4.WithYourIP(offerIP),
		dhcpv4.WithServerIP(h.state.Config.Router),
		dhcpv4.WithLeaseTime(uint32(h.state.Config.LeaseDuration/time.Second)),
		dhcpv4.WithNetmask(h.state.Config.SubnetMask),
		dhcpv4.WithRouter(h.state.Config.Router),
		dhcpv4.WithDNS(h.state.Config.DNS...),
	)
	if err != nil {
		return nil, err
	}

	h.state.Leases[request.ClientHWAddr.String()] = &Lease{
		IP:     offerIP,
		MAC:    macAddress,
		Expire: expire,
	}

	return response, nil
}

// handleRequest handles a DHCP Request request.
func (h *DHCPHandler) handleRequest(request *dhcpv4.DHCPv4) (*dhcpv4.DHCPv4, error) {
	h.leaseMutex.Lock()
	defer h.leaseMutex.Unlock()

	macAddress := request.ClientHWAddr.String()

	requestedIP := request.ClientIPAddr
	lease, exists := h.state.Leases[macAddress]
	if exists && lease.IP.Equal(requestedIP) && lease.Expire.After(time.Now()) {
		response, err := dhcpv4.NewReplyFromRequest(request,
			dhcpv4.WithMessageType(dhcpv4.MessageTypeAck),
			dhcpv4.WithYourIP(requestedIP),
			dhcpv4.WithServerIP(h.state.Config.Router),
			dhcpv4.WithLeaseTime(uint32(h.state.Config.LeaseDuration/time.Second)),
			dhcpv4.WithNetmask(h.state.Config.SubnetMask),
			dhcpv4.WithRouter(h.state.Config.Router),
			dhcpv4.WithDNS(h.state.Config.DNS...),
		)
		if err != nil {
			return nil, err
		}
		return response, nil
	}

	response, err := dhcpv4.NewReplyFromRequest(request,
		dhcpv4.WithMessageType(dhcpv4.MessageTypeNak),
	)
	if err != nil {
		return nil, err
	}
	return response, nil
}

// getAvailableIP returns an available IP address for the given MAC address.
func (h *DHCPHandler) getAvailableIP(macAddress string) net.IP {
	for ip := h.state.Config.IPRange[0]; !ip.Equal(h.state.Config.IPRange[1]); h.incrementIP(ip) {
		if h.isBroadcastAddress(ip) {
			continue
		}

		lease, taken := h.state.Leases[macAddress]
		if !taken || lease.Expire.Before(time.Now()) {
			if taken && lease.Expire.Before(time.Now()) {
				delete(h.state.Leases, macAddress)
			}
			return ip
		}
	}
	return nil
}

// isBroadcastAddress returns true if the IP address is a broadcast address.
func (h *DHCPHandler) isBroadcastAddress(ip net.IP) bool {
	broadcast := net.IP(make([]byte, len(ip)))
	for i := 0; i < len(ip); i++ {
		broadcast[i] = ip[i] | ^h.state.Config.SubnetMask[i]
	}
	return ip.Equal(broadcast)
}

// incrementIP increments the IP address by one.
func (h *DHCPHandler) incrementIP(ip net.IP) {
	for i := len(ip) - 1; i >= 0; i-- {
		ip[i]++
		if ip[i] > 0 {
			break
		}
	}
}

// CleanupExpiredLeases removes expired leases.
func (h *DHCPHandler) CleanupExpiredLeases() {
	h.leaseMutex.Lock()
	defer h.leaseMutex.Unlock()

	now := time.Now()
	for mac, lease := range h.state.Leases {
		if now.After(lease.Expire) {
			delete(h.state.Leases, mac)
		}
	}
}

// StartLeaseCleanup starts a goroutine that cleans up expired leases at the given interval.
func (h *DHCPHandler) StartLeaseCleanup(interval time.Duration) {
	ticker := time.NewTicker(interval)
	go func() {
		for range ticker.C {
			h.CleanupExpiredLeases()
		}
	}()
}

func (h *DHCPHandler) Reload(state *State) {
	h.state = h.initializeState(state)
}

func (h *DHCPHandler) initializeState(state *State) *State {
	if state == nil {
		state = &State{}
	}
	if state.Config == nil {
		state.Config = &Config{}
	}
	if state.Leases == nil {
		state.Leases = make(map[string]*Lease)
	}
	if state.Config.LeaseDuration == 0 {
		state.Config.LeaseDuration = DefaultLeaseDuration
	}
	if state.Config.IPRange[0] == nil {
		state.Config.IPRange[0] = DefaultRangeStartIp
	}
	if state.Config.IPRange[1] == nil {
		state.Config.IPRange[1] = DefaultRangeEndIp
	}
	if state.Config.SubnetMask == nil {
		state.Config.SubnetMask = DefaultSubnetMask
	}
	if state.Config.Router == nil {
		state.Config.Router = DefaultRouter
	}
	if len(state.Config.DNS) == 0 {
		state.Config.DNS = DefaultDNS
	}
	if state.Leases == nil {
		state.Leases = make(map[string]*Lease)
	}
	return state
}
