package main

import (
	"fmt"
	rpc "github.com/aau-network-security/haaukins-store/grpc"
	pb "github.com/aau-network-security/haaukins-store/proto"
	_ "github.com/lib/pq"
	"log"
	"net"
	"os"
)

const (
	port = ":50051"
)

func main() {
	certificate := os.Getenv("CERT")
	certificateKey := os.Getenv("CERT_KEY")
	s := rpc.InitilizegRPCServer(certificate,certificateKey)
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts, err := s.GrpcOpts()
	if err != nil {
		log.Fatal("failed to retrieve server options")
	}

	gRPCServer := s.GetGRPCServer(opts...)
	pb.RegisterStoreServer(gRPCServer, s)
	fmt.Println("waiting client")
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
