package main

import (
	grpc2 "github.com/startcodextech/octopuslb/internal/grpc"
	"github.com/startcodextech/octopuslb/internal/logs"
	"github.com/startcodextech/octopuslb/internal/managers"
	api "github.com/startcodextech/octopuslb/tools/proto"
	"google.golang.org/grpc"
	"net"
)

func init() {
	logs.Init()
}

func main() {
	lis, err := net.Listen("tcp", ":3000")
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()

	dhcpManager, err := managers.NewDHCPManager()
	if err != nil {
		panic(err)
	}

	grpcDHCP := grpc2.New(dhcpManager)
	api.RegisterDHCPServiceServer(server, grpcDHCP)

	if err := server.Serve(lis); err != nil {
		panic(err)
	}
}
