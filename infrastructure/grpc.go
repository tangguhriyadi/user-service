package infrastructure

import (
	"context"
	"log"
	"net"
	"strconv"
	"sync"

	"github.com/tangguhriyadi/user-service/repository"
	pb "github.com/tangguhriyadi/user-service/user"
	"google.golang.org/grpc"
)

type dataUserServiceServer struct {
	pb.UnimplementedUserServiceServer
	mu       sync.Mutex
	users    []*pb.ResponseMessage
	userRepo repository.UserRepository
}

func (d *dataUserServiceServer) GetUserById(ctx context.Context, user *pb.RequestMessage) (*pb.ResponseMessage, error) {
	userId := strconv.Itoa(int(user.Id))

	result, err := d.userRepo.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}

	ageString := strconv.Itoa(int(result.Age))

	var response pb.ResponseMessage
	response.Age = ageString
	response.Email = result.Email
	response.FullName = result.FullName

	return &response, nil

}

func NewUserServer() *dataUserServiceServer {
	s := dataUserServiceServer{}
	return &s
}

func RunGrpc() {
	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalln("error lister", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, NewUserServer())

	go func() {
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalln("erorr when serving grpc", err.Error())
		}
	}()
}
