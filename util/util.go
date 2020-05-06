package util

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"github.com/aau-network-security/haaukins-store/database"
	pb "github.com/aau-network-security/haaukins-store/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

type server struct {
	store   database.Store
	auth    Authenticator
	tls     bool
}

type certificate struct {
	cPath 	string
	cKeyPath  string
	caPath				string
}

func (s server) AddEvent(ctx context.Context, in *pb.AddEventRequest) (*pb.InsertResponse, error) {
	result, err := s.store.AddEvent(in)
	if err != nil {
		return &pb.InsertResponse{ErrorMessage: err.Error()}, nil
	}
	return &pb.InsertResponse{Message: result}, nil

}

func (s server) AddTeam(ctx context.Context, in *pb.AddTeamRequest) (*pb.InsertResponse, error) {
	result, err := s.store.AddTeam(in)
	if err != nil {
		return &pb.InsertResponse{ErrorMessage: err.Error()}, nil
	}
	return &pb.InsertResponse{Message: result}, nil
}

func (s server) GetEvents(context.Context, *pb.EmptyRequest) (*pb.GetEventResponse, error) {
	result, err := s.store.GetEvents()
	if err != nil {
		return &pb.GetEventResponse{ErrorMessage: err.Error()}, nil
	}

	var events []*pb.GetEventResponse_Events
	for _, e := range result {
		events = append(events, &pb.GetEventResponse_Events{
			Name:               e.Name,
			Tag:                e.Tag,
			Frontends:          e.Frontends,
			Exercises:          e.Exercises,
			Available:          int32(e.Available),
			Capacity:           int32(e.Capacity),
			StartedAt:          e.StartedAt,
			ExpectedFinishTime: e.ExpectedFinishTime,
			FinishedAt:         e.FinishedAt,
		})
	}

	return &pb.GetEventResponse{Events: events}, nil

}

func (s server) GetEventTeams(ctx context.Context, in *pb.GetEventTeamsRequest) (*pb.GetEventTeamsResponse, error) {
	result, err := s.store.GetTeams(in.EventTag)
	if err != nil {
		return &pb.GetEventTeamsResponse{ErrorMessage: err.Error()}, nil
	}

	var teams []*pb.GetEventTeamsResponse_Teams
	for _, t := range result {
		teams = append(teams, &pb.GetEventTeamsResponse_Teams{
			Id:               t.Id,
			Email:            t.Email,
			Name:             t.Name,
			HashPassword:     t.Password,
			CreatedAt:        t.CreatedAt,
			LastAccess:       t.LastAccess,
			SolvedChallenges: t.SolvedChallenges,
		})
	}

	return &pb.GetEventTeamsResponse{Teams: teams}, nil
}

func (s server) UpdateEventFinishDate(ctx context.Context, in *pb.UpdateEventRequest) (*pb.UpdateResponse, error) {
	result, err := s.store.UpdateEventFinishDate(in)
	if err != nil {
		return &pb.UpdateResponse{ErrorMessage: err.Error()}, nil
	}
	return &pb.UpdateResponse{Message: result}, nil
}

func (s server) UpdateTeamSolvedChallenge(ctx context.Context, in *pb.UpdateTeamSolvedChallengeRequest) (*pb.UpdateResponse, error) {
	result, err := s.store.UpdateTeamSolvedChallenge(in)
	if err != nil {
		return &pb.UpdateResponse{ErrorMessage: err.Error()}, nil
	}
	return &pb.UpdateResponse{Message: result}, nil
}

func (s server) UpdateTeamLastAccess(ctx context.Context, in *pb.UpdateTeamLastAccessRequest) (*pb.UpdateResponse, error) {
	result, err := s.store.UpdateTeamLastAccess(in)
	if err != nil {
		return &pb.UpdateResponse{ErrorMessage: err.Error()}, nil
	}
	return &pb.UpdateResponse{Message: result}, nil
}

func GetCreds() (credentials.TransportCredentials,error) {
	log.Printf("Preparing credentials for RPC")
	// todo: change environment variables into configuration
	// add handling functionality
	certificateProps := certificate{
		cPath:    			os.Getenv("CERT"),
		cKeyPath: 			os.Getenv("CERT_KEY"),
		caPath:             os.Getenv("CA"),
	}

	certificate, err := tls.LoadX509KeyPair(certificateProps.cPath, certificateProps.cKeyPath)
	if err != nil {
		return nil,fmt.Errorf("could not load server key pair: %s", err)
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(certificateProps.caPath)
	if err != nil {
		return nil, fmt.Errorf("could not read ca certificate: %s", err)
	}
	// CA file for let's encrypt is located under domain conf as `chain.pem`
	// pass chain.pem location
	// Append the client certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, errors.New("failed to append client certs")
	}

	// Create the TLS credentials
	creds := credentials.NewTLS(&tls.Config{
		ClientAuth:   tls.RequireAndVerifyClientCert,
		Certificates: []tls.Certificate{certificate},
		ClientCAs:    certPool,
	})
	return creds, nil
}

func (s server) GrpcOpts() ([]grpc.ServerOption, error) {

	if s.tls {
		creds,err := GetCreds()

		if err != nil {
			return []grpc.ServerOption{}, errors.New("Error on retrieving certificates: "+err.Error())
		}
		return []grpc.ServerOption{grpc.Creds(creds)}, nil
	}
	return []grpc.ServerOption{}, nil
}

func (s server) GetGRPCServer(opts ...grpc.ServerOption) *grpc.Server {

	streamInterceptor := func(srv interface{}, stream grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		if err := s.auth.AuthenticateContext(stream.Context()); err != nil {
			return err
		}
		return handler(srv, stream)
	}

	unaryInterceptor := func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if err := s.auth.AuthenticateContext(ctx); err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}

	opts = append([]grpc.ServerOption{
		grpc.StreamInterceptor(streamInterceptor),
		grpc.UnaryInterceptor(unaryInterceptor),
	}, opts...)
	return grpc.NewServer(opts...)
}

func readContent(path string) error {
	cont, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
		return err
	}
	_, err = io.Copy(os.Stdout, cont)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}



func InitilizegRPCServer() *server {

	store, err := database.NewStore()

	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	// todo: change handling of tls, the example below not good enough

	tls := true
	mode := os.Getenv("SSL_OFF")
	if strings.ToLower(mode) == "true" {
		tls = false
	}

	s := &server{
		store:   store,
		auth:    NewAuthenticator(os.Getenv("SIGNIN_KEY")),
		tls:     tls,
	}
	return s
}
