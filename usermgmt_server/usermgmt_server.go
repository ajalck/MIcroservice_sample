package main

import (
	"context"
	"log"
	"math/rand"
	"net"

	pb "github.com/ajalck/microservice_sample1/usermgmt/protos"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

type UserManagementServer struct {
	pb.UnimplementedUserManagementServer
	user_list *pb.UserList
}

func NewUserManagementServer() *UserManagementServer {
	return &UserManagementServer{}
}

func (s *UserManagementServer) CreateUser(ctx context.Context, in *pb.NewUser) (*pb.User, error) {
	log.Printf("Received :%v", in.GetName())

	var user_id int32 = int32(rand.Intn(1000))
	newUser := &pb.User{
		Name: in.GetName(),
		Age:  in.GetAge(),
		Id:   user_id}
	s.user_list.Users = append(s.user_list.Users, newUser)
	return newUser, nil
}
func (s *UserManagementServer) GetUsers(ctx context.Context, in *pb.GetUsersParams) (*pb.UserList, error) {
	return s.user_list, nil
}

func main() {
	NewUserManagementServer()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen :%v", err)
	}

	s := grpc.NewServer()
	pb.RegisterUserManagementServer(s, &UserManagementServer{})
	log.Printf("server listening at :%v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
