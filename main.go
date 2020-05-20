package main

import (
	"fmt"
	pb "github.com/aau-network-security/haaukins-store/proto"
	rpc "github.com/aau-network-security/haaukins-store/util"
	_ "github.com/lib/pq"
	"log"
	"net"
)

const (
	port = ":50051"
)

func main() {

	s, err := rpc.InitilizegRPCServer()
	if err != nil {
		return
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
		return
	}

	opts, err := s.GrpcOpts()
	if err != nil {
		log.Fatalf("failed to retrieve server options %s", err.Error())
		return
	}

	gRPCServer := s.GetGRPCServer(opts...)
	pb.RegisterStoreServer(gRPCServer, s)
	fmt.Println("waiting client")
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
