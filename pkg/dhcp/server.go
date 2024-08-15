package dhcp

import (
	"github.com/insomniacslk/dhcp/dhcpv4/server4"
	"github.com/phuslu/log"
	"net"
	"sync"
	"time"
)

const (
	SIGNAL_STOP         = "stop"
	SIGNAL_CHANGE_IFACE = "change_iface"
	SIGNAL_RELOAD       = "reload"
)

type (
	DHCPServer struct {
		handler    *DHCPHandler
		iface      string
		conn       net.PacketConn
		server     *server4.Server
		signalChan chan Signal
		wg         sync.WaitGroup
		stopChan   chan struct{}
	}

	Signal struct {
		Name  string
		Iface string
	}
)

func NewServer() (*DHCPServer, error) {
	handler := NewDHCPHandler("")

	conn, err := net.ListenPacket("udp4", ":67")
	if err != nil {
		log.Error().
			Str("service", "dhcp").
			Err(err).
			Msg("Failed to listen on interface")
		return nil, err
	}

	return &DHCPServer{
		handler:    handler,
		iface:      "",
		conn:       conn,
		signalChan: make(chan Signal),
		stopChan:   make(chan struct{}),
	}, nil
}

func (s *DHCPServer) Start() error {
	s.wg.Add(1)
	defer s.wg.Done()

	addr := &net.UDPAddr{
		IP:   net.IPv4zero,
		Port: 67,
	}

	server, err := server4.NewServer(s.iface, addr, s.handler.Handle)
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
		Msg("Starting DHCP server on interface " + s.iface)

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
			case "stop":
				log.Info().
					Str("service", "dhcp").
					Msg("Shutting down DHCP server...")
				return s.Stop()
			case "change_iface":
				newIface := sig.Iface
				log.Info().
					Str("service", "dhcp").
					Msg("Changing interface to " + newIface)
				if err := s.ChangeInterface(newIface); err != nil {
					log.Error().
						Str("service", "dhcp").
						Err(err).
						Msg("Failed to change interface")
				}
			case "reload":
				log.Info().
					Str("service", "dhcp").
					Msg("Reloading DHCP server configuration...")
				if err := s.Reload(); err != nil {
					log.Error().
						Str("service", "dhcp").
						Err(err).
						Msg("Failed to reload DHCP server configuration")
				}
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
		Msgf("Stopping DHCP server on interface %s...", s.iface)
	close(s.stopChan)
	return s.conn.Close()
}

func (s *DHCPServer) Reload() error {
	log.Info().
		Str("service", "dhcp").
		Msg("Reloading DHCP server configuration...")

	// Recargar el estado desde el archivo (configuración y arrendamientos)
	if err := s.handler.LoadFromFile(); err != nil {
		log.Printf("Failed to reload state: %v", err)
		return err
	}

	log.Info().
		Str("service", "dhcp").
		Msg("DHCP server configuration reloaded successfully.")
	return nil
}

func (s *DHCPServer) ChangeInterface(newIface string) error {
	log.Info().
		Str("service", "dhcp").
		Msgf("Changing interface from %s to %s...", s.iface, newIface)
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

	s.iface = newIface
	s.conn = conn

	return s.Start()
}

func (s *DHCPServer) SendSignal(signalName string, iface ...string) {
	sig := Signal{Name: signalName}
	if len(iface) > 0 {
		sig.Iface = iface[0]
	}
	s.signalChan <- sig
}
