package grpc

import (
	"context"
	authService "github.com/raffops/chat/internal/auth/service"
	"github.com/raffops/chat/pb"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log"
)

type Server struct {
	pb.AuthServer
	AuthService authService.Service
}

func (s *Server) SignUp(ctx context.Context, in *pb.SignUpRequest) (*pb.SignUpResponse, error) {
	log.Printf("SignUp: %v", in.GetName())
	user, err := s.AuthService.SignUp(ctx, in.GetName(), in.GetPassword())
	var responseErr error
	if err != nil {
		responseErr = err
	}
	return &pb.SignUpResponse{
		Id:        user.Id,
		Name:      user.Name,
		Role:      stringToRole(user.Role),
		CreatedAt: timestamppb.New(user.CreatedAt),
	}, responseErr
}

func stringToRole(roleStr string) pb.Role {
	switch roleStr {
	case "ADMIN":
		return pb.Role_ADMIN
	case "USER":
		return pb.Role_USER
	default:
		return pb.Role_USER // default value if the string doesn't match any Role
	}
}
