package main

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
	jwt "github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc/credentials"
	"io/ioutil"
	"log"
	"time"

	pb "github.com/aau-network-security/haaukins-store/proto"
	"google.golang.org/grpc"
)

const (
	address  = "localhost:50051"
	AUTH_KEY = "au"
)

var (
	UnreachableDaemonErr = errors.New("Daemon seems to be unreachable")
	UnauthorizedErr      = errors.New("You seem to not be logged in")
)

type Creds struct {
	Token    string
	Insecure bool
}

func (c Creds) GetRequestMetadata(context.Context, ...string) (map[string]string, error) {
	return map[string]string{
		"token": string(c.Token),
	}, nil
}

func (c Creds) RequireTransportSecurity() bool {
	return !c.Insecure
}

func main() {

	//todo just for test purpose
	test_auth_key := "c41ec030-db76-473f-a504-5a7323aa04ec"
	test_sign_key := "34b16c10-1a2c-4533-83e8-cfde78817501"
	testCertPath := "/home/ubuntu/haaukins_main/configs/certs/localhost.crt"
	testCertKeyPath:= "/home/ubuntu/haaukins_main/configs/certs/localhost.key"
	testCAPath := "/home/ubuntu/haaukins_main/configs/certs/haaukins-store.com.crt"

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		AUTH_KEY: test_auth_key,
	})

	tokenString, err := token.SignedString([]byte(test_sign_key))
	if err != nil {
		fmt.Println("Error creating the token")
	}

	authCreds := Creds{Token: tokenString}
	dialOpts := []grpc.DialOption{}

	ssl := true
	if ssl {
		// Load the client certificates from disk
		certificate, err := tls.LoadX509KeyPair(testCertPath, testCertKeyPath)
		if err != nil {
			log.Printf("could not load client key pair: %s", err)
		}

		// Create a certificate pool from the certificate authority
		certPool := x509.NewCertPool()
		ca, err := ioutil.ReadFile(testCAPath)
		if err != nil {
			log.Printf("could not read ca certificate: %s", err)
		}

		// Append the certificates from the CA
		if ok := certPool.AppendCertsFromPEM(ca); !ok {
			log.Println("failed to append ca certs")
		}

		creds := credentials.NewTLS(&tls.Config{
			ServerName:   address,
			Certificates: []tls.Certificate{certificate},
			RootCAs:      certPool,
		})

		dialOpts = append(dialOpts, grpc.WithTransportCredentials(creds),grpc.WithPerRPCCredentials(authCreds))

	} else {
		authCreds.Insecure = true
		dialOpts = append(dialOpts,
			grpc.WithInsecure(),
			grpc.WithPerRPCCredentials(authCreds))
	}

	conn, err := grpc.Dial(address, dialOpts...)
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewStoreClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	r, err := c.GetEvents(ctx, &pb.EmptyRequest{})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	if r.ErrorMessage != "" {
		log.Fatalf("my could not greet: %v", r.ErrorMessage)
	}
	//log.Println(r.Message)
	for _, e := range r.Events {
		fmt.Println(e)
	}
}
