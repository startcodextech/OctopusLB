package grpc

import (
	"github.com/startcodextech/octopuslb/internal/managers"
	api "github.com/startcodextech/octopuslb/tools/proto"
)

type gRPCServerDHCP struct {
	api.UnimplementedDHCPServiceServer
	manager *managers.DHCPManager
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

func (s *gRPCServerDHCP) ConfigureDHCP(request *api.RequestConfigureDHCP) (*api.Response, error) {
	response := &api.Response{
		Success: true,
	}

	config := managers.DHCPConfig{
		InterfaceName: request.InterfaceName,
		IP:            request.Ip,
	}
	err := s.manager.ConfigureDHCP(config)
	if err != nil {
		errMessage := err.Error()
		response.Success = false
		response.Error = &errMessage
		return response, nil
	}

	return response, nil
}
