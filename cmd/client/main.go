package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"os"
)

const (
	clientCertFile   = "cert/client-cert.pem"
	clientKeyFile    = "cert/client-key.pem"
	clientCACertFile = "cert/ca-cert.pem"
)

func main() {
	serverAddress := flag.String("address", ":9000", "Server address")
	enableTLS := flag.Bool("tls", false, "Enable SSL/TLS")
	flag.Parse()
	log.Printf("dial server %s, TLS %t", *serverAddress, *enableTLS)

	var transportOption grpc.DialOption
	transportOption = grpc.WithTransportCredentials(insecure.NewCredentials())
	//if *enableTLS {
	//	tlsCredentials, err := loadTLSCredentials()
	//	if err != nil {
	//		log.Fatalf("cannot load TLS credentials: %v", err)
	//	}
	//	transportOption = grpc.WithTransportCredentials(tlsCredentials)
	//}
	//
	cc1, err := grpc.Dial(*serverAddress, transportOption)
	if err != nil {
		log.Fatal("cannot dial server: ", err)
	}
	defer cc1.Close()

}

func loadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := os.ReadFile(clientCACertFile)
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Load client's certificate and private key
	clientCert, err := tls.LoadX509KeyPair(clientCertFile, clientKeyFile)
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{clientCert},
		RootCAs:      certPool,
	}

	return credentials.NewTLS(config), nil
}
