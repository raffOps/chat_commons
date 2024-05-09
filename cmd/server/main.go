package main

import (
	"context"
	"fmt"
	"github.com/raffops/chat/internal/database"
	"github.com/raffops/chat/pb"
	"log"
)

type Server struct {
	pb.LoginServer
}

func (s *Server) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	log.Printf("Received: %v", in)
	return &pb.SignUpResponse{Status: &pb.HttpStatus{
		Code:    200,
		Message: fmt.Sprintf("User %s: Ok", in.Name),
		Details: nil,
	}}, nil
}

func main() {
	db := database.GetPostgresConn()
	defer db.Close()
	//listener, err := net.Listen("tcp", ":9000")
	//if err != nil {
	//	log.Fatalf("Error: %v", err)
	//}
	//server := grpc.NewServer()
	//pb.RegisterLoginServer(server, &Server{})
	//
	//log.Println("Starting server on port :9000")
	//if err := server.Serve(listener); err != nil {
	//	log.Fatalf("Failed to server: %v", err)
	//}

}
