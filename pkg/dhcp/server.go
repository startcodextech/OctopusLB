package dhcp

import (
	"errors"
	"github.com/insomniacslk/dhcp/dhcpv4/server4"
	"github.com/phuslu/log"
	"net"
	"sync"
	"time"
)

const (
	SIGNAL_STOP                     = "stop"
	SIGNAL_CHANGE_NETWORK_INTERFACE = "change_network_interface"
	SIGNAL_RELOAD                   = "reload"
)

var (
	ErrEmptyInterfaceName = errors.New("interface name is empty")
)

type (
	DHCPServer struct {
		handler       *DHCPHandler
		interfaceName string
		conn          net.PacketConn
		server        *server4.Server
		signalChan    chan Signal
		wg            sync.WaitGroup
		stopChan      chan struct{}
	}

	Signal struct {
		Name          string
		InterfaceName string
		State         *State
	}
)

func NewServer(netInterface string, state *State) (*DHCPServer, error) {
	handler := NewDHCPHandler(state)

	conn, err := net.ListenPacket("udp4", ":67")
	if err != nil {
		log.Error().
			Str("service", "dhcp").
			Err(err).
			Msg("Failed to listen on interface")
		return nil, err
	}

	return &DHCPServer{
		handler:       handler,
		interfaceName: netInterface,
		conn:          conn,
		signalChan:    make(chan Signal),
		stopChan:      make(chan struct{}),
	}, nil
}

func (s *DHCPServer) Start() error {
	if s.interfaceName == "" {
		log.Error().
			Str("service", "dhcp").
			Msg("Interface name is empty")
		return ErrEmptyInterfaceName
	}

	s.wg.Add(1)
	defer s.wg.Done()

	addr := &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 67,
	}

	server, err := server4.NewServer(s.interfaceName, addr, s.handler.Handle)
	if err != nil {
		log.Error().
			Str("service", "dhcp").
			Err(err).
			Msg("Failed to create DHCP server")
		return err
	}

	s.server = server

	log.Info().
		Str("service", "dhcp").
		Msg("Starting DHCP server on interface " + s.interfaceName)

	go func() {
		if err := server.Serve(); err != nil {
			log.Error().
				Str("service", "dhcp").
				Err(err).
				Msg("Failed to start DHCP server")
		}
	}()

	s.handler.StartLeaseCleanup(1 * time.Minute)

	return s.handleSignal()
}

func (s *DHCPServer) handleSignal() error {
	for {
		select {
		case sig := <-s.signalChan:
			switch sig.Name {
			case SIGNAL_STOP:
				log.Info().
					Str("service", "dhcp").
					Msg("Shutting down DHCP server...")
				return s.Stop()
			case SIGNAL_CHANGE_NETWORK_INTERFACE:
				newIface := sig.InterfaceName
				log.Info().
					Str("service", "dhcp").
					Msg("Changing interface to " + newIface)
				if err := s.ChangeInterface(newIface); err != nil {
					log.Error().
						Str("service", "dhcp").
						Err(err).
						Msg("Failed to change interface")
				}
			case SIGNAL_RELOAD:
				log.Info().
					Str("service", "dhcp").
					Msg("Reloading DHCP server configuration...")
				s.Reload(sig.State)
			}
		case <-s.stopChan:
			log.Info().
				Str("service", "dhcp").
				Msg("Shutting down DHCP server...")
			return nil
		}
	}
}

func (s *DHCPServer) Stop() error {
	log.Info().
		Str("service", "dhcp").
		Msgf("Stopping DHCP server on interface %s...", s.interfaceName)
	close(s.stopChan)
	return s.conn.Close()
}

func (s *DHCPServer) Reload(state *State) {
	log.Info().
		Str("service", "dhcp").
		Msg("Reloading DHCP server configuration...")

	s.handler.Reload(state)

	log.Info().
		Str("service", "dhcp").
		Msg("DHCP server configuration reloaded successfully.")
}

func (s *DHCPServer) ChangeInterface(newIface string) error {
	log.Info().
		Str("service", "dhcp").
		Msgf("Changing interface from %s to %s...", s.interfaceName, newIface)
	err := s.Stop()
	if err != nil {
		log.Error().
			Str("service", "dhcp").
			Err(err).
			Msg("Failed to stop DHCP server")
		return err
	}

	conn, err := net.ListenPacket("udp4", ":67")
	if err != nil {
		log.Error().
			Str("service", "dhcp").
			Err(err).
			Msg("Failed to listen on new interface")
		return err
	}

	s.interfaceName = newIface
	s.conn = conn

	return s.Start()
}

func (s *DHCPServer) SendSignal(signalName string, interfaceName *string, state *State) {
	sig := Signal{Name: signalName}
	if interfaceName != nil {
		if *interfaceName != "" {
			sig.InterfaceName = *interfaceName
		}
	}
	if state != nil {
		sig.State = state
	}
	s.signalChan <- sig
}
