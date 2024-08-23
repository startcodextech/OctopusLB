package grpc

import (
	"context"
	"github.com/startcodextech/octopuslb/internal/managers"
	api "github.com/startcodextech/octopuslb/tools/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

type GRPCServerDHCP struct {
	api.UnimplementedDHCPServiceServer
	manager *managers.DHCPManager
}

func New(manager *managers.DHCPManager) *GRPCServerDHCP {
	return &GRPCServerDHCP{
		manager: manager,
	}
}

func (s *GRPCServerDHCP) GetNetworksInterfaces(ctx context.Context, empty *emptypb.Empty) (*api.ResponseGetNetworkInterfaces, error) {
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

func (s *GRPCServerDHCP) ConfigureDHCP(ctx context.Context, request *api.RequestConfigureDHCP) (*api.Response, error) {
	response := &api.Response{
		Success: true,
	}

	config := managers.DHCPConfig{
		InterfaceName: request.InterfaceName,
		IP:            request.Ip,
	}
	err := s.manager.ConfigureDHCP(config)
	if err != nil {
		response.Success = false
		response.Error = err.Error()
		return response, nil
	}

	return response, nil
}

func (s *GRPCServerDHCP) Start(ctx context.Context, empty *emptypb.Empty) (*api.Response, error) {
	response := &api.Response{
		Success: true,
	}
	err := s.manager.Start()
	if err != nil {
		return response, status.Error(codes.FailedPrecondition, err.Error())
	}
	return response, nil
}

func (s *GRPCServerDHCP) Stop(ctx context.Context, empty *emptypb.Empty) (*api.Response, error) {
	response := &api.Response{
		Success: true,
	}
	err := s.manager.Stop()
	if err != nil {
		response.Success = false
		response.Error = err.Error()
		return response, nil
	}
	return response, nil
}
