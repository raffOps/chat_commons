package main

import (
	grpcController "github.com/raffops/chat/internal/auth/controller/grpc"
	authRepository "github.com/raffops/chat/internal/auth/repository"
	authService "github.com/raffops/chat/internal/auth/service"
	"github.com/raffops/chat/internal/database"
	"github.com/raffops/chat/pb"
	passwordHasher "github.com/raffops/chat/pkg/password_hasher"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	db := database.GetPostgresConn()
	defer db.Close()

	authRepo := authRepository.NewPostgresRepo(db)
	passHasher := passwordHasher.NewBcryptHasher()
	authSrv := authService.NewUserService(authRepo, passHasher)

	authController := grpcController.Server{
		AuthServer:  &grpcController.Server{},
		AuthService: authSrv,
	}

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterAuthServer(server, &authController)

	log.Println("Listening on port 9000...")
	if err := server.Serve(listener); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}
