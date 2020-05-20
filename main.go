package main

import (
	"flag"
	"fmt"
	pb "github.com/aau-network-security/haaukins-store/proto"
	rpc "github.com/aau-network-security/haaukins-store/util"
	_ "github.com/lib/pq"
	"log"
	"net"
)

const (
	defaultConfigFile = "config.yml"
	port = ":50051"
)

func main() {

	confFilePtr := flag.String("config", defaultConfigFile, "configuration file")
	flag.Parse()

	c, err := rpc.NewConfigFromFile(*confFilePtr)
	if err != nil {
		log.Fatalf("unable to read configuration file \"%s\": %s\n", *confFilePtr, err)
	}

	s, err := rpc.InitilizegRPCServer(c)
	if err != nil {
		return
	}
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	opts, err := s.GrpcOpts(c)
	if err != nil {
		log.Fatalf("failed to retrieve server options %s", err.Error())
	}

	gRPCServer := s.GetGRPCServer(opts...)
	pb.RegisterStoreServer(gRPCServer, s)
	fmt.Println("waiting client")
	if err := gRPCServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
