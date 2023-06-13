package infrastructure

import (
	"context"
	"log"
	"net"
	"strconv"

	pb "github.com/tangguhriyadi/grpc-user/user"
	"github.com/tangguhriyadi/user-service/repository"
	"google.golang.org/grpc"
)

type dataUserServiceServer struct {
	pb.UnimplementedUserServiceServer
	// mu       sync.Mutex
	// users    []*pb.ResponseMessage
	userRepo repository.UserRepository
}

func (d *dataUserServiceServer) GetUserById(ctx context.Context, user *pb.RequestMessage) (*pb.ResponseMessage, error) {
	getUserId := user.GetId()
	userId := strconv.Itoa(int(getUserId))
	result, err := d.userRepo.GetById(ctx, userId)
	if err != nil {
		return nil, err
	}
	// ageString := strconv.Itoa(int(result.Age))

	var response pb.ResponseMessage
	response.Age = int32(result.Age)
	response.Email = result.Email
	response.FullName = result.FullName

	return &response, nil

}

func NewUserServer(userRepo repository.UserRepository) *dataUserServiceServer {
	// s := dataUserServiceServer{}
	return &dataUserServiceServer{
		userRepo: userRepo,
	}
}

func RunGrpc() {
	listen, err := net.Listen("tcp", ":1200")
	if err != nil {
		log.Fatalln("error lister", err)
	}

	userRepo := repository.NewUserRepository(DB)

	grpcServer := grpc.NewServer()
	pb.RegisterUserServiceServer(grpcServer, NewUserServer(userRepo))

	go func() {
		if err := grpcServer.Serve(listen); err != nil {
			log.Fatalln("erorr when serving grpc", err.Error())
		}
	}()
}
