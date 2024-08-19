package dhcp

import (
	"errors"
	"github.com/insomniacslk/dhcp/dhcpv4/server4"
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
		return err
	}

	s.server = server

	go func() {
		if err = server.Serve(); err != nil {
			panic(err)
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
				return s.Stop()
			case SIGNAL_CHANGE_NETWORK_INTERFACE:
				newIface := sig.InterfaceName
				if err := s.ChangeInterface(newIface); err != nil {
					return err
				}
			case SIGNAL_RELOAD:
				s.Reload(sig.State)
			}
		case <-s.stopChan:
			return s.server.Close()
		}
	}
}

func (s *DHCPServer) Stop() error {
	close(s.stopChan)
	return s.conn.Close()
}

func (s *DHCPServer) Reload(state *State) {
	s.handler.Reload(state)
}

func (s *DHCPServer) ChangeInterface(newIface string) error {
	err := s.Stop()
	if err != nil {
		return err
	}

	conn, err := net.ListenPacket("udp4", ":67")
	if err != nil {
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
