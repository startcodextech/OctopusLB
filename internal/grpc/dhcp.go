package grpc

import (
	"github.com/startcodextech/octopuslb/internal/dhcp"
	api "github.com/startcodextech/octopuslb/tools/proto"
)

type gRPCServerDHCP struct {
	api.UnimplementedDHCPServiceServer
	manager *dhcp.DHCPManager
}

func (s *gRPCServerDHCP) GetNetworksInterfaces() (*api.ResponseGetNetworkInterfaces, error) {
	response := &api.ResponseGetNetworkInterfaces{
		Success: true,
	}
	networksInterfaces, err := s.manager.GetNetworksInterfaces()
	if err != nil {
		response.Success = false
		response.Error = err.Error()
		return response, nil
	}

	response.Data = make([]*api.NetworkInterface, 0)
	for _, networkInterface := range networksInterfaces {
		response.Data = append(response.Data, &api.NetworkInterface{
			Name:       networkInterface.Name,
			Alias:      networkInterface.Alias,
			MacAddress: networkInterface.MacAddress,
			Ip:         networkInterface.IP,
			Mask:       networkInterface.Mask,
			Gateway:    networkInterface.Gateway,
			Dns:        networkInterface.DNS,
			Up:         networkInterface.Up,
		})
	}

	return response, nil
}

func (s *gRPCServerDHCP) ConfigureDHCP() (*api.Response, error) {
	response := &api.Response{
		Success: true,
	}
	err := s.manager.ConfigureDHCP()

}
