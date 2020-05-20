package tests

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"fmt"
	pb "github.com/aau-network-security/haaukins-store/proto"
	"github.com/dgrijalva/jwt-go"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/status"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const (
	AUTH_KEY = "au"
)

var (
	testCertPath    = strings.Replace(os.Getenv("CERT"), "./tests/","./",1)
	testCertKeyPath = strings.Replace(os.Getenv("CERT_KEY"), "./tests/","./",1)
	testCAPath 		= strings.Replace(os.Getenv("CA"), "./tests/","./",1)
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

func TestStoreConnection(t *testing.T) {
	addr := os.Getenv("HOST")

	tokenCorret := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		AUTH_KEY: os.Getenv("AUTH_KEY"),
	})

	tokenError := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		AUTH_KEY: "wrong-token",
	})

	tt := []struct {
		name  string
		token *jwt.Token
		err   string
	}{
		{name: "Normal Authentication", token: tokenCorret},
		{name: "Unauthorized", token: tokenError, err: "Invalid Authentication Key"},
	}

	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			signin_key := os.Getenv("SIGNIN_KEY")

			tokenString, err := tc.token.SignedString([]byte(signin_key))
			if err != nil {
				t.Fatalf("Error creating the token")
			}

			authCreds := Creds{Token: tokenString}

			// Load the client certificates from disk
			certificate, err := tls.LoadX509KeyPair(testCertPath, testCertKeyPath)
			if err != nil {
				t.Fatalf("could not load client key pair: %s", err)
			}

			// Create a certificate pool from the certificate authority
			certPool := x509.NewCertPool()
			ca, err := ioutil.ReadFile(testCAPath)
			if err != nil {
				t.Fatalf("could not read ca certificate: %s", err)
			}

			// Append the certificates from the CA
			if ok := certPool.AppendCertsFromPEM(ca); !ok {
				 t.Fatalf("failed to append ca certs")
			}

			creds := credentials.NewTLS(&tls.Config{
				ServerName:   addr,
				Certificates: []tls.Certificate{certificate},
				RootCAs:      certPool,
			})

			dialOpts := []grpc.DialOption{
				grpc.WithTransportCredentials(creds),
				grpc.WithPerRPCCredentials(authCreds),
			}
			// Create a connection with the TLS credentials

			conn, err := grpc.Dial(addr, dialOpts...)
			if err != nil {
				t.Fatalf("Connection error: %v", err)
			}
			defer conn.Close()

			c := pb.NewStoreClient(conn)

			_, err = c.GetEvents(context.Background(), &pb.EmptyRequest{})

			if err != nil {
				st, ok := status.FromError(err)
				if ok {
					err = fmt.Errorf(st.Message())
				}

				if tc.err != "" {
					if tc.err != err.Error() {
						t.Fatalf("unexpected error (expected: %s) received: %s", tc.err, err.Error())
					}
					return
				}
				t.Fatalf("expected no error, but received: %s", err)
			}

			if tc.err != "" {
				t.Fatalf("expected error, but received none")
			}
		})
	}
}

func createTestClientConn() (*grpc.ClientConn, error){
	addr := os.Getenv("HOST")

	tokenCorret := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		AUTH_KEY: os.Getenv("AUTH_KEY"),
	})

	signin_key := os.Getenv("SIGNIN_KEY")

	tokenString, err := tokenCorret.SignedString([]byte(signin_key))
	if err != nil {
		return nil, err
	}

	authCreds := Creds{Token: tokenString}

	// Load the client certificates from disk
	certificate, err := tls.LoadX509KeyPair(testCertPath, testCertKeyPath)
	if err != nil {
		return nil, err
	}

	// Create a certificate pool from the certificate authority
	certPool := x509.NewCertPool()
	ca, err := ioutil.ReadFile(testCAPath)
	if err != nil {
		return nil, err
	}

	// Append the certificates from the CA
	if ok := certPool.AppendCertsFromPEM(ca); !ok {
		return nil, err
	}

	creds := credentials.NewTLS(&tls.Config{
		ServerName:   addr,
		Certificates: []tls.Certificate{certificate},
		RootCAs:      certPool,
	})

	dialOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(creds),
		grpc.WithPerRPCCredentials(authCreds),
	}

	// Create a connection with the TLS credentials
	conn, err := grpc.Dial(addr, dialOpts...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

func TestAddEvent(t *testing.T){

	conn, err := createTestClientConn()

	defer conn.Close()

	c := pb.NewStoreClient(conn)

	if err != nil {
		t.Fatal(err)
	}

	req := pb.AddEventRequest{
		Name: 				"Test",
		Tag: 				"test",
		Frontends:			"kali",
		Exercises: 			"ftp,xss",
		Available: 			1,
		Capacity: 			2,
		StartTime:  		"2020-05-20 14:35:01",
		ExpectedFinishTime: "2020-05-21 14:35:01",

	}

	_, err = c.AddEvent(context.Background(), &req)
	if err != nil {
		t.Fatal(err)
	}
	events, err := c.GetEvents(context.Background(), &pb.EmptyRequest{})
	if err != nil {
		t.Fatal(err)
	}
	if len(events.Events) != 1 {
		t.Fatal("Error getting the stored events")
	}
}
