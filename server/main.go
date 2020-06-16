package main

import (
	pb "grpc/pb/user"
	"log"
	"net"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type Server struct {
	request []*pb.GetRequest
}

func (s *Server) CreateUser(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	return &pb.CreateResponse{}, nil
}

func (s *Server) FindUser(ctx context.Context, req *pb.FindRequest) (*pb.FindResponse, error) {
	return &pb.FindResponse{}, nil
}

func (s *Server) GetUsers(ctx context.Context, req *pb.GetRequest) (*pb.GetResponse, error) {
	users := &pb.GetResponse{
		Status: "success",
		Users: []*pb.Model{
			&pb.Model{
				Id:   "asdfsadfas",
				Name: req.Keyword,
			},
		},
	}
	return users, nil
}

func main() {
	lis, err := net.Listen("tcp", ":8000")
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	srv := grpc.NewServer()

	pb.RegisterUserServer(srv, &Server{})
	srv.Serve(lis)
}
