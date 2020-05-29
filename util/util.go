package util

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/aau-network-security/haaukins-store/database"
	"github.com/aau-network-security/haaukins-store/model"
	pb "github.com/aau-network-security/haaukins-store/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"gopkg.in/yaml.v2"
)

type server struct {
	store database.Store
	auth  Authenticator
	tls   bool
}

type certificate struct {
	cPath    string
	cKeyPath string
	caPath   string
}

func (s server) AddEvent(ctx context.Context, in *pb.AddEventRequest) (*pb.InsertResponse, error) {
	result, err := s.store.AddEvent(in)
	if err != nil {
		log.Printf("ERR: Error Add Event %s", err.Error())
		return &pb.InsertResponse{ErrorMessage: err.Error()}, nil
	}
	log.Printf("Event %s Saved", in.Tag)
	return &pb.InsertResponse{Message: result}, nil

}

func (s server) AddTeam(ctx context.Context, in *pb.AddTeamRequest) (*pb.InsertResponse, error) {
	result, err := s.store.AddTeam(in)
	if err != nil {
		log.Printf("ERR: Error Add Team %s", err.Error())
		return &pb.InsertResponse{ErrorMessage: err.Error()}, nil
	}
	log.Printf("Team %s Saved for the Event %s", in.Id, in.EventTag)
	return &pb.InsertResponse{Message: result}, nil
}

func (s server) GetEvents(context.Context, *pb.EmptyRequest) (*pb.GetEventResponse, error) {
	result, err := s.store.GetEvents()
	if err != nil {
		log.Printf("ERR: Error Get Events %s", err.Error())
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
			Status:             e.Status,
		})
	}
	log.Printf("Get Events")
	return &pb.GetEventResponse{Events: events}, nil
}

func (s server) GetEventStatus(ctx context.Context, in *pb.GetEventStatusRequest) (*pb.EventStatus, error) {
	result, err := s.store.GetEventStatus(in)
	if err != nil {
		return &pb.EventStatus{Status: "Error on getting status of event " + err.Error()}, err
	}
	log.Printf("Event status returned ! [Status: %s , Event: %s] ", result, in.EventTag)
	return &pb.EventStatus{Status: result}, nil

}

func (s server) SetEventStatus(ctx context.Context, in *pb.SetEventStatusRequest) (*pb.EventStatus, error) {
	log.Printf("Set event status for event %s to %s", in.EventTag, in.Status)
	result, err := s.store.SetEventStatus(in)
	if err != nil {
		return &pb.EventStatus{Status: err.Error()}, err
	}

	log.Printf("Event status updated ! [Status: %s , Event: %s] ", result, in.EventTag)

	return &pb.EventStatus{Status: result}, nil
}

func (s server) GetEventTeams(ctx context.Context, in *pb.GetEventTeamsRequest) (*pb.GetEventTeamsResponse, error) {
	result, err := s.store.GetTeams(in.EventTag)
	if err != nil {
		log.Printf("ERR: Error Get teams for Event %s : %s", in.EventTag, err.Error())
		return &pb.GetEventTeamsResponse{ErrorMessage: err.Error()}, nil
	}

	var teams []*pb.GetEventTeamsResponse_Teams
	for _, t := range result {
		teams = append(teams, &pb.GetEventTeamsResponse_Teams{
			Id:               t.Tag,
			Email:            t.Email,
			Name:             t.Name,
			HashPassword:     t.Password,
			CreatedAt:        t.CreatedAt,
			LastAccess:       t.LastAccess,
			SolvedChallenges: t.SolvedChallenges,
		})
	}
	log.Printf("Get Teams for the Event %s", in.EventTag)
	return &pb.GetEventTeamsResponse{Teams: teams}, nil
}

func (s server) UpdateEventFinishDate(ctx context.Context, in *pb.UpdateEventRequest) (*pb.UpdateResponse, error) {
	result, err := s.store.UpdateEventFinishDate(in)
	if err != nil {
		log.Printf("ERR: Error Update Event %s finish time: %s", in.EventId, err.Error())
		return &pb.UpdateResponse{ErrorMessage: err.Error()}, nil
	}
	log.Printf("Event %s Stopped", in.EventId)
	return &pb.UpdateResponse{Message: result}, nil
}

func (s server) UpdateTeamSolvedChallenge(ctx context.Context, in *pb.UpdateTeamSolvedChallengeRequest) (*pb.UpdateResponse, error) {
	result, err := s.store.UpdateTeamSolvedChallenge(in)
	if err != nil {
		log.Printf("ERR: Error Update team %s solve challenge: %s", in.TeamId, err.Error())
		return &pb.UpdateResponse{ErrorMessage: err.Error()}, nil
	}
	log.Printf("Team %s solved %s challenge", in.TeamId, in.Tag)
	return &pb.UpdateResponse{Message: result}, nil
}

func (s server) UpdateTeamLastAccess(ctx context.Context, in *pb.UpdateTeamLastAccessRequest) (*pb.UpdateResponse, error) {
	result, err := s.store.UpdateTeamLastAccess(in)
	if err != nil {
		log.Printf("ERR: Error Update team %s last access: %s", in.TeamId, err.Error())
		return &pb.UpdateResponse{ErrorMessage: err.Error()}, nil
	}
	return &pb.UpdateResponse{Message: result}, nil
}

func GetCreds(conf *model.Config) (credentials.TransportCredentials, error) {
	log.Printf("Preparing credentials for RPC")

	certificateProps := certificate{
		cPath:    conf.TLS.CertFile,
		cKeyPath: conf.TLS.CertKey,
		caPath:   conf.TLS.CAFile,
	}

	certificate, err := tls.LoadX509KeyPair(certificateProps.cPath, certificateProps.cKeyPath)
	if err != nil {
		return nil, fmt.Errorf("could not load server key pair: %s", err)
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

func (s server) GrpcOpts(conf *model.Config) ([]grpc.ServerOption, error) {

	if conf.TLS.Enabled {
		creds, err := GetCreds(conf)

		if err != nil {
			return []grpc.ServerOption{}, errors.New("Error on retrieving certificates: " + err.Error())
		}
		log.Printf("Server is running in secure mode !")
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

func InitilizegRPCServer(conf *model.Config) (*server, error) {

	store, err := database.NewStore(conf)

	if err != nil {
		return nil, err
	}

	s := &server{
		store: store,
		auth:  NewAuthenticator(conf.SigninKey, conf.AuthKey),
		tls:   conf.TLS.Enabled,
	}
	return s, nil
}

func NewConfigFromFile(path string) (*model.Config, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var c model.Config
	err = yaml.Unmarshal(f, &c)
	if err != nil {
		return nil, err
	}

	if c.Host == "" {
		log.Println("Host not provided in the configuration file")
		c.Host = "localhost:50051"
	}

	if c.SigninKey == "" {
		log.Println("SigninKey not provided in the configuration file")
		c.Host = "dev-env"
	}

	if c.AuthKey == "" {
		log.Println("AuthKey not provided in the configuration file")
		c.Host = "development-environment"
	}

	if c.DB.Host == "" || c.DB.User == "" || c.DB.Pass == "" || c.DB.Name == "" {
		return nil, errors.New("DB paramenters missing in the configuration file")
	}

	if c.DB.Port == 0 {
		c.DB.Port = 5432
	}

	if c.TLS.Enabled {
		if c.TLS.CAFile == "" || c.TLS.CertKey == "" || c.TLS.CertFile == "" {
			return nil, errors.New("Provide Certificates in the config file")
		}
	}

	return &c, nil
}
